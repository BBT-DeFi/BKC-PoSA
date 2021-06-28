pragma solidity <=0.8.6;

import "./System.sol";
import "./BKCValidatorSet.sol";
import "./StakePool.sol";
import "./Interface/IValidator.sol";

// ["0xCA35b7d915458EF540aDe6068dFe2F44E8fa733c",50,"BondStatus.UNBONDED",false]

contract ValidatorPool is System, IValidator {
    Validator[] public validators;
    mapping(address => uint256) public validatorsMap;

    BKCValidatorSet private validatorset;
    StakePool private stakepool;

    address private _BKCValidatorSetAddress;
    address private _StakePoolAddress;

    uint256 constant MIN_VALIDATOR_STAKE_AMOUNT = 10000 * (10**18); // 10000 KUB

    mapping(address => uint256) public validatorUnBondQueue;

    modifier onlyValidatorInValidatorSet(address consensusAddress) {
        require(
            validatorset.currentValidatorSetMap(consensusAddress) > 0,
            "not allow for non-active validator"
        );
        _;
    }

    modifier onlyOtherValidatorInValidatorSet(
        address consensusAddress,
        address validatorAddress
    ) {
        require(
            consensusAddress != validatorAddress,
            "only other validators can prove your innocence"
        );
        _;
    }

    modifier onlyBCKValidatorSetContract() {
        require(
            msg.sender == _BKCValidatorSetAddress,
            "can only be called from BKCValidatorSet Contract"
        );
        _;
    }

    function NumberOfValidator() public view returns (uint256) {
        return validators.length;
    }

    function init(address BKCValidatorSetAddress, address StakePoolAddress)
        external
        onlyNotInit
    {
        Validator memory firstValidator = Validator(
            0xCA35b7d915458EF540aDe6068dFe2F44E8fa733c,
            50,
            BondStatus.UNBONDED,
            false
        );
        Validator memory secondValidator = Validator(
            0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2,
            100,
            BondStatus.UNBONDED,
            false
        );
        Validator memory thirdValidator = Validator(
            0x4B20993Bc481177ec7E8f571ceCaE8A9e22C02db,
            75,
            BondStatus.UNBONDED,
            false
        );
        Validator memory fourthValidator = Validator(
            0x78731D3Ca6b7E34aC0F824c42a7cC18A495cabaB,
            120,
            BondStatus.UNBONDED,
            false
        );
        Validator memory fifthValidator = Validator(
            0x617F2E2fD72FD9D5503197092aC168c91465E7f2,
            20,
            BondStatus.UNBONDED,
            false
        );
        Validator memory sixthValidator = Validator(
            0x17F6AD8Ef982297579C203069C1DbfFE4348c372,
            200,
            BondStatus.UNBONDED,
            false
        );

        validatorset = BKCValidatorSet(BKCValidatorSetAddress);
        stakepool = StakePool(StakePoolAddress);

        _BKCValidatorSetAddress = BKCValidatorSetAddress;
        _StakePoolAddress = StakePoolAddress;

        addValidatorToSortedList(firstValidator);
        addValidatorToSortedList(secondValidator);
        addValidatorToSortedList(thirdValidator);
        addValidatorToSortedList(fourthValidator);
        addValidatorToSortedList(fifthValidator);
        addValidatorToSortedList(sixthValidator);

        alreadyInit = true;
    }

    function withdrawFund(address consensusAddress, uint256 amount)
        external
        onlyInit
    {
        require(msg.sender == consensusAddress, "can't withdraw others' fund");
        require(
            validators[validatorsMap[consensusAddress] - 1].stakeAmount >
                amount,
            "not enough fund to withdraw"
        );
        payable(consensusAddress).transfer(amount);
        validators[validatorsMap[consensusAddress] - 1].stakeAmount -= amount;
        reEvaluateValidator(consensusAddress);
    }

    function validatorTopUp(address consensusAddress)
        external
        payable
        onlyInit
    {
        require(msg.sender == consensusAddress, "can't top up for others");
    }

    function bondValidator(address consensusAddress)
        external
        onlyInit
        onlyBCKValidatorSetContract
    {
        validators[validatorsMap[consensusAddress] - 1].bondStatus = BondStatus
        .BONDED;
    }

    function unBondValidator(address consensusAddress)
        external
        onlyInit
        onlyBCKValidatorSetContract
    {
        validators[validatorsMap[consensusAddress] - 1].bondStatus = BondStatus
        .UNBONDING;
        validatorUnBondQueue[consensusAddress] =
            block.timestamp +
            unbondingPeriod;
    }

    function removeUnBondingValidatorFromQueue(address consensusAddress)
        public
        onlyInit
    {
        require(msg.sender == consensusAddress, "can't unbond other, only you");
        require(
            block.timestamp > validatorUnBondQueue[consensusAddress],
            "unbonding still in progress"
        );
        delete validatorUnBondQueue[consensusAddress];
        validators[validatorsMap[consensusAddress] - 1].bondStatus = BondStatus
        .UNBONDED;
    }

    function jailValidator(address consensusAddress)
        public
        onlyInit
        onlyValidatorInValidatorSet(msg.sender)
        onlyOtherValidatorInValidatorSet(consensusAddress, msg.sender)
    {
        validators[validatorsMap[consensusAddress]].isJail = true;
        // TODO remove the validator from active validatorset;
    }

    function unJailValidator(address consensusAddress)
        public
        onlyInit
        onlyValidatorInValidatorSet(msg.sender)
        onlyOtherValidatorInValidatorSet(consensusAddress, msg.sender)
    {
        validators[validatorsMap[consensusAddress]].isJail = false;
    }

    function registerValidator() external payable onlyInit {
        require(
            msg.value > MIN_VALIDATOR_STAKE_AMOUNT,
            "can't register to be a validator not enough staking amount sent"
        );
        addValidatorToSortedList(
            Validator(msg.sender, msg.value, BondStatus.UNBONDED, false)
        );
        validatorsMap[msg.sender] = validators.length;
    }

    function getTotalPower(address consensusAddress)
        public
        view
        returns (uint256)
    {
        return
            validators[validatorsMap[consensusAddress] - 1].stakeAmount +
            stakepool.getTotalDelegation(consensusAddress);
    }

    function addValidatorToSortedList(Validator memory new_validator) private {
        uint256 have = validatorsMap[new_validator.consensusAddress];
        if (have > 0) {
            // validator is already in the list
            return;
        } else {
            if (validators.length == 0) {
                validators.push(new_validator);
                validatorsMap[new_validator.consensusAddress] = 1;
            } else if (validators.length == 1) {
                // the incoming validator cannot have any delegation before being added into the validator list.
                uint256 totalPower = getTotalPower(
                    validators[0].consensusAddress
                );
                if (totalPower > new_validator.stakeAmount) {
                    validators.push(new_validator);
                    validatorsMap[new_validator.consensusAddress] = 2;
                } else {
                    validators.push(validators[0]);
                    validators[0] = new_validator;
                    validatorsMap[validators[0].consensusAddress] = 1;
                    validatorsMap[validators[1].consensusAddress] = 2;
                }
            } else {
                if (
                    new_validator.stakeAmount >
                    getTotalPower(validators[0].consensusAddress)
                ) {
                    validators.push(new_validator);
                    for (uint256 i = validators.length - 1; i > 0; i--) {
                        validators[i] = validators[i - 1];
                        validatorsMap[validators[i].consensusAddress] = i + 1;
                    }
                    validators[0] = new_validator;
                    validatorsMap[new_validator.consensusAddress] = 1;
                } else if (
                    new_validator.stakeAmount <=
                    getTotalPower(
                        validators[validators.length - 1].consensusAddress
                    )
                ) {
                    validators.push(new_validator);
                    validatorsMap[new_validator.consensusAddress] = validators
                    .length;
                } else {
                    for (uint256 i = 0; i < validators.length; i++) {
                        // change the sequence according to the validator's total power;
                        if (
                            new_validator.stakeAmount >=
                            getTotalPower(validators[i].consensusAddress)
                        ) {
                            validators.push(new_validator);
                            for (
                                uint256 j = validators.length - 2;
                                j >= i;
                                j--
                            ) {
                                validators[j + 1] = validators[j];
                                validatorsMap[
                                    validators[j + 1].consensusAddress
                                ] = j + 2;
                            }
                            validators[i] = new_validator;
                            validatorsMap[new_validator.consensusAddress] =
                                i +
                                1;
                            break;
                        }
                    }
                }
            }
        }
    }

    function reEvaluateValidator(address consensusAddress) private {
        // TODO : check if the validator still has enough fund to stay as a validator.
    }
}

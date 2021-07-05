// SPDX-License-Identifier: MIT
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

    uint256 constant MIN_VALIDATOR_STAKE_AMOUNT = 10 ether; //10000 * (10**18); // 10000 KUB

    address constant ADMIN = 0x090fb1c3d66303358806836DF2B5b44fcd3e582f; // account 18

    mapping(address => uint256) public validatorUnBondQueue;

    mapping(address => uint256) public validatorJailQueue;

    mapping(address => uint256) public validatorRemoveQueue;

    modifier onlyAdmin() {
        require(msg.sender == ADMIN, "admin only");
        _;
    }

    modifier onlyValidator() {
        require(
            validatorsMap[msg.sender] > 0,
            "only validators can call this function"
        );
        _;
    }

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
            "only other validators can do this operation"
        );
        _;
    }

    modifier onlyBKCValidatorSetContract() {
        require(
            msg.sender == _BKCValidatorSetAddress,
            "can only be called from BKCValidatorSet Contract"
        );
        _;
    }

    modifier onlyNotBonded(address consensusAddress) {
        require(
            validators[validatorsMap[consensusAddress] - 1].bondStatus ==
                BondStatus.UNBONDED,
            "can't do operation when unbonding or bonded"
        );
        _;
    }

    modifier onlyStakePoolContractAndThisContract {
        require(msg.sender == address(this) || msg.sender == _StakePoolAddress);
        _;
    }

    function NumberOfValidator() public view returns (uint256) {
        return validators.length;
    }

    function init(address BKCValidatorSetAddress, address StakePoolAddress)
        external
        payable
        onlyNotInit
    {
        // Validator memory firstValidator = Validator(
        //     0xCA35b7d915458EF540aDe6068dFe2F44E8fa733c,
        //     50 ether,
        //     BondStatus.UNBONDED,
        //     false
        // );
        // Validator memory secondValidator = Validator(
        //     0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2,
        //     100,
        //     BondStatus.UNBONDED,
        //     false
        // );
        // Validator memory thirdValidator = Validator(
        //     0x4B20993Bc481177ec7E8f571ceCaE8A9e22C02db,
        //     75,
        //     BondStatus.UNBONDED,
        //     false
        // );
        // Validator memory fourthValidator = Validator(
        //     0x78731D3Ca6b7E34aC0F824c42a7cC18A495cabaB,
        //     120,
        //     BondStatus.UNBONDED,
        //     false
        // );
        // Validator memory fifthValidator = Validator(
        //     0x617F2E2fD72FD9D5503197092aC168c91465E7f2,
        //     20,
        //     BondStatus.UNBONDED,
        //     false
        // );
        // Validator memory sixthValidator = Validator(
        //     0x17F6AD8Ef982297579C203069C1DbfFE4348c372,
        //     200,
        //     BondStatus.UNBONDED,
        //     false
        // );

        validatorset = BKCValidatorSet(BKCValidatorSetAddress);
        stakepool = StakePool(StakePoolAddress);

        _BKCValidatorSetAddress = BKCValidatorSetAddress;
        _StakePoolAddress = StakePoolAddress;

        // validators.push(firstValidator);
        // validatorsMap[firstValidator.consensusAddress] = 1;
        // validators.push(secondValidator);
        // validatorsMap[secondValidator.consensusAddress] = 2;
        // validators.push(thirdValidator);
        // validatorsMap[thirdValidator.consensusAddress] = 3;
        // validators.push(fourthValidator);
        // validatorsMap[fourthValidator.consensusAddress] = 4;
        // validators.push(fifthValidator);
        // validatorsMap[fifthValidator.consensusAddress] = 5;
        // validators.push(sixthValidator);
        // validatorsMap[sixthValidator.consensusAddress] = 6;

        // addValidatorToSortedList(firstValidator);
        // addValidatorToSortedList(secondValidator);
        // addValidatorToSortedList(thirdValidator);
        // addValidatorToSortedList(fourthValidator);
        // addValidatorToSortedList(fifthValidator);
        // addValidatorToSortedList(sixthValidator);

        alreadyInit = true;
    }

    function withdrawFund(uint256 amount)
        external
        onlyInit
        onlyNotBonded(msg.sender)
    {
        require(
            validators[validatorsMap[msg.sender] - 1].stakeAmount > amount,
            "not enough fund to withdraw"
        );
        payable(msg.sender).transfer(amount);
        validators[validatorsMap[msg.sender] - 1].stakeAmount -= amount;
        bool res = reEvaluateValidator(msg.sender);
        if (res) {
            // validator still have enough
            //rePositionValidator(consensusAddress, false);
        }
    }

    function validatorTopUp()
        external
        payable
        onlyInit
        onlyNotBonded(msg.sender)
    {
        require(msg.value > 0, "can't top up with 0 amount");
        validators[validatorsMap[msg.sender] - 1].stakeAmount += msg.value;
        //rePositionValidator(consensusAddress, true);
    }

    // bonding validator does not bond its delegators' delegations, just that it will uses all the bonded delegation from the before.
    function bondValidator(address consensusAddress)
        external
        onlyInit
        onlyBKCValidatorSetContract
    {
        validators[validatorsMap[consensusAddress] - 1].bondStatus = BondStatus
        .BONDED;
        delete validatorUnBondQueue[consensusAddress]; // remove from the unbondedqueue if any. (this will only be called by updateValidatorSet)
    }

    // unbonding validator does not unbond its delegators' delegations, just that it will allow delegators to unbond fund quickly after unbonded.
    function unBondValidator(address consensusAddress)
        external
        onlyInit
        onlyBKCValidatorSetContract
    {
        validators[validatorsMap[consensusAddress] - 1].bondStatus = BondStatus
        .UNBONDING;
        validatorUnBondQueue[consensusAddress] =
            block.timestamp +
            unbondingPeriod;
    }

    function removeUnBondingValidatorFromQueue() public onlyInit onlyValidator {
        require(
            block.timestamp > validatorUnBondQueue[msg.sender],
            "unbonding still in progress"
        );
        require(
            validatorJailQueue[msg.sender] == 0,
            "can't unbond, validator is jailed, also the validator should not be in the unbonding queue"
        );
        delete validatorUnBondQueue[msg.sender];
        validators[validatorsMap[msg.sender] - 1].bondStatus = BondStatus
        .UNBONDED;
    }

    function removeJailValidatorFromQueue() public onlyInit onlyValidator {
        require(
            block.timestamp > validatorJailQueue[msg.sender],
            "jailing still in progress"
        );
        delete validatorJailQueue[msg.sender];
        validators[validatorsMap[msg.sender] - 1].bondStatus = BondStatus
        .UNBONDED;
        validators[validatorsMap[msg.sender] - 1].isJail = false;
    }

    function removeRemovingValidatorFromQueue() public onlyInit onlyValidator {
        require(
            block.timestamp > validatorRemoveQueue[msg.sender],
            "removing still in progress"
        );
        require(
            validatorRemoveQueue[msg.sender] > 0,
            "validator must be in the queue"
        );
        delete validatorRemoveQueue[msg.sender];
        address[] memory delegators = stakepool.getDelegators(msg.sender);
        for (uint256 i = 0; i < delegators.length; i++) {
            stakepool.removeDelegation(delegators[i], msg.sender); // return all delegations
        }
        payable(msg.sender).transfer(
            validators[validatorsMap[msg.sender] - 1].stakeAmount
        ); // returns the fund to the validator
        uint256 index = validatorsMap[msg.sender] - 1;
        for (uint256 i = index; i < validators.length - 1; i++) {
            // remove the validator from pool
            validators[i] = validators[i + 1]; // shift left
            validatorsMap[validators[i].consensusAddress] = i + 1;
        }
        validators.pop();
        delete validatorsMap[msg.sender]; // remove the validator from pool
    }

    function jailValidator(address consensusAddress) public onlyInit onlyAdmin {
        require(
            validatorset.currentValidatorSetMap(consensusAddress) > 0,
            "can't jail non-active validator"
        );
        validators[validatorsMap[consensusAddress] - 1].isJail = true;
        validatorset.jailValidator(consensusAddress); // remove from validator set and deal with rewards.
        updateBondStatus(consensusAddress, BondStatus.UNBONDING);
        validatorJailQueue[consensusAddress] =
            block.timestamp +
            VALIDATOR_JAIL_PERIOD;
        validators[validatorsMap[consensusAddress] - 1].stakeAmount -=
            5 *
            (10**18);
        reEvaluateValidator(consensusAddress);
    }

    function registerValidator() external payable onlyInit {
        require(
            msg.value >= MIN_VALIDATOR_STAKE_AMOUNT,
            "can't register to be a validator not enough staking amount sent"
        );
        require(
            validatorsMap[msg.sender] == 0,
            "can't register same validator"
        );
        //addValidatorToSortedList(Validator(msg.sender, msg.value, BondStatus.UNBONDED, false));
        validators.push(
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

    function updateBondStatus(address consensusAddress, BondStatus b)
        private
        onlyInit
    {
        validators[validatorsMap[consensusAddress] - 1].bondStatus = b;
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
                    // be the top
                    validators.push(new_validator);
                    for (uint256 i = validators.length - 1; i > 0; i--) {
                        validators[i] = validators[i - 1]; // shift right
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
                    // be the least
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

    function reEvaluateValidator(address consensusAddress)
        private
        returns (bool)
    {
        if (
            validators[validatorsMap[consensusAddress] - 1].stakeAmount <
            MIN_VALIDATOR_STAKE_AMOUNT
        ) {
            uint256 index = validatorsMap[consensusAddress] - 1;
            for (uint256 i = index; i < validators.length - 1; i++) {
                validators[i] = validators[i + 1]; // shift left
                validatorsMap[validators[i].consensusAddress] = i + 1;
            }
            payable(consensusAddress).transfer(
                validators[validatorsMap[consensusAddress] - 1].stakeAmount
            );
            address[] memory delegators = stakepool.getDelegators(
                consensusAddress
            );
            for (uint256 i = 0; i < delegators.length; i++) {
                stakepool.removeDelegation(delegators[i], consensusAddress);
            }
            validators.pop();
            delete validatorsMap[consensusAddress];
            return false;
        }
        return true;
    }

    function validatorRemove() external onlyInit {
        require(validatorsMap[msg.sender] > 0, "can only remove a validator");
        validatorRemoveQueue[msg.sender] =
            block.timestamp +
            VALIDATOR_REMOVE_PERIOD; // remove period
        delete validatorUnBondQueue[msg.sender];
    }

    function rePositionValidator(address consensusAddress, bool operation)
        public
        onlyStakePoolContractAndThisContract
    {
        // operation true for top up and false for withdraw
        uint256 index = validatorsMap[consensusAddress] - 1;
        uint256 this_totalPower = getTotalPower(consensusAddress);
        if (operation) {
            // top up
            if (index == 0) {
                return; // already the largest
            }
            if (
                this_totalPower >= getTotalPower(validators[0].consensusAddress)
            ) {
                // become the top
                Validator memory v = validators[index];
                for (uint256 i = index; i > 0; i--) {
                    validators[i] = validators[i - 1]; // shift right
                    validatorsMap[validators[i].consensusAddress] = i + 1;
                }
                validators[0] = v;
                validatorsMap[v.consensusAddress] = 1;
                return;
            }
            for (uint256 i = index; i > 0; i--) {
                // look to the left
                if (
                    getTotalPower(validators[i - 1].consensusAddress) >
                    this_totalPower
                ) {
                    // found the should-be place
                    Validator memory v = validators[index];
                    for (uint256 j = index; j > i; j--) {
                        validators[j] = validators[j - 1]; // shift right
                        validatorsMap[validators[j].consensusAddress] = j + 1;
                    }
                    validators[i] = v;
                    validatorsMap[v.consensusAddress] = i + 1;
                    return;
                }
            }
        } else {
            // withdraw
            if (index == validators.length) {
                return; // already the smallest
            }
            if (
                this_totalPower <=
                getTotalPower(
                    validators[validators.length - 1].consensusAddress
                )
            ) {
                // become the least
                Validator memory v = validators[index];
                for (uint256 i = index; i < validators.length - 1; i++) {
                    validators[i] = validators[i + 1]; // shift left
                    validatorsMap[validators[i].consensusAddress] = i + 1;
                }
                validators[validators.length - 1] = v;
                validatorsMap[v.consensusAddress] = validators.length;
                return;
            }
            for (uint256 i = index; i < validators.length; i++) {
                // look to the right
                if (
                    getTotalPower(validators[i + 1].consensusAddress) <
                    this_totalPower
                ) {
                    // found the should-be place
                    Validator memory v = validators[index];
                    for (uint256 j = index; j < i; j++) {
                        validators[j] = validators[j + 1]; // shift left
                        validatorsMap[validators[j].consensusAddress] = j + 1;
                    }
                    validators[i] = v;
                    validatorsMap[v.consensusAddress] = i + 1;
                    return;
                }
            }
        }
    }
}

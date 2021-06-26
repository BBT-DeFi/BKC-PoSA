pragma solidity <=0.8.6;

import "./System.sol";
import "./BKCValidatorSet.sol";
import "./Interface/IValidator.sol";
import "./Interface/IDelegation.sol";

contract ValidatorPool is System, IValidator {
    Validator[] public validators;
    mapping(address => uint256) public validatorsMap;
    BKCValidatorSet private validatorset;

    address private _BKCValidatorSetAddress;

    uint256 constant MIN_VALIDATOR_STAKE_AMOUNT = 10000 * (10**18); // 10000 KUB

    mapping(address => uint256) public validatorUnBondQueue;

    modifier onlyValidatorInValidatorSet(address concensusAddress) {
        require(
            validatorset.currentValidatorSetMap(concensusAddress) > 0,
            "not allow for non-active validator"
        );
        _;
    }

    modifier onlyOtherValidatorInValidatorSet(
        address concensusAddress,
        address validatorAddress
    ) {
        require(
            concensusAddress != validatorAddress,
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

    function init(address BKCValidatorSetAddress) external onlyNotInit {
        Validator memory firstValidator = Validator(
            0xCA35b7d915458EF540aDe6068dFe2F44E8fa733c,
            48648428258251948625,
            BondStatus.UNBONDED,
            false
        );
        Validator memory secondValidator = Validator(
            0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2,
            35481719561823448621,
            BondStatus.UNBONDED,
            false
        );

        validators.push(firstValidator);
        validators.push(secondValidator);
        validatorsMap[address(0xCA35b7d915458EF540aDe6068dFe2F44E8fa733c)] = 1;
        validatorsMap[address(0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2)] = 2;
        validatorset = BKCValidatorSet(BKCValidatorSetAddress);
        _BKCValidatorSetAddress = BKCValidatorSetAddress;
        alreadyInit = true;
    }

    function withdrawFund(address concensusAddress, uint256 amount)
        external
        onlyInit
    {
        require(msg.sender == concensusAddress, "can't withdraw others' fund");
        require(
            validators[validatorsMap[concensusAddress] - 1].stakeAmount >
                amount,
            "not enough fund to withdraw"
        );
        payable(concensusAddress).transfer(amount);
        validators[validatorsMap[concensusAddress] - 1].stakeAmount -= amount;
        reEvaluateValidator(concensusAddress);
    }

    function bondValidator(address concensusAddress)
        external
        onlyInit
        onlyBCKValidatorSetContract
    {
        require(msg.sender == concensusAddress, "you can't bond others");
        validators[validatorsMap[concensusAddress]].bondStatus = BondStatus
        .BONDED;
    }

    function unBondValidator(address concensusAddress)
        external
        onlyInit
        onlyBCKValidatorSetContract
    {
        require(msg.sender == concensusAddress, "you can't unbond others");
        validators[validatorsMap[concensusAddress]].bondStatus = BondStatus
        .UNBONDING;
        validatorUnBondQueue[concensusAddress] =
            block.timestamp +
            unbondingPeriod;
    }

    function removeUnBondingValidatorFromQueue(address concensusAddress)
        public
        onlyInit
    {
        require(msg.sender == concensusAddress, "can't unbond other, only you");
        require(
            block.timestamp > validatorUnBondQueue[concensusAddress],
            "unbonding still in progress"
        );
        delete validatorUnBondQueue[concensusAddress];
    }

    function jailValidator(address concensusAddress)
        public
        onlyInit
        onlyValidatorInValidatorSet(msg.sender)
        onlyOtherValidatorInValidatorSet(concensusAddress, msg.sender)
    {
        validators[validatorsMap[concensusAddress]].isJail = true;
        // TODO remove the validator from validatorset;
    }

    function unJailValidator(address concensusAddress)
        public
        onlyInit
        onlyValidatorInValidatorSet(msg.sender)
        onlyOtherValidatorInValidatorSet(concensusAddress, msg.sender)
    {
        validators[validatorsMap[concensusAddress]].isJail = false;
    }

    function registerValidator() external payable onlyInit {
        require(
            msg.value > MIN_VALIDATOR_STAKE_AMOUNT,
            "can't register to be a validator not enough staking amount sent"
        );
        validators.push(
            Validator(msg.sender, msg.value, BondStatus.UNBONDED, false)
        );
        validatorsMap[msg.sender] = validators.length;
    }

    function reEvaluateValidator(address concensusAddress) private {
        // TODO : check if the validator still has enough fund to stay as a validator.
    }
}

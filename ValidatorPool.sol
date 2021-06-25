pragma solidity <=0.8.6;

import "./System.sol";
import "./BKCValidatorSet.sol";
import "./Interface/IValidator.sol";

contract ValidatorPool is System, IValidator {
    Validator[] public validators;
    mapping(address => uint256) public validatorsMap;
    BKCValidatorSet public validatorset;

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

    function init() external override onlyNotInit {
        Validator memory aValidator = Validator(
            address(0x70F657164e5b75689b64B7fd1fA275F334f28e18),
            63306535440133,
            false,
            false
        );
        validators.push(aValidator);
        validatorsMap[address(0x70F657164e5b75689b64B7fd1fA275F334f28e18)] = 1;
        validatorset = BKCValidatorSet(
            address(0xaE036c65C649172b43ef7156b009c6221B596B8b)
        );
        alreadyInit = true;
    }

    function bondValidator(address concensusAddress)
        public
        onlyInit
        onlyValidatorInValidatorSet(msg.sender)
    {
        require(msg.sender == concensusAddress, "you can't bond others");
        validators[validatorsMap[concensusAddress]].isBond = true;
    }

    function unBondValidator(address concensusAddress)
        public
        onlyInit
        onlyValidatorInValidatorSet(msg.sender)
    {
        require(msg.sender == concensusAddress, "you can't unbond others");
        validators[validatorsMap[concensusAddress]].isBond = false;
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
}

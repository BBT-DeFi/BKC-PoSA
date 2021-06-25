pragma solidity <=0.8.6;

import "./System.sol";
import "./ValidatorPool.sol";
import "./Interface/IValidator.sol";
import "./Interface/IDelegation.sol";

contract StakePool is System, IValidator, IDelegation {
    mapping(address => DelegationObject) public ValidatorDelegation;

    function init() external override onlyNotInit {
        alreadyInit = true;
    }

    function delegate(address validatorConcensus, uint256 amount)
        external
        onlyInit
    {
        // TODO remove amount from msg.sender
        ValidatorDelegation[validatorConcensus].delegateAmountOfEach[
            msg.sender
        ] += amount;
        ValidatorDelegation[validatorConcensus].totalDelegation += amount;
    }

    function undelegate(address validatorConcensus, uint256 amount)
        external
        onlyInit
    {
        require(
            ValidatorDelegation[validatorConcensus].delegateAmountOfEach[
                msg.sender
            ] > amount,
            "not enough amount to undelegate"
        );
        address validatorPoolAddress = address(
            0xe2899bddFD890e320e643044c6b95B9B0b84157A
        );
        ValidatorPool vldpool = ValidatorPool(validatorPoolAddress);
        Validator memory aValidator = vldpool.validators(
            vldpool.validatorsMap(validatorConcensus) - 1
        );
        // address aValidatorAddress = vldpool.validators(vldpool.validatorsMap(validatorConcensus)-1).concensusAddress;
        // Validator memory aValidatorObject = Validator(vldpool.validators(vldpool.validatorsMap(validatorConcensus)-1));
        if (aValidator.isBond) {
            // do logic for unbonding the delegation
        } else {
            ValidatorDelegation[validatorConcensus].delegateAmountOfEach[
                msg.sender
            ] -= amount;
            vldpool.validatorsMap[validatorConcensus].stakeAmount -= amount;
        }
    }
}

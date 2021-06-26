pragma solidity <=0.8.6;

import "./System.sol";
import "./ValidatorPool.sol";
import "./Interface/IValidator.sol";
import "./Interface/IDelegation.sol";

contract StakePool is System, IValidator, IDelegation {
    mapping(address => DelegationObject) public ValidatorDelegation;
    mapping(address => userDelegation) public UserDelegation;

    mapping(address => uint256) public UserUnDelegateQueue;

    ValidatorPool vldpool;

    function init(address validatorPoolAddress) external onlyNotInit {
        vldpool = ValidatorPool(validatorPoolAddress);
        alreadyInit = true;
    }

    function delegate(address validatorConcensus) external payable onlyInit {
        require(
            vldpool.validatorsMap(validatorConcensus) > 0,
            "can't delegate to a non-validator"
        );
        if (UserDelegation[msg.sender].concensusAddress == validatorConcensus) {
            // if contribute more to the same validator
            ValidatorDelegation[validatorConcensus].delegateAmountOfEach[
                msg.sender
            ] += msg.value;
            ValidatorDelegation[validatorConcensus].totalDelegation += msg
            .value;
            UserDelegation[msg.sender].bondedAmount += msg.value;
        } else {
            if (
                UserDelegation[msg.sender].bondedAmount == 0 &&
                UserDelegation[msg.sender].unbondingAmount == 0
            ) {
                // change validator
                ValidatorDelegation[validatorConcensus].delegateAmountOfEach[
                    msg.sender
                ] += msg.value;
                ValidatorDelegation[validatorConcensus].totalDelegation += msg
                .value;
                UserDelegation[msg.sender] = userDelegation(
                    validatorConcensus,
                    msg.value,
                    0
                );
            } else {
                revert(
                    "can't make delegation to another validator when already delegated to one"
                ); // one validator at a time.
            }
        }
    }

    function undelegate(address validatorConcensus, uint256 amount)
        external
        onlyInit
    {
        require(
            ValidatorDelegation[validatorConcensus].delegateAmountOfEach[
                msg.sender
            ] > amount,
            "not enough amount to unbond"
        );
        require(
            UserDelegation[msg.sender].bondedAmount > amount,
            "not enough amount to unbond"
        );

        uint256 index = vldpool.validatorsMap(validatorConcensus) - 1;
        (, , BondStatus bond, ) = vldpool.validators(index);
        if (bond == BondStatus.BONDED) {
            ValidatorDelegation[validatorConcensus].delegateAmountOfEach[
                msg.sender
            ] -= amount;
            UserDelegation[msg.sender].bondedAmount -= amount;
            UserDelegation[msg.sender].unbondingAmount += amount;
            UserUnDelegateQueue[msg.sender] = block.timestamp + unbondingPeriod;
        } else {
            ValidatorDelegation[validatorConcensus].delegateAmountOfEach[
                msg.sender
            ] -= amount;
            UserDelegation[msg.sender].bondedAmount -= amount;
            payable(msg.sender).transfer(amount);
        }
        // ValidatorDelegation[validatorConcensus].delegateAmountOfEach[msg.sender] -= amount;
        // UserDelegation[msg.sender].bondedAmount -= amount;
        // UserDelegation[msg.sender].unbondingAmount += amount;
        // UserUnDelegateQueue[msg.sender] = block.timestamp + unbondingPeriod;
    }

    function removeUnbondingUserFromUnbondingQueue() external onlyInit {
        require(
            block.timestamp > UserUnDelegateQueue[msg.sender],
            "unbonding still in progress"
        );
        delete UserUnDelegateQueue[msg.sender];
        payable(msg.sender).transfer(
            UserDelegation[msg.sender].unbondingAmount
        );
        UserDelegation[msg.sender].unbondingAmount = 0;
    }
}

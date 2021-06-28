pragma solidity <=0.8.6;

import "./System.sol";
import "./ValidatorPool.sol";
import "./Interface/IValidator.sol";
import "./Interface/IDelegation.sol";

contract StakePool is System, IValidator, IDelegation {
    mapping(address => DelegationObject) public ValidatorDelegation;
    mapping(address => userDelegation) internal UserDelegation;

    mapping(address => UnBondingQueueStruct[]) public UserUnDelegateQueue;

    ValidatorPool vldpool;

    struct UnBondingQueueStruct {
        uint256 amount;
        uint256 time;
        address validator;
    }

    function init(address validatorPoolAddress) external onlyNotInit {
        vldpool = ValidatorPool(validatorPoolAddress);
        alreadyInit = true;
    }

    function getTotalDelegation(address consensusAddress)
        public
        view
        returns (uint256)
    {
        return ValidatorDelegation[consensusAddress].totalDelegation;
    }

    function delegate(address validatorConsensus) external payable onlyInit {
        require(
            vldpool.validatorsMap(validatorConsensus) > 0,
            "can't delegate to a non-validator"
        );
        (, , BondStatus b, ) = vldpool.validators(
            vldpool.validatorsMap(validatorConsensus) - 1
        );
        require(
            b == BondStatus.UNBONDED,
            "can't delegate to an bonding delegator or an unbonding delegator"
        );
        ValidatorDelegation[validatorConsensus].delegateAmountOfEach[
            msg.sender
        ] += msg.value;
        ValidatorDelegation[validatorConsensus].totalDelegation += msg.value;
        vldpool.rePositionValidator(validatorConsensus, true); // the same as validator top up ; comment this line for firstMethod
        UserDelegation[msg.sender].bondedAmountForEach[
            validatorConsensus
        ] += msg.value;
    }

    function undelegate(address validatorConsensus, uint256 amount)
        external
        onlyInit
    {
        require(
            ValidatorDelegation[validatorConsensus].delegateAmountOfEach[
                msg.sender
            ] >= amount,
            "not enough amount to unbond"
        );

        uint256 index = vldpool.validatorsMap(validatorConsensus) - 1;
        (, , BondStatus bond, ) = vldpool.validators(index);
        if (bond == BondStatus.BONDED) {
            // if bond then all the delegation should be in bondedAmountForEach of the validator so unbond it.
            //ValidatorDelegation[validatorConsensus].delegateAmountOfEach[msg.sender] -= amount;
            UserDelegation[msg.sender].bondedAmountForEach[
                validatorConsensus
            ] -= amount;
            UserDelegation[msg.sender].unbondingAmountForEach[
                validatorConsensus
            ] += amount;
            UserUnDelegateQueue[msg.sender].push(
                UnBondingQueueStruct(
                    amount,
                    block.timestamp + unbondingPeriod,
                    validatorConsensus
                )
            );
        } else {
            // if not bond then transfer the amount (when not bond all the delegation for this validator should be in unbondedAmountForEach of the validator)
            ValidatorDelegation[validatorConsensus].delegateAmountOfEach[
                msg.sender
            ] -= amount;
            ValidatorDelegation[validatorConsensus].totalDelegation -= amount;
            vldpool.rePositionValidator(validatorConsensus, false); // the same as validator withdraw ; comment this line for firstMethod
            UserDelegation[msg.sender].bondedAmountForEach[
                validatorConsensus
            ] -= amount;
            payable(msg.sender).transfer(amount);
        }
    }

    function removeUnbondingUserFromUnbondingQueue() external onlyInit {
        // remove does not mean come to unbondedAmountForEach, it transfers out.
        for (uint256 i = 0; i < UserUnDelegateQueue[msg.sender].length; i++) {
            uint256 amount = UserUnDelegateQueue[msg.sender][i].amount;
            uint256 time = UserUnDelegateQueue[msg.sender][i].time;
            address validator = UserUnDelegateQueue[msg.sender][i].validator;
            if (time < block.timestamp) {
                // unbonding finishes for this undelegation.
                delete UserUnDelegateQueue[msg.sender][i];
                payable(msg.sender).transfer(amount);
                UserDelegation[msg.sender].unbondingAmountForEach[
                    validator
                ] -= amount;
                // the delegation for the validator is deducted when the unbonding is finished and user remove it from queue
                ValidatorDelegation[validator].delegateAmountOfEach[
                    msg.sender
                ] -= amount;
                ValidatorDelegation[validator].totalDelegation -= amount;
                vldpool.rePositionValidator(validator, false); // the same as validator withdraw ; comment this line for firstMethod
            }
        }
    }
}

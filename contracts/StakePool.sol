// SPDX-License-Identifier: MIT
pragma solidity <=0.8.6;

import "./System.sol";
import "./ValidatorPool.sol";
import "./Interface/IValidator.sol";
import "./Interface/IDelegation.sol";

contract StakePool is System, IValidator, IDelegation {
    mapping(address => DelegationObject) public ValidatorDelegation;
    mapping(address => userDelegation) internal UserDelegation;

    mapping(address => UnBondingQueueStruct[]) public UserUnDelegateQueue;

    address _ValidatorPoolAddress;

    uint256[] private queueAmount;
    uint256[] private queueTime;
    address[] private queueValidator;

    ValidatorPool vldpool;

    struct UnBondingQueueStruct {
        uint256 amount;
        uint256 time;
        address validator;
    }

    modifier onlyValidatorPoolContract() {
        require(msg.sender == _ValidatorPoolAddress);
        _;
    }

    function init(address validatorPoolAddress) external onlyNotInit {
        _ValidatorPoolAddress = validatorPoolAddress;
        vldpool = ValidatorPool(validatorPoolAddress);
        alreadyInit = true;
    }

    function getUnbondingValue(address delegator) external view onlyInit returns(UnBondingQueueStruct[] memory){
        return UserUnDelegateQueue[delegator];
    }

    function getUnbondingQueueAmount(address delegator)
        external
        onlyInit
        returns (uint256[] memory)
    {
        delete queueAmount;
        UnBondingQueueStruct[] memory Queue = UserUnDelegateQueue[delegator];
        for (uint256 i = 0; i < Queue.length; i++) {
            uint256 amount = Queue[i].amount;
            queueAmount.push(amount);
        }
        return queueAmount;
    }

    function getUnbondingQueueTime(address delegator)
        external
        onlyInit
        returns (uint256[] memory)
    {
        delete queueTime;
        UnBondingQueueStruct[] memory Queue = UserUnDelegateQueue[delegator];
        for (uint256 i = 0; i < Queue.length; i++) {
            uint256 time = Queue[i].time;
            queueTime.push(time);
        }
        return queueTime;
    }

    function getUnbondingQueueValidator(address delegator)
        external
        onlyInit
        returns (address[] memory)
    {
        delete queueValidator;
        UnBondingQueueStruct[] memory Queue = UserUnDelegateQueue[delegator];
        for (uint256 i = 0; i < Queue.length; i++) {
            address validator = Queue[i].validator;
            queueValidator.push(validator);
        }
        return queueValidator;
    }

    function getDelegators(address consensusAddress)
        external
        view
        returns (address[] memory)
    {
        return ValidatorDelegation[consensusAddress].delegators;
    }

    function getTotalDelegation(address consensusAddress)
        public
        view
        returns (uint256)
    {
        return ValidatorDelegation[consensusAddress].totalDelegation;
    }

    function getTotalDelegationExcludeUnbonding(address consensusAddress)
        public
        view
        returns (uint256)
    {
        DelegationObject storage DelOb = ValidatorDelegation[consensusAddress];
        address[] storage delegators = DelOb.delegators;
        uint256 totalDelegation = ValidatorDelegation[consensusAddress]
        .totalDelegation;
        for (uint256 i = 0; i < delegators.length; i++) {
            address delegator = delegators[i];
            for (
                uint256 j = 0;
                j < UserUnDelegateQueue[delegator].length;
                j++
            ) {
                UnBondingQueueStruct storage unbondStruct = UserUnDelegateQueue[
                    delegator
                ][j];
                uint256 amount = unbondStruct.amount;
                if (unbondStruct.validator == consensusAddress) {
                    totalDelegation -= amount;
                }
            }
        }
        return totalDelegation;
    }

    function getUserDelegationBondedAmount(address consensusAddress)
        public
        view
        returns (uint256)
    {
        return UserDelegation[msg.sender].bondedAmountForEach[consensusAddress];
    }

    function getUserDelegationBondedAmountCallable(
        address delegator,
        address consensusAddress
    ) public view returns (uint256) {
        return UserDelegation[delegator].bondedAmountForEach[consensusAddress];
    }

    function getUserDelegationUnbondingAmount(address consensusAddress)
        public
        view
        returns (uint256)
    {
        return
            UserDelegation[msg.sender].unbondingAmountForEach[consensusAddress];
    }

    function getUserDelegationUnbondingAmountCallable(
        address delegator,
        address consensusAddress
    ) public view returns (uint256) {
        return
            UserDelegation[delegator].unbondingAmountForEach[consensusAddress];
    }

    function getUserTotalDelegationAmount(address consensusAddress)
        external
        view
        returns (uint256)
    {
        return
            getUserDelegationBondedAmount(consensusAddress) +
            getUserDelegationUnbondingAmount(consensusAddress);
    }

    function getUserTotalDelegationAmountCallable(
        address delegator,
        address consensusAddress
    ) external view returns (uint256) {
        return
            getUserDelegationBondedAmountCallable(delegator, consensusAddress) +
            getUserDelegationUnbondingAmountCallable(
                delegator,
                consensusAddress
            );
    }

    function getDelegationAmountOfEach(
        address consensusAddress,
        address delegator
    ) external view returns (uint256) {
        return
            ValidatorDelegation[consensusAddress].delegateAmountOfEach[
                delegator
            ];
    }

    function removeDelegation(address delegator, address validator)
        external
        onlyInit
        onlyValidatorPoolContract
    {
        uint256 totalDelegationOfDelegator = ValidatorDelegation[validator]
        .delegateAmountOfEach[delegator];
        require(
            address(this).balance >= totalDelegationOfDelegator,
            "not enough fund to return to delegator, for some reason..."
        );
        // require(
        //     UserDelegation[delegator].unbondingAmountForEach[validator] == 0,
        //     "still have unbonding fund of this validator, which should not be possible"
        // );
        payable(delegator).transfer(totalDelegationOfDelegator);
        ValidatorDelegation[validator]
        .totalDelegation -= totalDelegationOfDelegator;
        ValidatorDelegation[validator].delegateAmountOfEach[delegator] = 0;
        UserDelegation[delegator].bondedAmountForEach[validator] = 0;
        UserDelegation[delegator].unbondingAmountForEach[validator] = 0;
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
        // update the ValidatorDelegation object in delegateAmountOfEach, totalDelegation, and delegators array
        ValidatorDelegation[validatorConsensus].delegateAmountOfEach[
            msg.sender
        ] += msg.value;
        ValidatorDelegation[validatorConsensus].totalDelegation += msg.value;
        bool found_delegator = false;
        for (
            uint256 i = 0;
            i < ValidatorDelegation[validatorConsensus].delegators.length;
            i++
        ) {
            if (
                ValidatorDelegation[validatorConsensus].delegators[i] ==
                msg.sender
            ) {
                found_delegator = true;
                break;
            }
        }
        if (!found_delegator) {
            ValidatorDelegation[validatorConsensus].delegators.push(msg.sender);
        }
        //vldpool.rePositionValidator(validatorConsensus, true); // the same as validator top up
        // update the UserDelegation object in bondedAmountForEach, and validators array
        UserDelegation[msg.sender].bondedAmountForEach[
            validatorConsensus
        ] += msg.value;
        bool found = false;
        for (
            uint256 i = 0;
            i < UserDelegation[msg.sender].validators.length;
            i++
        ) {
            if (
                UserDelegation[msg.sender].validators[i] == validatorConsensus
            ) {
                found = true;
                break;
            }
        }
        if (!found) {
            UserDelegation[msg.sender].validators.push(validatorConsensus);
        }
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
        if (bond == BondStatus.BONDED || bond == BondStatus.UNBONDING) {
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
            //vldpool.rePositionValidator(validatorConsensus, false); // the same as validator withdraw
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
                ValidatorDelegation[validator].delegateAmountOfEach[
                    msg.sender
                ] -= amount; // remove from the validator's delegation
                ValidatorDelegation[validator].totalDelegation -= amount;
                // no need to remove user delegation here because the undelegate only affect power of validator in the pool and not in the active set.
                //vldpool.rePositionValidator(validator, false); // the same as validator withdraw
            }
        }
    }
}

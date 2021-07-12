// SPDX-License-Identifier: MIT
pragma solidity <=0.8.6;

import "./System.sol";
import "./Interface/IValidator.sol";
import "./Interface/IBond.sol";
import "./ValidatorPool.sol";
import "./SystemReward.sol";

/*[
    ["0x70F657164e5b75689b64B7fd1fA275F334f28e18","35481719561823448621","BondStatus.UNBONDED",63306535440133,false],
    ["0x2465176C461AfB316ebc773C61fAEe85A6515DAA", "35481719561823448621", "BondStatus.UNBONDED",  64328738321847, false]
  ]
*/
/*
  [
    ["0x70F657164e5b75689b64B7fd1fA275F334f28e18",63306535440133],
    ["0x2465176C461AfB316ebc773C61fAEe85A6515DAA",64328738321847]
  ]
*/

contract BKCValidatorSet is System, IValidator {
    Validator[] public currentValidatorSet; // the array holding the validator structs
    mapping(address => uint256) public currentValidatorSetMap; // map validator's consensus address to the position in the array

    uint256 public number_of_validators;

    mapping(uint256 => bool) private alreadyIn;
    uint256[] private powers;

    Validator[] private tmpValidatorSet;
    uint256 public endTime;

    address private _validatorPoolAddress;
    address private _systemRewardAddress;

    ValidatorPool private vldpool;
    SystemReward private systemreward;

    modifier onlyAfterEndTime() {
        require(block.timestamp >= endTime, "can't update before end time");
        _;
    }

    modifier onlyValidatorInValidatorSet(address consensusAddress) {
        require(
            currentValidatorSetMap[consensusAddress] > 0,
            "not allow for non-active validator"
        );
        _;
    }

    modifier onlyValidatorPoolContract {
        require(msg.sender == _validatorPoolAddress);
        _;
    }

    function updateEndTime() private {
        endTime += unbondingPeriod;
    }

    function init(address validatorPoolAddress, address systemRewardAddress)
        external
        onlyNotInit
    {
        endTime = block.timestamp;
        _validatorPoolAddress = validatorPoolAddress;
        _systemRewardAddress = systemRewardAddress;
        vldpool = ValidatorPool(validatorPoolAddress);
        systemreward = SystemReward(systemRewardAddress);
        updateValidatorSet();
        endTime = block.timestamp + unbondingPeriod;
        alreadyInit = true;
    }

    function getValidators() public view onlyInit returns (Validator[] memory) {
        return currentValidatorSet;
    }

    function removeValidator(address consensusAddress) private onlyInit {
        uint256 index = currentValidatorSetMap[consensusAddress] - 1;
        for (uint256 i = index; i < number_of_validators - 1; i++) {
            currentValidatorSet[i] = currentValidatorSet[i + 1]; // shift left
            Validator memory v = currentValidatorSet[i];
            currentValidatorSetMap[v.consensusAddress] = i + 1;
        }
        currentValidatorSet.pop();
        delete currentValidatorSetMap[consensusAddress];
        number_of_validators -= 1;
    }

    function jailValidator(address consensusAddress)
        external
        onlyInit
        onlyValidatorPoolContract
    {
        removeValidator(consensusAddress);
        systemreward.AllocateJailValidatorReward(consensusAddress);
    }

    // need to add onlyAfterEndTime here as modifier
    function updateValidatorSet() public {

            uint256 _number_of_validators
         = MAX_NUMBER_OF_VALIDATORS_IN_VALIDATOR_SET;
        if (
            vldpool.NumberOfValidator() <
            MAX_NUMBER_OF_VALIDATORS_IN_VALIDATOR_SET
        ) {
            _number_of_validators = vldpool.NumberOfValidator();
        }
        // uint256 _number_of_validators = vldpool.NumberOfValidator() >
        //     MAX_NUMBER_OF_VALIDATORS_IN_VALIDATOR_SET
        //     ? MAX_NUMBER_OF_VALIDATORS_IN_VALIDATOR_SET
        //     : vldpool.NumberOfValidator();
        for (uint256 i = 0; i < currentValidatorSet.length; i++) {
            Validator memory validator = currentValidatorSet[i];
            vldpool.unBondValidator(validator.consensusAddress); // change the bond state to UNBONDED
            delete currentValidatorSetMap[validator.consensusAddress];
        }

        delete currentValidatorSet;

        delete powers;

        number_of_validators = _number_of_validators;

        uint256 len = vldpool.NumberOfValidator();
        for (uint256 j = 0; j < len; j++) {
            (address a, , , ) = vldpool.validators(j);
            uint256 power = vldpool.getTotalPowerExcludeUnbonding(a);
            powers.push(power);
            alreadyIn[j] = false;
        }
        for (uint256 i = 0; i < number_of_validators; i++) {
            uint256 max_power = 0;
            uint256 max_index = 0;
            for (uint256 j = 0; j < len; j++) {
                (address a, , , bool Jail) = vldpool.validators(j);
                uint256 time = vldpool.validatorRemoveQueue(a); // time is 0 not in queue
                if (
                    powers[j] > max_power && !alreadyIn[j] && !Jail && time == 0
                ) {
                    max_power = powers[j];
                    max_index = j;
                }
            }
            alreadyIn[max_index] = true;
            (address a, uint256 stake, BondStatus b, bool isJail) = vldpool
            .validators(max_index);
            vldpool.bondValidator(a); // bond the validator
            (, , b, ) = vldpool.validators(max_index);
            currentValidatorSet.push(Validator(a, stake, b, isJail));
            currentValidatorSetMap[a] = i + 1;
        }

        systemreward.distributeReward(); // distribute reward

        // uint j = 0;

        // for(uint i = 0; i < number_of_validators; i++){
        //     (address a, uint stake, BondStatus b, bool isJail) = vldpool.validators(i);
        //     if (!isJail){
        //         Validator memory v = Validator(a,stake,b,isJail);
        //         currentValidatorSet.push(v);
        //         vldpool.bondValidator(v.consensusAddress); // change the bond state to BONDED
        //         currentValidatorSetMap[v.consensusAddress] = j+1;
        //         j++;
        //     }
        // }
        updateEndTime(); // for the next round.
    }
}

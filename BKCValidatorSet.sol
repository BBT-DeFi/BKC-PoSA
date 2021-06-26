pragma solidity <=0.8.6;

import "./System.sol";
import "./Interface/IValidator.sol";
import "./Interface/IBond.sol";
import "./ValidatorPool.sol";

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
    mapping(address => uint256) public currentValidatorSetMap; // map validator's concensus address to the position in the array

    Validator[] private tmpValidatorSet;

    uint256 public endTime;

    ValidatorPool private vldpool;

    modifier onlyAfterEndTime() {
        require(block.timestamp >= endTime);
        _;
    }

    modifier onlyValidatorInValidatorSet(address concensusAddress) {
        require(
            currentValidatorSetMap[concensusAddress] > 0,
            "not allow for non-active validator"
        );
        _;
    }

    function updateEndTime() private {
        endTime += unbondingPeriod;
    }

    function init(address validatorPoolAddress) external onlyNotInit {
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

        tmpValidatorSet.push(firstValidator);
        tmpValidatorSet.push(secondValidator);

        for (uint256 i = 0; i < tmpValidatorSet.length; i++) {
            currentValidatorSet.push(tmpValidatorSet[i]);
            currentValidatorSetMap[tmpValidatorSet[i].consensusAddress] = i + 1;
        }
        vldpool = ValidatorPool(validatorPoolAddress);
        endTime = block.timestamp + unbondingPeriod;
        alreadyInit = true;
    }

    function getValidators() public view onlyInit returns (Validator[] memory) {
        return currentValidatorSet;
    }

    /* TODO : need to change from inputting newValidatorSet into 
    calculating the top validators in validator pool contract
    */
    function updateValidatorSet(Validator[] memory newValidatorSet)
        public
        onlyInit
        onlyAfterEndTime
        onlyValidatorInValidatorSet(msg.sender)
    {
        for (uint256 i = 0; i < currentValidatorSet.length; i++) {
            Validator memory validator = currentValidatorSet[i];
            vldpool.unBondValidator(validator.consensusAddress); // change the bond state to UNBONDED
            delete currentValidatorSetMap[validator.consensusAddress];
        }
        delete currentValidatorSet;
        for (uint256 i = 0; i < newValidatorSet.length; i++) {
            currentValidatorSet.push(newValidatorSet[i]);
            vldpool.bondValidator(newValidatorSet[i].consensusAddress); // change the bond state to BONDED
            currentValidatorSetMap[newValidatorSet[i].consensusAddress] = i + 1;
        }
        updateEndTime(); // for the next round.
    }
}

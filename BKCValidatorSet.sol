pragma solidity <=0.8.6;

import "./System.sol";
import "./Interface/IValidator.sol";

/*[
    ["0x70F657164e5b75689b64B7fd1fA275F334f28e18","0x5CF5d69852151952b1b1faF04d9B6140dAe30aFD","0x022180D9cF057A1d4bda8AeF1c581F0F17c1745c",63306535440133,false],
    ["0x2465176C461AfB316ebc773C61fAEe85A6515DAA", "0x14cC4EcB0DA0B7b0281B36081912A6AF07F779e9", "0x99CF09e8f5F8Dc3465c834fbf9154C2050075203",  64328738321847, false]
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

    function init() external override onlyNotInit {
        Validator memory firstValidator = Validator(
            0x295e26495CEF6F69dFA69911d9D8e4F3bBadB89B,
            // payable(0x812625152e78C2a7027B9767cC5ab6Fc4B3D2e28),
            // 0xF5cACa6d5fCb6F819147892191716cf443D224be,
            // 64619916293869,
            // false
            48648428258251948625,
            false,
            false
        );
        tmpValidatorSet.push(firstValidator);

        for (uint256 i = 0; i < tmpValidatorSet.length; i++) {
            currentValidatorSet.push(tmpValidatorSet[i]);
            currentValidatorSetMap[tmpValidatorSet[i].consensusAddress] = i + 1;
        }
        alreadyInit = true;
    }

    function getValidators() public view onlyInit returns (Validator[] memory) {
        return currentValidatorSet;
    }

    function updateValidatorSet(Validator[] memory newValidatorSet)
        public
        onlyInit
    {
        for (uint256 i = 0; i < currentValidatorSet.length; i++) {
            Validator memory validator = currentValidatorSet[i];
            delete currentValidatorSetMap[validator.consensusAddress];
        }
        delete currentValidatorSet;
        for (uint256 i = 0; i < newValidatorSet.length; i++) {
            currentValidatorSet.push(newValidatorSet[i]);
            currentValidatorSetMap[newValidatorSet[i].consensusAddress] = i + 1;
        }
    }
}

// SPDX-License-Identifier: MIT
pragma solidity <=0.8.6;

import "./IBond.sol";

abstract contract IDelegation {
    struct DelegationObject {
        mapping(address => uint256) delegateAmountOfEach;
        address[] delegators;
        uint256 totalDelegation;
    }

    struct userDelegation {
        address[] validators;
        mapping(address => uint256) bondedAmountForEach;
        mapping(address => uint256) unbondingAmountForEach;
        // address consensusAddress;
        // uint256 bondedAmount;
        // uint256 unbondingAmount;
    }
}

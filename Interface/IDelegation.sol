pragma solidity <=0.8.6;

import "./IBond.sol";

abstract contract IDelegation {
    struct DelegationObject {
        mapping(address => uint256) delegateAmountOfEach;
        uint256 totalDelegation;
    }

    struct userDelegation {
        address consensusAddress;
        uint256 bondedAmount;
        uint256 unbondingAmount;
    }
}

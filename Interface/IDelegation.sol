pragma solidity <=0.8.6;

abstract contract IDelegation {
    struct DelegationObject {
        mapping(address => uint256) delegateAmountOfEach;
        uint256 totalDelegation;
    }
}

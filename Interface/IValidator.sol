pragma solidity <=0.8.6;

abstract contract IValidator {
    struct Validator {
        address consensusAddress;
        uint256 stakeAmount;
        bool isBond;
        bool isJail;
        // address payable feeAddress;
        // address BBCFeeAddress;
        // uint64  votingPower;

        // // only in state
        // bool jailed;
        // uint256 incoming;
    }
}

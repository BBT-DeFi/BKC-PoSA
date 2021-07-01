pragma solidity <=0.8.6;

import "./IBond.sol";

abstract contract IValidator is IBond {
    struct Validator {
        address consensusAddress;
        uint256 stakeAmount;
        BondStatus bondStatus;
        bool isJail;
        // address payable feeAddress;
        // address BBCFeeAddress;
        // uint64  votingPower;

        // // only in state
        // bool jailed;
        // uint256 incoming;
    }

    uint256 constant MAX_NUMBER_OF_VALIDATORS_IN_VALIDATOR_SET = 2;

    uint256 constant VALIDATOR_JAIL_PERIOD = 10 seconds;
}

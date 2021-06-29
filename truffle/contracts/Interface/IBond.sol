pragma solidity <=0.8.6;

abstract contract IBond {
    enum BondStatus {
        BONDED,
        UNBONDING,
        UNBONDED
    }

    uint256 unbondingPeriod = 1 days;
}

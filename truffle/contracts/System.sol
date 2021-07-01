// SPDX-License-Identifier: MIT
pragma solidity <=0.8.6;

abstract contract System {
    bool public alreadyInit;
    modifier onlyCoinbase() {
        require(
            msg.sender == block.coinbase,
            "the message sender must be the block producer"
        );
        _;
    }

    modifier onlyNotInit() {
        require(!alreadyInit, "the contract already init");
        _;
    }

    modifier onlyInit() {
        require(alreadyInit, "the contract not init yet");
        _;
    }

    // function init() external virtual ;
}

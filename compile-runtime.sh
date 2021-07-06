mkdir build-runtime
solc --abi --bin-runtime ./Contracts/BKCValidatorSet.sol -o build-runtime
solc --abi --bin-runtime ./Contracts/StakePool.sol -o build-runtime
solc --abi --bin-runtime ./Contracts/SystemReward.sol -o build-runtime
solc --abi --bin-runtime ./Contracts/ValidatorPool.sol -o build-runtime

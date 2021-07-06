mkdir build
solc --abi --bin ./Contracts/BKCValidatorSet.sol -o build
solc --abi --bin ./Contracts/StakePool.sol -o build
solc --abi --bin ./Contracts/SystemReward.sol -o build
solc --abi --bin ./Contracts/ValidatorPool.sol -o build

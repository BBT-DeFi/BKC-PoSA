mkdir build

solc --abi --bin --overwrite ./contracts/BKCValidatorSet.sol -o build
solc --abi --bin --overwrite ./contracts/StakePool.sol -o build
solc --abi --bin --overwrite ./contracts/SystemReward.sol -o build
solc --abi --bin --overwrite ./contracts/ValidatorPoool.sol -o build

mkdir abi
mkdir abi/validatorset
mkdir abi/stakepool
mkdir abi/systemreward
mkdir abi/vldpool

abigen --bin=./build/BKCValidatorSet.bin --abi=./build/BKCValidatorSet.abi --pkg=validatorset --out=abi/validatorset/BKCValidatorSet.go
abigen --bin=./build/StakePool.bin --abi=./build/StakePool.abi --pkg=stakepool --out=abi/stakepool/StakePool.go
abigen --bin=./build/SystemReward.bin --abi=./build/SystemReward.abi --pkg=systemreward --out=abi/systemreward/SystemReward.go
abigen --bin=./build/ValidatorPool.bin --abi=./build/ValidatorPool.abi --pkg=vldpool --out=abi/vldpool/ValidatorPool.go
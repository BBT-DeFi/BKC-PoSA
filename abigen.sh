mkdir posa/BKCValidatorSet
mkdir posa/StakePool
mkdir posa/SystemReward
mkdir posa/ValidatorPool
abigen --abi=./build/BKCValidatorSet.abi --bin=./build/BKCValidatorSet.bin --pkg=validatorset --out=posa/BKCValidatorSet/BKCValidatorSet.go
abigen --abi=./build/StakePool.abi --bin=./build/StakePool.bin --pkg=stakepool --out=posa/StakePool/StakePool.go
abigen --abi=./build/SystemReward.abi --bin=./build/SystemReward.bin --pkg=systemreward --out=posa/SystemReward/SystemReward.go
abigen --abi=./build/ValidatorPool.abi --bin=./build/ValidatorPool.bin --pkg=validatorpool --out=posa/ValidatorPool/ValidatorPool.go

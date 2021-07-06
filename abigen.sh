mkdir posa/BKCValidatorSet
mkdir posa/StakePool
mkdir posa/SystemReward
mkdir posa/ValidatorPool
abigen --abi=./build/BKCValidatorSet.abi --pkg=validatorset --out=posa/BKCValidatorSet/BKCValidatorSet.go
abigen --abi=./build/StakePool.abi --pkg=stakepool --out=posa/StakePool/StakePool.go
abigen --abi=./build/SystemReward.abi --pkg=systemreward --out=posa/SystemReward/SystemReward.go
abigen --abi=./build/ValidatorPool.abi --pkg=validatorpool --out=posa/ValidatorPool/ValidatorPool.go

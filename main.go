package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	validatorset "ethclient/posa/BKCValidatorSet"
	stakepool "ethclient/posa/StakePool"
	systemreward "ethclient/posa/SystemReward"
	validatorpool "ethclient/posa/ValidatorPool"
)

func main() {
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Fatal(err)
	}

	validatorsetAddress := "0x61a75be699e3B21488aFae66f1133A06301F5B96"
	stakepoolAddress := "0x46629fa29d04227b422739aD26f4c559fe1EfE14"
	systemrewardAddress := "0xe6BDc0d5ff96A4842AC209c0c1Ebe3FE53db818b"
	validatorpoolAddress := "0x0F750D25B5d791bE6896D6c56689E1DD9c601F8b"

	bkcaddress := common.HexToAddress(validatorsetAddress)
	BKCValidatorSetInstance, err := validatorset.NewValidatorset(bkcaddress, client)
	if err != nil {
		log.Fatal(err)
	}

	stakepooladdress := common.HexToAddress(stakepoolAddress)
	StakePoolInstance, err := stakepool.NewStakepool(stakepooladdress, client)
	if err != nil {
		log.Fatal(err)
	}

	systemrewardaddress := common.HexToAddress(systemrewardAddress)
	SystemRewardInstance, err := systemreward.NewSystemreward(systemrewardaddress, client)
	if err != nil {
		log.Fatal(err)
	}

	validatorpooladdress := common.HexToAddress(validatorpoolAddress)
	ValidatorPoolInstance, err := validatorpool.NewValidatorpool(validatorpooladdress, client)

	ValidatorPoolInstance.Init(&bind.TransactOpts{}, bkcaddress, stakepooladdress)
	StakePoolInstance.Init(&bind.TransactOpts{}, validatorpooladdress)
	SystemRewardInstance.Init(&bind.TransactOpts{}, bkcaddress, stakepooladdress, validatorpooladdress)
	BKCValidatorSetInstance.Init(&bind.TransactOpts{}, validatorpooladdress, systemrewardaddress)

	alreadyInit1, err := BKCValidatorSetInstance.AlreadyInit(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	alreadyInit2, err := StakePoolInstance.AlreadyInit(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	alreadyInit3, err := SystemRewardInstance.AlreadyInit(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	alreadyInit4, err := ValidatorPoolInstance.AlreadyInit(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(alreadyInit1)
	fmt.Println(alreadyInit2)
	fmt.Println(alreadyInit3)
	fmt.Println(alreadyInit4)

}

// func loadContracts(){

// }

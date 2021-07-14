/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"math/big"
	"time"
	IValidator "win/Code/IValidator"
	Init "win/Code/Init"
	vldpool "win/abi/vldpool"
	ETHclient "win/client"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

// vldpoolCmd represents the vldpool command
var vldpoolCmd = &cobra.Command{
	Use:   "vldpool",
	Short: "query the validator pool contract",
	Long:  `This contract consits of validators and their stake amount`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := ETHclient.Client()
		handleError(err)

		name := "validator pool"

		//ValidatorPoolAddress := "0x282659B28f9acCaC9fBd81317Ce43b4044E5d4A7"

		ValidatorPoolInstance, err := vldpool.NewVldpool(common.HexToAddress(ValidatorPoolAddress), client)
		handleError(err)

		val, err := cmd.Flags().GetBool("alreadyInit")
		handleError(err)
		num, err := cmd.Flags().GetBool("numberOfValidators")
		handleError(err)
		v, err := cmd.Flags().GetBool("getAllValidators")
		handleError(err)
		m, err := cmd.Flags().GetString("validatorsMap")
		handleError(err)
		addr, err := cmd.Flags().GetString("validatorUnbondQueue")
		handleError(err)
		addr2, err := cmd.Flags().GetString("validatorJailQueue")
		handleError(err)
		addr3, err := cmd.Flags().GetString("validatorRemoveQueue")
		handleError(err)
		idx, err := cmd.Flags().GetInt("getValidator")
		handleError(err)
		addr4, err := cmd.Flags().GetString("totalPower")
		handleError(err)
		if val {
			Init.AlreadyInit(*client, ValidatorPoolInstance, name)
		} else if idx != -1 {
			GetValidator(ValidatorPoolInstance, idx)
		} else if num {
			n := GetNumberOfValdiatorInPool(ValidatorPoolInstance)
			fmt.Println()
			fmt.Println("There are", n, "validators in the pool.")
			fmt.Println()
		} else if v {
			GetAllValidators(ValidatorPoolInstance)
		} else if m != "" {
			GetValidatorsMap(ValidatorPoolInstance, m)
		} else if addr != "" {
			GetVaidatorUnbondQueue(ValidatorPoolInstance, addr)
		} else if addr2 != "" {
			GetValidatorUnJailQueue(ValidatorPoolInstance, addr2)
		} else if addr3 != "" {
			GetValidatorRemoveQueue(ValidatorPoolInstance, addr3)
		} else if addr4 != "" {
			GetTotalPowerExcludeUnbonding(ValidatorPoolInstance, addr4)
		} else {
			fmt.Println("please specify the field you want to query")
		}
	},
}

func init() {
	rootCmd.AddCommand(vldpoolCmd)
	vldpoolCmd.Flags().BoolVarP(&alreadyInit, "alreadyInit", "i", false, "the init status")

	vldpoolCmd.Flags().BoolVarP(&number_of_validators, "numberOfValidators", "n", false, "the number of validators in the pool")

	vldpoolCmd.Flags().IntVarP(&index, "getValidator", "a", -1, "get the validator at the specified index (specify the index starts at 0 for first validator)")

	vldpoolCmd.Flags().BoolVarP(&getValidators, "getAllValidators", "v", false, "All the validators in the pool")

	vldpoolCmd.Flags().StringVarP(&validatormap, "validatorsMap", "m", "", "The map of the validators in the pool (specify validator address)")

	vldpoolCmd.Flags().StringVarP(&consensusAddress, "validatorUnbondQueue", "u", "", "The unbond time of the validator (specify validator address")

	vldpoolCmd.Flags().StringVarP(&consensusAddress2, "validatorJailQueue", "j", "", "The unjail time of the validator (specify validator address)")

	vldpoolCmd.Flags().StringVarP(&consensusAddress3, "validatorRemoveQueue", "r", "", "The remove time of the validator (specify validator address)")

	vldpoolCmd.Flags().StringVarP(&consensusAddress4, "totalPower", "p", "", "the total power (exclude unbonding delegation) of a validator (specify validator address)")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// vldpoolCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// vldpoolCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func GetNumberOfValdiatorInPool(instance *vldpool.Vldpool) int {
	n, err := instance.NumberOfValidator(&bind.CallOpts{})
	handleError(err)
	return int(n.Int64())
}

func GetAllValidators(instance *vldpool.Vldpool) {
	n := GetNumberOfValdiatorInPool(instance)
	for i := 0; i < n; i++ {
		validator, err := instance.Validators(&bind.CallOpts{}, big.NewInt(int64(i)))
		handleError(err)
		fmt.Println()
		fmt.Println("Validator", i+1)
		IValidator.PrintValidator(IValidator.Validator(validator))
		fmt.Println()
	}
}

func GetValidatorsMap(instance *vldpool.Vldpool, m string) {
	valmap, err := instance.ValidatorsMap(&bind.CallOpts{}, common.HexToAddress(m))
	handleError(err)
	if valmap.Int64() == int64(0) {
		fmt.Println()
		fmt.Println("The validator address", m, "is not in the validator pool")
		fmt.Println()
		return
	}
	fmt.Println()
	fmt.Println("The validator", m, "is at index", valmap, "(index starts at 1)")
	fmt.Println()
}

func GetVaidatorUnbondQueue(instance *vldpool.Vldpool, addr string) {
	unbond, err := instance.ValidatorUnBondQueue(&bind.CallOpts{}, common.HexToAddress(addr))
	handleError(err)
	fmt.Println()
	if unbond.Int64() != int64(0) {
		fmt.Println("The validator", addr, "will finish unbond at", time.Unix(int64(unbond.Uint64()), 0))
	} else {
		fmt.Println("The validator", addr, "isn't unbonding or the address is not a validator >> use the command vldpool -m <validator address> to check for index of the validator.")
	}
	fmt.Println()
}

func GetValidatorUnJailQueue(instance *vldpool.Vldpool, addr2 string) {
	unjail, err := instance.ValidatorJailQueue(&bind.CallOpts{}, common.HexToAddress(addr2))
	handleError(err)
	fmt.Println()
	if unjail.Int64() != int64(0) {
		fmt.Println("The validator", addr2, "will be able to unjail at", time.Unix(int64(unjail.Uint64()), 0))
	} else {
		fmt.Println("The validator", addr2, "is not in jailed state, or the address provided is not a validator >> use the command vldpool -m <validator address> to check for index of the validator.")
	}
	fmt.Println()
}

func GetValidatorRemoveQueue(instance *vldpool.Vldpool, addr3 string) {
	remove, err := instance.ValidatorRemoveQueue(&bind.CallOpts{}, common.HexToAddress(addr3))
	handleError(err)
	fmt.Println()
	if remove.Int64() != int64(0) {
		fmt.Println("The validator", addr3, "can remove itself from the pool at", time.Unix(int64(remove.Uint64()), 0))
	} else {
		fmt.Println("The validator", addr3, "isn't in the remove queue or the address is not a validator >> use the command vldpool -m <validator address> to check for index of the validator.")
	}
	fmt.Println()
}

func GetValidator(instance *vldpool.Vldpool, idx int) {
	validator, err := instance.Validators(&bind.CallOpts{}, big.NewInt(int64(idx)))
	handleError(err)
	fmt.Println()
	fmt.Println("Validator at index", idx, "in the pool")
	IValidator.PrintValidator(IValidator.Validator(validator))
	fmt.Println()
}

func GetTotalPowerExcludeUnbonding(instance *vldpool.Vldpool, addr4 string) {
	power, err := instance.GetTotalPowerExcludeUnbonding(&bind.CallOpts{}, common.HexToAddress(addr4))
	handleError(err)
	fmt.Println()
	fmt.Println("The total power of the validator", addr4, "is", power)
	fmt.Println()
}

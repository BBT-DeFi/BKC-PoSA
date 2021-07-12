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
	"log"
	"math/big"
	"time"
	IValidator "win/Code/IValidator"
	Init "win/Code/Init"
	"win/abi/validatorset"
	ETHclient "win/client"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

var alreadyInit bool
var index int
var consensusAddress string
var number_of_validators bool
var endTime bool
var validators bool

// validatorsetCmd represents the validatorset command
var validatorsetCmd = &cobra.Command{
	Use:   "validatorset",
	Short: "query the validator set contract",
	Long: `The validator set contract contains the active validator set as well as the functions for updating 
	the new validator set`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := ETHclient.Client()
		if err != nil {
			log.Fatal(err)
		}

		name := "validator set"

		//BKCValidatorSetAddress := "0x7eA14c6696EB86a9c7C7a8aCbaC4Bd7BFa80F974"

		BKCValidatorSetInstance, err := validatorset.NewValidatorset(common.HexToAddress(BKCValidatorSetAddress), client)

		in, err := cmd.Flags().GetBool("alreadyInit")
		cu, err := cmd.Flags().GetInt("currentValidatorSet")
		cumap, err := cmd.Flags().GetString("currentValidatorSetMap")
		num, err := cmd.Flags().GetBool("numberOfValidators")
		end, err := cmd.Flags().GetBool("endTime")
		vl, err := cmd.Flags().GetBool("validators")
		if err != nil {
			log.Fatal(err)
		}
		if in {
			Init.AlreadyInit(*client, BKCValidatorSetInstance, name)
		} else if cu >= 0 {
			GetValidatorInSet(BKCValidatorSetInstance, cu)
		} else if cumap != "" {
			GetValidatorSetMap(BKCValidatorSetInstance, cumap)
		} else if num {
			GetNumberOfValdiator(BKCValidatorSetInstance)
		} else if end {
			GetEndTime(BKCValidatorSetInstance)
		} else if vl {
			vlds, err := BKCValidatorSetInstance.GetValidators(&bind.CallOpts{})
			if err != nil {
				log.Fatal(err)
			}
			for i, vld := range vlds {
				fmt.Println()
				fmt.Println("active validator", i+1)
				IValidator.PrintValidator(IValidator.Validator(vld))
				fmt.Println()
			}
		} else {
			fmt.Println("please specify the field you want to query")
		}
	},
}

func init() {
	rootCmd.AddCommand(validatorsetCmd)
	validatorsetCmd.Flags().BoolVarP(&alreadyInit, "alreadyInit", "a", false, "the init status")

	validatorsetCmd.Flags().IntVarP(&index, "currentValidatorSet", "s", -1, "the current validator set (specify index)")

	validatorsetCmd.Flags().StringVarP(&consensusAddress, "currentValidatorSetMap", "m", "", "the map of the current validator set (specify address)")

	validatorsetCmd.Flags().BoolVarP(&number_of_validators, "numberOfValidators", "n", false, "the number of validators")

	validatorsetCmd.Flags().BoolVarP(&endTime, "endTime", "e", false, "the end time of each epoch")

	validatorsetCmd.Flags().BoolVarP(&validators, "validators", "v", false, "the validators in active set")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// validatorsetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// validatorsetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func PrintValue(instance *validatorset.Validatorset, name string, value interface{}) {

	fmt.Println("The contract's", name, "field is", value)
}

func GetValidatorInSet(instance *validatorset.Validatorset, i int) {

	set, err := instance.CurrentValidatorSet(&bind.CallOpts{}, big.NewInt(int64(i)))
	set = IValidator.Validator(set)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Validator in validator set at index", i)
	IValidator.PrintValidator(set)
}

func GetValidatorSetMap(instance *validatorset.Validatorset, addr string) {
	idx, err := instance.CurrentValidatorSetMap(&bind.CallOpts{}, common.HexToAddress(addr))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The index (starts at 1) of the validator", addr, "is", idx)
}

func GetNumberOfValdiator(instance *validatorset.Validatorset) {
	n, err := instance.NumberOfValidators(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("There are", n, "validators in the active set")
}

func GetEndTime(instance *validatorset.Validatorset) {
	endtime, err := instance.EndTime(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	et := time.Unix(int64(endtime.Uint64()), 0)
	fmt.Println("The end time for this epoch is", et)
}

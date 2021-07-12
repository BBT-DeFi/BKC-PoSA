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
	Init "win/Code/Init"
	"win/abi/stakepool"
	ETHclient "win/client"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

var delegatorAddress string
var consensusAddress2 string

// stakepoolCmd represents the stakepool command
var stakepoolCmd = &cobra.Command{
	Use:   "stakepool",
	Short: "query the stake pool contract",
	Long:  `This contract consists of staking module and delegation logics`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := ETHclient.Client()
		if err != nil {
			log.Fatal(err)
		}

		name := "stake pool"

		//StakePoolAddress := "0xD825E46b1f610Aa96e6A9454aD99B1CA321B2751"

		StakePoolInstance, err := stakepool.NewStakepool(common.HexToAddress(StakePoolAddress), client)

		val, err := cmd.Flags().GetBool("alreadyInit")
		addrs, err := cmd.Flags().GetStringSlice("ValidatorDelegation")
		addr, err := cmd.Flags().GetString("totalDelegation")
		addr2, err := cmd.Flags().GetString("Delegators")
		if val {
			Init.AlreadyInit(*client, StakePoolInstance, name)
		} else if addrs[0] != "" && addrs[1] != "" {
			GetDelegationAmountOfEach(StakePoolInstance, addrs)
		} else if addr != "" {
			GetTotalDelegation(StakePoolInstance, addr)
		} else if addr2 != "" {
			GetDelegators(StakePoolInstance, addr2)
		} else {
			fmt.Println("please specify the field you want to query")
		}
	},
}

func init() {
	rootCmd.AddCommand(stakepoolCmd)
	stakepoolCmd.Flags().BoolVarP(&alreadyInit, "alreadyInit", "i", false, "the init status")

	stakepoolCmd.Flags().StringSliceP("ValidatorDelegation", "v", []string{"", ""}, "The delegation object of the validator (specify validator address then delegator address separated by a comma)")

	stakepoolCmd.Flags().StringVarP(&consensusAddress, "totalDelegation", "t", "", "the total delegation of a validator (specify address)")

	stakepoolCmd.Flags().StringVarP(&consensusAddress2, "Delegators", "d", "", "the list of delegators of this validator (specify address)")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stakepoolCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stakepoolCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func GetDelegationAmountOfEach(instance *stakepool.Stakepool, addr []string) {
	deleg, err := instance.GetDelegationAmountOfEach(&bind.CallOpts{}, common.HexToAddress(addr[0]), common.HexToAddress(addr[1]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The total delegation of delegator", addr[1], "for validator", addr[0], "is", deleg)
}

func GetTotalDelegation(instance *stakepool.Stakepool, addr string) {
	tt, err := instance.GetTotalDelegation(&bind.CallOpts{}, common.HexToAddress(addr))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The total delegation of the validator", addr, "is", tt)
}

func GetDelegators(instance *stakepool.Stakepool, addr2 string) {
	delegators, err := instance.GetDelegators(&bind.CallOpts{}, common.HexToAddress(addr2))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("for validator:", addr2)
	for i, d := range delegators {
		fmt.Println("delegator", i, ":", d)
	}
}

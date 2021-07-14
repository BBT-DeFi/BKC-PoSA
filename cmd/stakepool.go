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
	"strconv"
	"time"
	Init "win/Code/Init"
	"win/abi/stakepool"
	ETHclient "win/client"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

// stakepoolCmd represents the stakepool command
var stakepoolCmd = &cobra.Command{
	Use:   "stakepool",
	Short: "query the stake pool contract",
	Long:  `This contract consists of staking module and delegation logics`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := ETHclient.Client()
		handleError(err)

		name := "stake pool"

		//StakePoolAddress := "0xD825E46b1f610Aa96e6A9454aD99B1CA321B2751"

		StakePoolInstance, err := stakepool.NewStakepool(common.HexToAddress(StakePoolAddress), client)
		handleError(err)

		val, err := cmd.Flags().GetBool("alreadyInit")
		handleError(err)
		addrs, err := cmd.Flags().GetStringSlice("ValidatorDelegation")
		handleError(err)
		addr, err := cmd.Flags().GetString("totalDelegation")
		handleError(err)
		addr2, err := cmd.Flags().GetString("Delegators")
		handleError(err)
		addrs2, err := cmd.Flags().GetStringSlice("UserDelegationBonded")
		handleError(err)
		addrs3, err := cmd.Flags().GetStringSlice("UserDelegationUnbonding")
		handleError(err)
		e, err := cmd.Flags().GetBool("ExcludeUnbonding")
		handleError(err)
		deleg, err := cmd.Flags().GetString("getUnbondQueueValue")
		handleError(err)
		if val {
			Init.AlreadyInit(*client, StakePoolInstance, name)
		} else if addrs[0] != "" && addrs[1] != "" {
			GetDelegationAmountOfEach(StakePoolInstance, addrs)
		} else if addr != "" {
			if !e {
				GetTotalDelegation(StakePoolInstance, addr)
			} else {
				GetTotalDelegationExcludeUnbonding(StakePoolInstance, addr)
			}
		} else if addr2 != "" {
			GetDelegators(StakePoolInstance, addr2)
		} else if addrs2[0] != "" && addrs2[1] != "" {
			GetUserDelegationBondedAmount(StakePoolInstance, addrs2)
		} else if addrs3[0] != "" && addrs3[1] != "" {
			GetUserDelegationUnbondingAmount(StakePoolInstance, addrs3)
		} else if deleg != "" {
			GetUserUnDelegateValue(StakePoolInstance, deleg)
		} else {
			fmt.Println("please specify the field you want to query")
		}
	},
}

func init() {
	rootCmd.AddCommand(stakepoolCmd)
	stakepoolCmd.Flags().BoolVarP(&alreadyInit, "alreadyInit", "i", false, "the init status")

	stakepoolCmd.Flags().StringSliceP("ValidatorDelegation", "v", []string{"", ""}, "The delegate amount of each of delegator for a validator (specify validator address then delegator address separated by a comma)")

	stakepoolCmd.Flags().StringVarP(&consensusAddress, "totalDelegation", "t", "", "the total delegation of a validator (specify address)")

	stakepoolCmd.Flags().StringVarP(&consensusAddress2, "Delegators", "d", "", "the list of delegators of this validator (specify address)")

	stakepoolCmd.Flags().BoolVarP(&exclude, "ExcludeUnbonding", "e", false, "the flag to add to totalDelegation flag to get total delegation exclude unbonding amounts")

	stakepoolCmd.Flags().StringSliceP("UserDelegationBonded", "b", []string{"", ""}, "The bonded amount of delegator for the validator (specify validator address then delegator address separated by a comma)")

	stakepoolCmd.Flags().StringSliceP("UserDelegationUnbonding", "u", []string{"", ""}, "The unbonding amount of delegator for validator (specify validator address then delegator address separated by a comma)")

	stakepoolCmd.Flags().StringVarP(&delegatorAddress, "getUnbondQueueValue", "q", "", "The unbonding value in the queue of this delegator (specify the delegator address)")
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
	handleError(err)
	fmt.Println()
	fmt.Println("The total delegation of delegator", addr[1], "for validator", addr[0], "is", deleg)
	fmt.Println()
}

func GetTotalDelegation(instance *stakepool.Stakepool, addr string) {
	tt, err := instance.GetTotalDelegation(&bind.CallOpts{}, common.HexToAddress(addr))
	handleError(err)
	fmt.Println()
	fmt.Println("The total delegation of the validator", addr, "is", tt)
	fmt.Println()
}

func GetTotalDelegationExcludeUnbonding(instance *stakepool.Stakepool, addr string) {
	totalExclude, err := instance.GetTotalDelegationExcludeUnbonding(&bind.CallOpts{}, common.HexToAddress(addr))
	handleError(err)
	fmt.Println()
	fmt.Println("The total delegation exclude unbonding amount of the validator", addr, "is", totalExclude)
	fmt.Println()
}

func GetUserDelegationBondedAmount(instance *stakepool.Stakepool, addrs2 []string) {
	b, err := instance.GetUserDelegationBondedAmountCallable(&bind.CallOpts{}, common.HexToAddress(addrs2[1]), common.HexToAddress(addrs2[0]))
	handleError(err)
	fmt.Println()
	fmt.Println("The bonded amount of delegator", addrs2[1], "for validator", addrs2[0], "is", b)
	fmt.Println()
}

func GetUserDelegationUnbondingAmount(instance *stakepool.Stakepool, addrs3 []string) {
	u, err := instance.GetUserDelegationUnbondingAmountCallable(&bind.CallOpts{}, common.HexToAddress(addrs3[1]), common.HexToAddress(addrs3[0]))
	handleError(err)
	fmt.Println()
	fmt.Println("The unbonding amount of delegator", addrs3[1], "for validator", addrs3[0], "is", u)
	fmt.Println()
}

func GetUserUnDelegateValue(instance *stakepool.Stakepool, deleg string) {
	q, err := instance.GetUnbondingValue(&bind.CallOpts{}, common.HexToAddress(deleg))
	handleError(err)
	fmt.Println()
	fmt.Println("for delegator", deleg)
	for i, qval := range q {
		fmt.Println()
		fmt.Println("undelegate no." + strconv.Itoa(i+1))
		fmt.Println("Amount: ", qval.Amount)
		fmt.Println("Time: ", time.Unix(int64(qval.Time.Uint64()), 0))
		fmt.Println("Validator: ", qval.Validator)
	}
	fmt.Println()
}

func GetDelegators(instance *stakepool.Stakepool, addr2 string) {
	delegators, err := instance.GetDelegators(&bind.CallOpts{}, common.HexToAddress(addr2))
	handleError(err)
	fmt.Println()
	fmt.Println("for validator:", addr2)
	for i, d := range delegators {
		fmt.Println("delegator", i, ":", d)
	}
	fmt.Println()
}

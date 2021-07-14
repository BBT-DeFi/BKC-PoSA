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
	Init "win/Code/Init"
	"win/abi/systemreward"
	ETHclient "win/client"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

// systemrewardCmd represents the systemreward command
var systemrewardCmd = &cobra.Command{
	Use:   "systemreward",
	Short: "query the system reward contract",
	Long:  `This contract consists of reward distributing logics`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := ETHclient.Client()
		handleError(err)

		name := "system reward"

		//SystemRewardAddress := "0xd5267A4551754F858EFe111073b0fB96d79409b7"

		SystemRewardInstance, err := systemreward.NewSystemreward(common.HexToAddress(SystemRewardAddress), client)
		handleError(err)

		val, err := cmd.Flags().GetBool("alreadyInit")
		handleError(err)
		addr, err := cmd.Flags().GetString("rewardMapping")
		handleError(err)
		b, err := cmd.Flags().GetBool("getBalance")
		handleError(err)

		if val {
			Init.AlreadyInit(*client, SystemRewardInstance, name)
		} else if addr != "" {
			GetRewardMapping(SystemRewardInstance, addr)
		} else if b {
			GetBalance(SystemRewardInstance)
		} else {
			fmt.Println("please specify the field you want to query")
		}
	},
}

func init() {
	rootCmd.AddCommand(systemrewardCmd)
	systemrewardCmd.Flags().BoolVarP(&alreadyInit, "alreadyInit", "i", false, "the init status")

	systemrewardCmd.Flags().StringVarP(&consensusAddress, "rewardMapping", "r", "", "the reward of the validator (specify address)")

	systemrewardCmd.Flags().BoolVarP(&balance, "getBalance", "b", false, "the balance of the system reward contract")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// systemrewardCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// systemrewardCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func GetRewardMapping(instance *systemreward.Systemreward, addr string) {
	reward, err := instance.RewardMapping(&bind.CallOpts{}, common.HexToAddress(addr))
	handleError(err)
	fmt.Println()
	fmt.Println("The reward for validator address", addr, "is", reward)
	fmt.Println()
}

func GetBalance(instance *systemreward.Systemreward) {
	bal, err := instance.GetBalance(&bind.CallOpts{})
	handleError(err)

	fmt.Println()
	fmt.Println("The balance of the system reward contract is", bal)
	fmt.Println()
}

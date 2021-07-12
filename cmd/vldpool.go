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
	vldpool "win/abi/vldpool"
	ETHclient "win/client"

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
		if err != nil {
			log.Fatal(err)
		}

		name := "validator pool"

		//ValidatorPoolAddress := "0x282659B28f9acCaC9fBd81317Ce43b4044E5d4A7"

		ValidatorPoolInstance, err := vldpool.NewVldpool(common.HexToAddress(ValidatorPoolAddress), client)

		val, err := cmd.Flags().GetBool("alreadyInit")
		if val {
			Init.AlreadyInit(*client, ValidatorPoolInstance, name)
		} else {
			fmt.Println("please specify the field you want to query")
		}
	},
}

func init() {
	rootCmd.AddCommand(vldpoolCmd)
	vldpoolCmd.Flags().BoolVarP(&alreadyInit, "alreadyInit", "i", false, "the init status")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// vldpoolCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// vldpoolCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

package IValidator

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Validator struct {
	ConsensusAddress common.Address
	StakeAmount      *big.Int
	BondStatus       uint8
	IsJail           bool
}

func PrintValidator(v Validator) {
	BondStatus := []string{"BONDED", "UNBONDING", "UNBONDED"}
	fmt.Println("Consensus Address:", v.ConsensusAddress)
	fmt.Println("Stake Amount:", v.StakeAmount)
	fmt.Println("Bond Status:", BondStatus[v.BondStatus])
	fmt.Println("Is Jail:", v.IsJail)
}

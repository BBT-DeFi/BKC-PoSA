package init

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

func AlreadyInit(client ethclient.Client, instance interface {
	AlreadyInit(b *bind.CallOpts) (bool, error)
}, name string) {
	alreadyInit, err := instance.AlreadyInit(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	fmt.Println("The contract", name+"'s", "AlreadyInit", "field is", alreadyInit)
	fmt.Println()
}

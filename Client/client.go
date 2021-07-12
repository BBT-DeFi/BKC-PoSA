package client

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

func Client() (*ethclient.Client, error) {
	client, err := ethclient.Dial("http://localhost:7545")

	return client, err
}

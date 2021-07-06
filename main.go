package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"

	validatorpool "ethclient/posa/ValidatorPool"
)

func main() {
	client, err := ethclient.Dial("https://rpc-testnet.bitkubchain.io")
	if err != nil {
		log.Fatal(err)
	}

	validatorsetAddress := common.HexToAddress("0x61a75be699e3B21488aFae66f1133A06301F5B96")
	stakePoolAddress := common.HexToAddress("0x46629fa29d04227b422739aD26f4c559fe1EfE14")
	// systemrewardAddress := common.HexToAddress("0xe6BDc0d5ff96A4842AC209c0c1Ebe3FE53db818b")
	validatorPoolAddress := common.HexToAddress("0x756F84C5F00DcA7141e1CdD5F00Dd508BBD53410")

	ValidatorPoolInstance, err := validatorpool.NewValidatorpool(validatorPoolAddress, client)

	alreadyInitValPoolBefore, err := ValidatorPoolInstance.AlreadyInit(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Before Init: ", alreadyInitValPoolBefore)

	txHash, err := initValidatorPool(*client, validatorPoolAddress, validatorsetAddress, stakePoolAddress)
	log.Println(txHash)

	time.Sleep(5 * time.Second)

	alreadyInitValPoolAfter, err := ValidatorPoolInstance.AlreadyInit(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("After Init: ", alreadyInitValPoolAfter)
}

//----------------------------An example of init() function----------------------------
func initValidatorPool(client ethclient.Client, valPoolAddr common.Address, valSetAddr common.Address, stakePoolAddr common.Address) (string, error) {
	privateKey, err := crypto.HexToECDSA("PRIVATE KEY")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	transferFnSignature := []byte("init(address,address)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)

	methodID := hash.Sum(nil)[:4]
	arg1 := common.LeftPadBytes(valSetAddr.Bytes(), 32)
	arg2 := common.LeftPadBytes(stakePoolAddr.Bytes(), 32)

	var data []byte
	data = append(data, methodID...)
	data = append(data, arg1...)
	data = append(data, arg2...)

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &valPoolAddr,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}

	tx := types.NewTransaction(nonce, valPoolAddr, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Transaction hash=%s", signedTx.Hash().Hex()) // tx sent: 0xa56316b637a94c4cc0331c73ef26389d6c097506d581073f927275e7a6ece0bc
	txHash := signedTx.Hash().Hex()
	return txHash, nil
}

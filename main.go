package main

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/binary"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"

	test "tasks/abi"
)

func main() {

	client, err := ethclient.Dial("https://rpc-testnet.bitkubchain.io")
	//client, err := ethclient.Dial("https://data-seed-prebsc-1-s1.binance.org:8545")
	//client, err := ethclient.Dial("http://localhost:7545")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// privatekey, publickey, address, err := createNewWallet()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// _ = privatekey
	// _ = publickey
	// _ = address

	//getBlockNumber(*client, big.NewInt(1588791))

	privateKey := "Private Key"
	_ = privateKey

	// sendETH(*client, privateKey,
	// 	"0xf5751Aef65a4ED92C109f045964a5B0C6cf7E1Fe", *big.NewInt(1000000000000000000))

	testAddress := common.HexToAddress("0x92035dF7e4589398EDAcd4AaB4c61E73B00F6841")

	testContract, err := test.NewTest(testAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	_ = testContract

	// value := getValue(*client, privateKey, "test", test)
	// fmt.Println(value)

	//setNum(*client, privateKey, 1000, testAddress)
	//setValue(*client, privateKey, "test", big.NewInt(15), testAddress)
	//setName(*client, privateKey, "test", testAddress)
	//setAddress(*client, privateKey, "0x60bB6c1B4a0F5B1ea820be6c762384982b8a658c", testAddress)
	// t := common.HexToAddress("0x64A915DbBfce21a526B7febfae25ca641840d8ce")
	// fmt.Println(t)
	//setAddressV2(*client, privateKey, testAddress.String(), testAddress)
	//testf(*client, privateKey, "0x60bB6c1B4a0F5B1ea820be6c762384982b8a658c", testAddress)
	// addr, err := testContract.Addr(&bind.CallOpts{})
	// fmt.Println(addr)
	// name, err := testContract.Name(&bind.CallOpts{})
	// fmt.Println(name)
	// value, err := testContract.Pair(&bind.CallOpts{}, "test")
	// fmt.Println(value)
	num, err := testContract.Num(&bind.CallOpts{})
	fmt.Println(num)
}

func createNewWallet() (string, string, string, error) {

	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", "", err
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	privatekey := hexutil.Encode(privateKeyBytes)[2:]
	fmt.Println("The private key is", hexutil.Encode(privateKeyBytes)[2:])

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey) // type assertion
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	publickey := hexutil.Encode(publicKeyBytes)[4:]
	fmt.Println("The public key is", hexutil.Encode(publicKeyBytes)[4:])
	//fmt.Println("The public key is", publicKey)
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("The address is", address)
	return privatekey, publickey, address, nil
}

func getBlockNumber(client ethclient.Client, number *big.Int) {
	block, err := client.BlockByNumber(context.Background(), number)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The block number is", block.Number().Uint64())
	fmt.Println("The time is", block.Time())
	fmt.Println("The block difficulty is", block.Difficulty().Uint64())
	fmt.Println("The block hash is", block.Hash().Hex())
	fmt.Println("The number of transaction is ", len(block.Transactions()))

	for i, tx := range block.Transactions() {
		fmt.Println("For transaction number", i+1)
		fmt.Println(tx.Hash().Hex())
		fmt.Println(tx.Value().String())
		fmt.Println(tx.Gas())
		fmt.Println(tx.GasPrice().Uint64())
		fmt.Println(tx.Nonce())
		fmt.Println(tx.Data())
		fmt.Println(tx.To().Hex())
	}
}

func generatePrivateKey() (string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", err
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	return hexutil.Encode(privateKeyBytes)[2:], nil
}

func sendETH(client ethclient.Client, privatekey string, to string, value big.Int) {
	privateKey, err := crypto.HexToECDSA(privatekey)
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

	gasLimit := uint64(21000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	toAddress := common.HexToAddress(to)
	tx := types.NewTransaction(nonce, toAddress, &value, gasLimit, gasPrice, nil)
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

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}

func sendTransaction(client ethclient.Client, privatekey string, data []byte, to common.Address) (string, error) {
	privateKey, err := crypto.HexToECDSA(privatekey)
	if err != nil {
		fmt.Println("1")
		fmt.Println(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Println("2")
		log.Fatal(err)
	}

	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println("3")
		log.Fatal(err)
	}
	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &to,
		Data: data,
	})
	if err != nil {
		fmt.Println("4")
		return "", err
	}

	tx := types.NewTransaction(nonce, to, value, gasLimit, gasPrice, data)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		fmt.Println("5")
		return "", err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		fmt.Println("6")
		return "", err
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}

	fmt.Printf("Transaction hash=%s\n", signedTx.Hash().Hex())
	txHash := signedTx.Hash().Hex()
	return txHash, nil
}

func getAuth(client ethclient.Client, privatekey string) *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA(privatekey)
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

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	return auth
}

func getValue(client ethclient.Client, privatekey string, key string, test *test.Test) *big.Int {
	value, err := test.Pair(&bind.CallOpts{}, "test")
	if err != nil {
		log.Fatal(err)
	}
	return value
}

func setValue(client ethclient.Client, privatekey string, key string, value *big.Int, to common.Address) {
	transferFnSignature := []byte("setValue(string, uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	var data []byte
	data = append(data, methodID...)
	data = append(data, key...)
	data = append(data, value.String()...)
	Txs, err := sendTransaction(client, privatekey, data, to)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(Txs)
}

func setName(client ethclient.Client, privatekey string, name string, to common.Address) {
	transferFnSignature := []byte("setName(string)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	var data []byte
	data = append(data, methodID...)
	bs := []byte(name)
	arg1 := common.LeftPadBytes(bs, 32)
	fmt.Println(arg1)
	data = append(data, arg1...)
	Txs, err := sendTransaction(client, privatekey, data, to)
	if err != nil {
		log.Fatal(err)
	}
	_ = Txs

}

func setNum(client ethclient.Client, privatekey string, num uint64, to common.Address) {
	transferFnSignature := []byte("setNum(uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	var data []byte
	data = append(data, methodID...)
	//arg1 := common.LeftPadBytes([]byte(strconv.Itoa(num)), 32)
	bs := new(bytes.Buffer)
	err := binary.Write(bs, binary.BigEndian, num)
	if err != nil {
		log.Fatal(err)
	}
	arg1 := bs.Bytes()
	arg1 = common.LeftPadBytes(arg1, 32)
	//fmt.Println(common.LeftPadBytes(arg1, 32))
	//fmt.Println(common.LeftPadBytes([]byte{1}, 32))
	data = append(data, arg1...)
	Txs, err := sendTransaction(client, privatekey, data, to)
	if err != nil {
		log.Fatal(err)
	}
	_ = Txs
}

// func setNumV2(client ethclient.Client, privatekey string, to common.Address) {
// 	abiTest := "[{\"inputs\":[],\"name\":\"addr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"num\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"pair\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"setAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_num\",\"type\":\"uint256\"}],\"name\":\"setNum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// 	address := to
// 	d := time.Now().Add(1000 * time.Millisecond)
// 	ctx, cancel := context.WithDeadline(context.Background(), d)
// 	defer cancel()
// 	testabi, err := abi.JSON(strings.NewReader(abiTest))
// 	privateKey, err := crypto.HexToECDSA(privatekey)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	publicKey := privateKey.Public()
// 	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
// 	if !ok {
// 		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
// 	}
// 	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
// 	nonce, _ := client.NonceAt(ctx, fromAddress, nil)
// 	bytesData, _= testabi.Pack("setNum", 5)
// }

func setAddress(client ethclient.Client, privatekey string, a string, to common.Address) {
	transferFnSignature := []byte("setAddress(address)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	var data []byte
	data = append(data, methodID...)
	arg1 := common.LeftPadBytes(common.HexToAddress(a).Bytes(), 32)
	data = append(data, arg1...)
	Txs, err := sendTransaction(client, privatekey, data, to)
	if err != nil {
		log.Fatal(err)
	}

	_ = Txs

}

func setAddressV2(client ethclient.Client, privatekey string, a string, to common.Address) {
	auth := getAuth(client, privatekey)
	test, err := test.NewTest(to, &client)
	if err != nil {
		log.Fatal(err)
	}
	txs, err := test.SetAddress(auth, common.HexToAddress(a))
	if err != nil {
		log.Fatal(err)
	}
	_ = txs
}

func testf(client ethclient.Client, privatekey string, a string, to common.Address) (string, error) {
	privateKey, err := crypto.HexToECDSA(privatekey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	// fmt.Println(fromAddress)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	transferFnSignature := []byte("setAddress(address)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)

	methodID := hash.Sum(nil)[:4]
	arg1 := common.LeftPadBytes(common.HexToAddress(a).Bytes(), 32)

	var data []byte
	data = append(data, methodID...)
	data = append(data, arg1...)

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &to,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}

	tx := types.NewTransaction(nonce, to, value, gasLimit, gasPrice, data)

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

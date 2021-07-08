// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package test

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TestABI is the input ABI used to generate the binding from.
const TestABI = "[{\"inputs\":[],\"name\":\"addr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"num\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"pair\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"setAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_num\",\"type\":\"uint256\"}],\"name\":\"setNum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TestBin is the compiled bytecode used for deploying new contracts.
var TestBin = "0x608060405234801561001057600080fd5b5061091c806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80637c2da7aa1161005b5780637c2da7aa14610117578063c47f002714610133578063cd16ecbf1461014f578063e30081a01461016b57610088565b806306fdde031461008d5780631a384e10146100ab5780634e70b1dc146100db578063767800de146100f9575b600080fd5b610095610187565b6040516100a29190610690565b60405180910390f35b6100c560048036038101906100c0919061056a565b610215565b6040516100d291906106b2565b60405180910390f35b6100e3610243565b6040516100f091906106b2565b60405180910390f35b610101610249565b60405161010e9190610675565b60405180910390f35b610131600480360381019061012c919061050a565b61026f565b005b61014d600480360381019061014891906104bd565b610299565b005b610169600480360381019061016491906105b3565b6102af565b005b61018560048036038101906101809190610490565b6102b9565b005b60028054610194906107c8565b80601f01602080910402602001604051908101604052809291908181526020018280546101c0906107c8565b801561020d5780601f106101e25761010080835404028352916020019161020d565b820191906000526020600020905b8154815290600101906020018083116101f057829003601f168201915b505050505081565b6000818051602081018201805184825260208301602085012081835280955050505050506000915090505481565b60015481565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b806000848460405161028292919061065c565b908152602001604051809103902081905550505050565b8181600291906102aa9291906102fd565b505050565b8060018190555050565b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b828054610309906107c8565b90600052602060002090601f01602090048101928261032b5760008555610372565b82601f1061034457803560ff1916838001178555610372565b82800160010185558215610372579182015b82811115610371578235825591602001919060010190610356565b5b50905061037f9190610383565b5090565b5b8082111561039c576000816000905550600101610384565b5090565b60006103b36103ae846106f2565b6106cd565b9050828152602081018484840111156103cf576103ce610898565b5b6103da848285610786565b509392505050565b6000813590506103f1816108b8565b92915050565b60008083601f84011261040d5761040c61088e565b5b8235905067ffffffffffffffff81111561042a57610429610889565b5b60208301915083600182028301111561044657610445610893565b5b9250929050565b600082601f8301126104625761046161088e565b5b81356104728482602086016103a0565b91505092915050565b60008135905061048a816108cf565b92915050565b6000602082840312156104a6576104a56108a2565b5b60006104b4848285016103e2565b91505092915050565b600080602083850312156104d4576104d36108a2565b5b600083013567ffffffffffffffff8111156104f2576104f161089d565b5b6104fe858286016103f7565b92509250509250929050565b600080600060408486031215610523576105226108a2565b5b600084013567ffffffffffffffff8111156105415761054061089d565b5b61054d868287016103f7565b935093505060206105608682870161047b565b9150509250925092565b6000602082840312156105805761057f6108a2565b5b600082013567ffffffffffffffff81111561059e5761059d61089d565b5b6105aa8482850161044d565b91505092915050565b6000602082840312156105c9576105c86108a2565b5b60006105d78482850161047b565b91505092915050565b6105e98161074a565b82525050565b60006105fb838561073f565b9350610608838584610786565b82840190509392505050565b600061061f82610723565b610629818561072e565b9350610639818560208601610795565b610642816108a7565b840191505092915050565b6106568161077c565b82525050565b60006106698284866105ef565b91508190509392505050565b600060208201905061068a60008301846105e0565b92915050565b600060208201905081810360008301526106aa8184610614565b905092915050565b60006020820190506106c7600083018461064d565b92915050565b60006106d76106e8565b90506106e382826107fa565b919050565b6000604051905090565b600067ffffffffffffffff82111561070d5761070c61085a565b5b610716826108a7565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b60006107558261075c565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b838110156107b3578082015181840152602081019050610798565b838111156107c2576000848401525b50505050565b600060028204905060018216806107e057607f821691505b602082108114156107f4576107f361082b565b5b50919050565b610803826108a7565b810181811067ffffffffffffffff821117156108225761082161085a565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b6108c18161074a565b81146108cc57600080fd5b50565b6108d88161077c565b81146108e357600080fd5b5056fea26469706673582212206a49c99c6ced7f54eec0dbebe61e7ec20cd6a1711985f9630e478cffa722649d64736f6c63430008050033"

// DeployTest deploys a new Ethereum contract, binding an instance of Test to it.
func DeployTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Test, error) {
	parsed, err := abi.JSON(strings.NewReader(TestABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Test{TestCaller: TestCaller{contract: contract}, TestTransactor: TestTransactor{contract: contract}, TestFilterer: TestFilterer{contract: contract}}, nil
}

// Test is an auto generated Go binding around an Ethereum contract.
type Test struct {
	TestCaller     // Read-only binding to the contract
	TestTransactor // Write-only binding to the contract
	TestFilterer   // Log filterer for contract events
}

// TestCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestSession struct {
	Contract     *Test             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestCallerSession struct {
	Contract *TestCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestTransactorSession struct {
	Contract     *TestTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestRaw struct {
	Contract *Test // Generic contract binding to access the raw methods on
}

// TestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestCallerRaw struct {
	Contract *TestCaller // Generic read-only contract binding to access the raw methods on
}

// TestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestTransactorRaw struct {
	Contract *TestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTest creates a new instance of Test, bound to a specific deployed contract.
func NewTest(address common.Address, backend bind.ContractBackend) (*Test, error) {
	contract, err := bindTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Test{TestCaller: TestCaller{contract: contract}, TestTransactor: TestTransactor{contract: contract}, TestFilterer: TestFilterer{contract: contract}}, nil
}

// NewTestCaller creates a new read-only instance of Test, bound to a specific deployed contract.
func NewTestCaller(address common.Address, caller bind.ContractCaller) (*TestCaller, error) {
	contract, err := bindTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestCaller{contract: contract}, nil
}

// NewTestTransactor creates a new write-only instance of Test, bound to a specific deployed contract.
func NewTestTransactor(address common.Address, transactor bind.ContractTransactor) (*TestTransactor, error) {
	contract, err := bindTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestTransactor{contract: contract}, nil
}

// NewTestFilterer creates a new log filterer instance of Test, bound to a specific deployed contract.
func NewTestFilterer(address common.Address, filterer bind.ContractFilterer) (*TestFilterer, error) {
	contract, err := bindTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestFilterer{contract: contract}, nil
}

// bindTest binds a generic wrapper to an already deployed contract.
func bindTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Test *TestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Test.Contract.TestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Test *TestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test.Contract.TestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Test *TestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Test.Contract.TestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Test *TestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Test.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Test *TestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Test *TestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Test.Contract.contract.Transact(opts, method, params...)
}

// Addr is a free data retrieval call binding the contract method 0x767800de.
//
// Solidity: function addr() view returns(address)
func (_Test *TestCaller) Addr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Test.contract.Call(opts, &out, "addr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addr is a free data retrieval call binding the contract method 0x767800de.
//
// Solidity: function addr() view returns(address)
func (_Test *TestSession) Addr() (common.Address, error) {
	return _Test.Contract.Addr(&_Test.CallOpts)
}

// Addr is a free data retrieval call binding the contract method 0x767800de.
//
// Solidity: function addr() view returns(address)
func (_Test *TestCallerSession) Addr() (common.Address, error) {
	return _Test.Contract.Addr(&_Test.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Test *TestCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Test.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Test *TestSession) Name() (string, error) {
	return _Test.Contract.Name(&_Test.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Test *TestCallerSession) Name() (string, error) {
	return _Test.Contract.Name(&_Test.CallOpts)
}

// Num is a free data retrieval call binding the contract method 0x4e70b1dc.
//
// Solidity: function num() view returns(uint256)
func (_Test *TestCaller) Num(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Test.contract.Call(opts, &out, "num")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Num is a free data retrieval call binding the contract method 0x4e70b1dc.
//
// Solidity: function num() view returns(uint256)
func (_Test *TestSession) Num() (*big.Int, error) {
	return _Test.Contract.Num(&_Test.CallOpts)
}

// Num is a free data retrieval call binding the contract method 0x4e70b1dc.
//
// Solidity: function num() view returns(uint256)
func (_Test *TestCallerSession) Num() (*big.Int, error) {
	return _Test.Contract.Num(&_Test.CallOpts)
}

// Pair is a free data retrieval call binding the contract method 0x1a384e10.
//
// Solidity: function pair(string ) view returns(uint256)
func (_Test *TestCaller) Pair(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _Test.contract.Call(opts, &out, "pair", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Pair is a free data retrieval call binding the contract method 0x1a384e10.
//
// Solidity: function pair(string ) view returns(uint256)
func (_Test *TestSession) Pair(arg0 string) (*big.Int, error) {
	return _Test.Contract.Pair(&_Test.CallOpts, arg0)
}

// Pair is a free data retrieval call binding the contract method 0x1a384e10.
//
// Solidity: function pair(string ) view returns(uint256)
func (_Test *TestCallerSession) Pair(arg0 string) (*big.Int, error) {
	return _Test.Contract.Pair(&_Test.CallOpts, arg0)
}

// SetAddress is a paid mutator transaction binding the contract method 0xe30081a0.
//
// Solidity: function setAddress(address a) returns()
func (_Test *TestTransactor) SetAddress(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _Test.contract.Transact(opts, "setAddress", a)
}

// SetAddress is a paid mutator transaction binding the contract method 0xe30081a0.
//
// Solidity: function setAddress(address a) returns()
func (_Test *TestSession) SetAddress(a common.Address) (*types.Transaction, error) {
	return _Test.Contract.SetAddress(&_Test.TransactOpts, a)
}

// SetAddress is a paid mutator transaction binding the contract method 0xe30081a0.
//
// Solidity: function setAddress(address a) returns()
func (_Test *TestTransactorSession) SetAddress(a common.Address) (*types.Transaction, error) {
	return _Test.Contract.SetAddress(&_Test.TransactOpts, a)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string _name) returns()
func (_Test *TestTransactor) SetName(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _Test.contract.Transact(opts, "setName", _name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string _name) returns()
func (_Test *TestSession) SetName(_name string) (*types.Transaction, error) {
	return _Test.Contract.SetName(&_Test.TransactOpts, _name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string _name) returns()
func (_Test *TestTransactorSession) SetName(_name string) (*types.Transaction, error) {
	return _Test.Contract.SetName(&_Test.TransactOpts, _name)
}

// SetNum is a paid mutator transaction binding the contract method 0xcd16ecbf.
//
// Solidity: function setNum(uint256 _num) returns()
func (_Test *TestTransactor) SetNum(opts *bind.TransactOpts, _num *big.Int) (*types.Transaction, error) {
	return _Test.contract.Transact(opts, "setNum", _num)
}

// SetNum is a paid mutator transaction binding the contract method 0xcd16ecbf.
//
// Solidity: function setNum(uint256 _num) returns()
func (_Test *TestSession) SetNum(_num *big.Int) (*types.Transaction, error) {
	return _Test.Contract.SetNum(&_Test.TransactOpts, _num)
}

// SetNum is a paid mutator transaction binding the contract method 0xcd16ecbf.
//
// Solidity: function setNum(uint256 _num) returns()
func (_Test *TestTransactorSession) SetNum(_num *big.Int) (*types.Transaction, error) {
	return _Test.Contract.SetNum(&_Test.TransactOpts, _num)
}

// SetValue is a paid mutator transaction binding the contract method 0x7c2da7aa.
//
// Solidity: function setValue(string key, uint256 value) returns()
func (_Test *TestTransactor) SetValue(opts *bind.TransactOpts, key string, value *big.Int) (*types.Transaction, error) {
	return _Test.contract.Transact(opts, "setValue", key, value)
}

// SetValue is a paid mutator transaction binding the contract method 0x7c2da7aa.
//
// Solidity: function setValue(string key, uint256 value) returns()
func (_Test *TestSession) SetValue(key string, value *big.Int) (*types.Transaction, error) {
	return _Test.Contract.SetValue(&_Test.TransactOpts, key, value)
}

// SetValue is a paid mutator transaction binding the contract method 0x7c2da7aa.
//
// Solidity: function setValue(string key, uint256 value) returns()
func (_Test *TestTransactorSession) SetValue(key string, value *big.Int) (*types.Transaction, error) {
	return _Test.Contract.SetValue(&_Test.TransactOpts, key, value)
}

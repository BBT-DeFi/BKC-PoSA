// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package systemreward

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

// SystemrewardABI is the input ABI used to generate the binding from.
const SystemrewardABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"consensusAddress\",\"type\":\"address\"}],\"name\":\"AllocateJailValidatorReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"consensusAddress\",\"type\":\"address\"}],\"name\":\"addReward\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"alreadyInit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributeReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fund\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"consensusAddress\",\"type\":\"address\"}],\"name\":\"getTotalPower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"consensusAddress\",\"type\":\"address\"}],\"name\":\"getTotalPowerExcludeUnbonding\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"BKCValidatorSetAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"StakePoolAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ValidatorPoolAddress\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardMapping\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Systemreward is an auto generated Go binding around an Ethereum contract.
type Systemreward struct {
	SystemrewardCaller     // Read-only binding to the contract
	SystemrewardTransactor // Write-only binding to the contract
	SystemrewardFilterer   // Log filterer for contract events
}

// SystemrewardCaller is an auto generated read-only Go binding around an Ethereum contract.
type SystemrewardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemrewardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SystemrewardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemrewardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SystemrewardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemrewardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SystemrewardSession struct {
	Contract     *Systemreward     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SystemrewardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SystemrewardCallerSession struct {
	Contract *SystemrewardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// SystemrewardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SystemrewardTransactorSession struct {
	Contract     *SystemrewardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SystemrewardRaw is an auto generated low-level Go binding around an Ethereum contract.
type SystemrewardRaw struct {
	Contract *Systemreward // Generic contract binding to access the raw methods on
}

// SystemrewardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SystemrewardCallerRaw struct {
	Contract *SystemrewardCaller // Generic read-only contract binding to access the raw methods on
}

// SystemrewardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SystemrewardTransactorRaw struct {
	Contract *SystemrewardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSystemreward creates a new instance of Systemreward, bound to a specific deployed contract.
func NewSystemreward(address common.Address, backend bind.ContractBackend) (*Systemreward, error) {
	contract, err := bindSystemreward(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Systemreward{SystemrewardCaller: SystemrewardCaller{contract: contract}, SystemrewardTransactor: SystemrewardTransactor{contract: contract}, SystemrewardFilterer: SystemrewardFilterer{contract: contract}}, nil
}

// NewSystemrewardCaller creates a new read-only instance of Systemreward, bound to a specific deployed contract.
func NewSystemrewardCaller(address common.Address, caller bind.ContractCaller) (*SystemrewardCaller, error) {
	contract, err := bindSystemreward(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SystemrewardCaller{contract: contract}, nil
}

// NewSystemrewardTransactor creates a new write-only instance of Systemreward, bound to a specific deployed contract.
func NewSystemrewardTransactor(address common.Address, transactor bind.ContractTransactor) (*SystemrewardTransactor, error) {
	contract, err := bindSystemreward(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SystemrewardTransactor{contract: contract}, nil
}

// NewSystemrewardFilterer creates a new log filterer instance of Systemreward, bound to a specific deployed contract.
func NewSystemrewardFilterer(address common.Address, filterer bind.ContractFilterer) (*SystemrewardFilterer, error) {
	contract, err := bindSystemreward(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SystemrewardFilterer{contract: contract}, nil
}

// bindSystemreward binds a generic wrapper to an already deployed contract.
func bindSystemreward(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SystemrewardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Systemreward *SystemrewardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Systemreward.Contract.SystemrewardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Systemreward *SystemrewardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Systemreward.Contract.SystemrewardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Systemreward *SystemrewardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Systemreward.Contract.SystemrewardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Systemreward *SystemrewardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Systemreward.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Systemreward *SystemrewardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Systemreward.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Systemreward *SystemrewardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Systemreward.Contract.contract.Transact(opts, method, params...)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_Systemreward *SystemrewardCaller) AlreadyInit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Systemreward.contract.Call(opts, &out, "alreadyInit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_Systemreward *SystemrewardSession) AlreadyInit() (bool, error) {
	return _Systemreward.Contract.AlreadyInit(&_Systemreward.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_Systemreward *SystemrewardCallerSession) AlreadyInit() (bool, error) {
	return _Systemreward.Contract.AlreadyInit(&_Systemreward.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Systemreward *SystemrewardCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Systemreward.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Systemreward *SystemrewardSession) GetBalance() (*big.Int, error) {
	return _Systemreward.Contract.GetBalance(&_Systemreward.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Systemreward *SystemrewardCallerSession) GetBalance() (*big.Int, error) {
	return _Systemreward.Contract.GetBalance(&_Systemreward.CallOpts)
}

// GetTotalPower is a free data retrieval call binding the contract method 0xd7dc81cd.
//
// Solidity: function getTotalPower(address consensusAddress) view returns(uint256)
func (_Systemreward *SystemrewardCaller) GetTotalPower(opts *bind.CallOpts, consensusAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Systemreward.contract.Call(opts, &out, "getTotalPower", consensusAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPower is a free data retrieval call binding the contract method 0xd7dc81cd.
//
// Solidity: function getTotalPower(address consensusAddress) view returns(uint256)
func (_Systemreward *SystemrewardSession) GetTotalPower(consensusAddress common.Address) (*big.Int, error) {
	return _Systemreward.Contract.GetTotalPower(&_Systemreward.CallOpts, consensusAddress)
}

// GetTotalPower is a free data retrieval call binding the contract method 0xd7dc81cd.
//
// Solidity: function getTotalPower(address consensusAddress) view returns(uint256)
func (_Systemreward *SystemrewardCallerSession) GetTotalPower(consensusAddress common.Address) (*big.Int, error) {
	return _Systemreward.Contract.GetTotalPower(&_Systemreward.CallOpts, consensusAddress)
}

// GetTotalPowerExcludeUnbonding is a free data retrieval call binding the contract method 0x322923b9.
//
// Solidity: function getTotalPowerExcludeUnbonding(address consensusAddress) view returns(uint256)
func (_Systemreward *SystemrewardCaller) GetTotalPowerExcludeUnbonding(opts *bind.CallOpts, consensusAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Systemreward.contract.Call(opts, &out, "getTotalPowerExcludeUnbonding", consensusAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPowerExcludeUnbonding is a free data retrieval call binding the contract method 0x322923b9.
//
// Solidity: function getTotalPowerExcludeUnbonding(address consensusAddress) view returns(uint256)
func (_Systemreward *SystemrewardSession) GetTotalPowerExcludeUnbonding(consensusAddress common.Address) (*big.Int, error) {
	return _Systemreward.Contract.GetTotalPowerExcludeUnbonding(&_Systemreward.CallOpts, consensusAddress)
}

// GetTotalPowerExcludeUnbonding is a free data retrieval call binding the contract method 0x322923b9.
//
// Solidity: function getTotalPowerExcludeUnbonding(address consensusAddress) view returns(uint256)
func (_Systemreward *SystemrewardCallerSession) GetTotalPowerExcludeUnbonding(consensusAddress common.Address) (*big.Int, error) {
	return _Systemreward.Contract.GetTotalPowerExcludeUnbonding(&_Systemreward.CallOpts, consensusAddress)
}

// RewardMapping is a free data retrieval call binding the contract method 0x571a9154.
//
// Solidity: function rewardMapping(address ) view returns(uint256)
func (_Systemreward *SystemrewardCaller) RewardMapping(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Systemreward.contract.Call(opts, &out, "rewardMapping", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardMapping is a free data retrieval call binding the contract method 0x571a9154.
//
// Solidity: function rewardMapping(address ) view returns(uint256)
func (_Systemreward *SystemrewardSession) RewardMapping(arg0 common.Address) (*big.Int, error) {
	return _Systemreward.Contract.RewardMapping(&_Systemreward.CallOpts, arg0)
}

// RewardMapping is a free data retrieval call binding the contract method 0x571a9154.
//
// Solidity: function rewardMapping(address ) view returns(uint256)
func (_Systemreward *SystemrewardCallerSession) RewardMapping(arg0 common.Address) (*big.Int, error) {
	return _Systemreward.Contract.RewardMapping(&_Systemreward.CallOpts, arg0)
}

// AllocateJailValidatorReward is a paid mutator transaction binding the contract method 0x348749e0.
//
// Solidity: function AllocateJailValidatorReward(address consensusAddress) returns()
func (_Systemreward *SystemrewardTransactor) AllocateJailValidatorReward(opts *bind.TransactOpts, consensusAddress common.Address) (*types.Transaction, error) {
	return _Systemreward.contract.Transact(opts, "AllocateJailValidatorReward", consensusAddress)
}

// AllocateJailValidatorReward is a paid mutator transaction binding the contract method 0x348749e0.
//
// Solidity: function AllocateJailValidatorReward(address consensusAddress) returns()
func (_Systemreward *SystemrewardSession) AllocateJailValidatorReward(consensusAddress common.Address) (*types.Transaction, error) {
	return _Systemreward.Contract.AllocateJailValidatorReward(&_Systemreward.TransactOpts, consensusAddress)
}

// AllocateJailValidatorReward is a paid mutator transaction binding the contract method 0x348749e0.
//
// Solidity: function AllocateJailValidatorReward(address consensusAddress) returns()
func (_Systemreward *SystemrewardTransactorSession) AllocateJailValidatorReward(consensusAddress common.Address) (*types.Transaction, error) {
	return _Systemreward.Contract.AllocateJailValidatorReward(&_Systemreward.TransactOpts, consensusAddress)
}

// AddReward is a paid mutator transaction binding the contract method 0x9c9b2e21.
//
// Solidity: function addReward(address consensusAddress) payable returns()
func (_Systemreward *SystemrewardTransactor) AddReward(opts *bind.TransactOpts, consensusAddress common.Address) (*types.Transaction, error) {
	return _Systemreward.contract.Transact(opts, "addReward", consensusAddress)
}

// AddReward is a paid mutator transaction binding the contract method 0x9c9b2e21.
//
// Solidity: function addReward(address consensusAddress) payable returns()
func (_Systemreward *SystemrewardSession) AddReward(consensusAddress common.Address) (*types.Transaction, error) {
	return _Systemreward.Contract.AddReward(&_Systemreward.TransactOpts, consensusAddress)
}

// AddReward is a paid mutator transaction binding the contract method 0x9c9b2e21.
//
// Solidity: function addReward(address consensusAddress) payable returns()
func (_Systemreward *SystemrewardTransactorSession) AddReward(consensusAddress common.Address) (*types.Transaction, error) {
	return _Systemreward.Contract.AddReward(&_Systemreward.TransactOpts, consensusAddress)
}

// DistributeReward is a paid mutator transaction binding the contract method 0x8f73c5ae.
//
// Solidity: function distributeReward() returns()
func (_Systemreward *SystemrewardTransactor) DistributeReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Systemreward.contract.Transact(opts, "distributeReward")
}

// DistributeReward is a paid mutator transaction binding the contract method 0x8f73c5ae.
//
// Solidity: function distributeReward() returns()
func (_Systemreward *SystemrewardSession) DistributeReward() (*types.Transaction, error) {
	return _Systemreward.Contract.DistributeReward(&_Systemreward.TransactOpts)
}

// DistributeReward is a paid mutator transaction binding the contract method 0x8f73c5ae.
//
// Solidity: function distributeReward() returns()
func (_Systemreward *SystemrewardTransactorSession) DistributeReward() (*types.Transaction, error) {
	return _Systemreward.Contract.DistributeReward(&_Systemreward.TransactOpts)
}

// Fund is a paid mutator transaction binding the contract method 0xb60d4288.
//
// Solidity: function fund() payable returns()
func (_Systemreward *SystemrewardTransactor) Fund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Systemreward.contract.Transact(opts, "fund")
}

// Fund is a paid mutator transaction binding the contract method 0xb60d4288.
//
// Solidity: function fund() payable returns()
func (_Systemreward *SystemrewardSession) Fund() (*types.Transaction, error) {
	return _Systemreward.Contract.Fund(&_Systemreward.TransactOpts)
}

// Fund is a paid mutator transaction binding the contract method 0xb60d4288.
//
// Solidity: function fund() payable returns()
func (_Systemreward *SystemrewardTransactorSession) Fund() (*types.Transaction, error) {
	return _Systemreward.Contract.Fund(&_Systemreward.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0x184b9559.
//
// Solidity: function init(address BKCValidatorSetAddress, address StakePoolAddress, address ValidatorPoolAddress) returns()
func (_Systemreward *SystemrewardTransactor) Init(opts *bind.TransactOpts, BKCValidatorSetAddress common.Address, StakePoolAddress common.Address, ValidatorPoolAddress common.Address) (*types.Transaction, error) {
	return _Systemreward.contract.Transact(opts, "init", BKCValidatorSetAddress, StakePoolAddress, ValidatorPoolAddress)
}

// Init is a paid mutator transaction binding the contract method 0x184b9559.
//
// Solidity: function init(address BKCValidatorSetAddress, address StakePoolAddress, address ValidatorPoolAddress) returns()
func (_Systemreward *SystemrewardSession) Init(BKCValidatorSetAddress common.Address, StakePoolAddress common.Address, ValidatorPoolAddress common.Address) (*types.Transaction, error) {
	return _Systemreward.Contract.Init(&_Systemreward.TransactOpts, BKCValidatorSetAddress, StakePoolAddress, ValidatorPoolAddress)
}

// Init is a paid mutator transaction binding the contract method 0x184b9559.
//
// Solidity: function init(address BKCValidatorSetAddress, address StakePoolAddress, address ValidatorPoolAddress) returns()
func (_Systemreward *SystemrewardTransactorSession) Init(BKCValidatorSetAddress common.Address, StakePoolAddress common.Address, ValidatorPoolAddress common.Address) (*types.Transaction, error) {
	return _Systemreward.Contract.Init(&_Systemreward.TransactOpts, BKCValidatorSetAddress, StakePoolAddress, ValidatorPoolAddress)
}

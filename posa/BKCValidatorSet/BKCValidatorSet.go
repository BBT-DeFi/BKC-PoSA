// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package validatorset

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

// IValidatorValidator is an auto generated low-level Go binding around an user-defined struct.
type IValidatorValidator struct {
	ConsensusAddress common.Address
	StakeAmount      *big.Int
	BondStatus       uint8
	IsJail           bool
}

// ValidatorsetABI is the input ABI used to generate the binding from.
const ValidatorsetABI = "[{\"inputs\":[],\"name\":\"alreadyInit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"currentValidatorSet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"consensusAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"enumIBond.BondStatus\",\"name\":\"bondStatus\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isJail\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"currentValidatorSetMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"consensusAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"enumIBond.BondStatus\",\"name\":\"bondStatus\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isJail\",\"type\":\"bool\"}],\"internalType\":\"structIValidator.Validator[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorPoolAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"systemRewardAddress\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"consensusAddress\",\"type\":\"address\"}],\"name\":\"jailValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"number_of_validators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Validatorset is an auto generated Go binding around an Ethereum contract.
type Validatorset struct {
	ValidatorsetCaller     // Read-only binding to the contract
	ValidatorsetTransactor // Write-only binding to the contract
	ValidatorsetFilterer   // Log filterer for contract events
}

// ValidatorsetCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorsetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorsetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorsetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorsetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorsetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorsetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorsetSession struct {
	Contract     *Validatorset     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorsetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorsetCallerSession struct {
	Contract *ValidatorsetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ValidatorsetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorsetTransactorSession struct {
	Contract     *ValidatorsetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ValidatorsetRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorsetRaw struct {
	Contract *Validatorset // Generic contract binding to access the raw methods on
}

// ValidatorsetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorsetCallerRaw struct {
	Contract *ValidatorsetCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorsetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorsetTransactorRaw struct {
	Contract *ValidatorsetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorset creates a new instance of Validatorset, bound to a specific deployed contract.
func NewValidatorset(address common.Address, backend bind.ContractBackend) (*Validatorset, error) {
	contract, err := bindValidatorset(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Validatorset{ValidatorsetCaller: ValidatorsetCaller{contract: contract}, ValidatorsetTransactor: ValidatorsetTransactor{contract: contract}, ValidatorsetFilterer: ValidatorsetFilterer{contract: contract}}, nil
}

// NewValidatorsetCaller creates a new read-only instance of Validatorset, bound to a specific deployed contract.
func NewValidatorsetCaller(address common.Address, caller bind.ContractCaller) (*ValidatorsetCaller, error) {
	contract, err := bindValidatorset(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorsetCaller{contract: contract}, nil
}

// NewValidatorsetTransactor creates a new write-only instance of Validatorset, bound to a specific deployed contract.
func NewValidatorsetTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorsetTransactor, error) {
	contract, err := bindValidatorset(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorsetTransactor{contract: contract}, nil
}

// NewValidatorsetFilterer creates a new log filterer instance of Validatorset, bound to a specific deployed contract.
func NewValidatorsetFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorsetFilterer, error) {
	contract, err := bindValidatorset(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorsetFilterer{contract: contract}, nil
}

// bindValidatorset binds a generic wrapper to an already deployed contract.
func bindValidatorset(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorsetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validatorset *ValidatorsetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Validatorset.Contract.ValidatorsetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validatorset *ValidatorsetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validatorset.Contract.ValidatorsetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validatorset *ValidatorsetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validatorset.Contract.ValidatorsetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validatorset *ValidatorsetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Validatorset.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validatorset *ValidatorsetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validatorset.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validatorset *ValidatorsetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validatorset.Contract.contract.Transact(opts, method, params...)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_Validatorset *ValidatorsetCaller) AlreadyInit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "alreadyInit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_Validatorset *ValidatorsetSession) AlreadyInit() (bool, error) {
	return _Validatorset.Contract.AlreadyInit(&_Validatorset.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_Validatorset *ValidatorsetCallerSession) AlreadyInit() (bool, error) {
	return _Validatorset.Contract.AlreadyInit(&_Validatorset.CallOpts)
}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x6969a25c.
//
// Solidity: function currentValidatorSet(uint256 ) view returns(address consensusAddress, uint256 stakeAmount, uint8 bondStatus, bool isJail)
func (_Validatorset *ValidatorsetCaller) CurrentValidatorSet(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ConsensusAddress common.Address
	StakeAmount      *big.Int
	BondStatus       uint8
	IsJail           bool
}, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "currentValidatorSet", arg0)

	outstruct := new(struct {
		ConsensusAddress common.Address
		StakeAmount      *big.Int
		BondStatus       uint8
		IsJail           bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConsensusAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.StakeAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BondStatus = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.IsJail = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x6969a25c.
//
// Solidity: function currentValidatorSet(uint256 ) view returns(address consensusAddress, uint256 stakeAmount, uint8 bondStatus, bool isJail)
func (_Validatorset *ValidatorsetSession) CurrentValidatorSet(arg0 *big.Int) (struct {
	ConsensusAddress common.Address
	StakeAmount      *big.Int
	BondStatus       uint8
	IsJail           bool
}, error) {
	return _Validatorset.Contract.CurrentValidatorSet(&_Validatorset.CallOpts, arg0)
}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x6969a25c.
//
// Solidity: function currentValidatorSet(uint256 ) view returns(address consensusAddress, uint256 stakeAmount, uint8 bondStatus, bool isJail)
func (_Validatorset *ValidatorsetCallerSession) CurrentValidatorSet(arg0 *big.Int) (struct {
	ConsensusAddress common.Address
	StakeAmount      *big.Int
	BondStatus       uint8
	IsJail           bool
}, error) {
	return _Validatorset.Contract.CurrentValidatorSet(&_Validatorset.CallOpts, arg0)
}

// CurrentValidatorSetMap is a free data retrieval call binding the contract method 0xad3c9da6.
//
// Solidity: function currentValidatorSetMap(address ) view returns(uint256)
func (_Validatorset *ValidatorsetCaller) CurrentValidatorSetMap(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "currentValidatorSetMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentValidatorSetMap is a free data retrieval call binding the contract method 0xad3c9da6.
//
// Solidity: function currentValidatorSetMap(address ) view returns(uint256)
func (_Validatorset *ValidatorsetSession) CurrentValidatorSetMap(arg0 common.Address) (*big.Int, error) {
	return _Validatorset.Contract.CurrentValidatorSetMap(&_Validatorset.CallOpts, arg0)
}

// CurrentValidatorSetMap is a free data retrieval call binding the contract method 0xad3c9da6.
//
// Solidity: function currentValidatorSetMap(address ) view returns(uint256)
func (_Validatorset *ValidatorsetCallerSession) CurrentValidatorSetMap(arg0 common.Address) (*big.Int, error) {
	return _Validatorset.Contract.CurrentValidatorSetMap(&_Validatorset.CallOpts, arg0)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Validatorset *ValidatorsetCaller) EndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "endTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Validatorset *ValidatorsetSession) EndTime() (*big.Int, error) {
	return _Validatorset.Contract.EndTime(&_Validatorset.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Validatorset *ValidatorsetCallerSession) EndTime() (*big.Int, error) {
	return _Validatorset.Contract.EndTime(&_Validatorset.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns((address,uint256,uint8,bool)[])
func (_Validatorset *ValidatorsetCaller) GetValidators(opts *bind.CallOpts) ([]IValidatorValidator, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "getValidators")

	if err != nil {
		return *new([]IValidatorValidator), err
	}

	out0 := *abi.ConvertType(out[0], new([]IValidatorValidator)).(*[]IValidatorValidator)

	return out0, err

}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns((address,uint256,uint8,bool)[])
func (_Validatorset *ValidatorsetSession) GetValidators() ([]IValidatorValidator, error) {
	return _Validatorset.Contract.GetValidators(&_Validatorset.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns((address,uint256,uint8,bool)[])
func (_Validatorset *ValidatorsetCallerSession) GetValidators() ([]IValidatorValidator, error) {
	return _Validatorset.Contract.GetValidators(&_Validatorset.CallOpts)
}

// NumberOfValidators is a free data retrieval call binding the contract method 0x74de1669.
//
// Solidity: function number_of_validators() view returns(uint256)
func (_Validatorset *ValidatorsetCaller) NumberOfValidators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validatorset.contract.Call(opts, &out, "number_of_validators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberOfValidators is a free data retrieval call binding the contract method 0x74de1669.
//
// Solidity: function number_of_validators() view returns(uint256)
func (_Validatorset *ValidatorsetSession) NumberOfValidators() (*big.Int, error) {
	return _Validatorset.Contract.NumberOfValidators(&_Validatorset.CallOpts)
}

// NumberOfValidators is a free data retrieval call binding the contract method 0x74de1669.
//
// Solidity: function number_of_validators() view returns(uint256)
func (_Validatorset *ValidatorsetCallerSession) NumberOfValidators() (*big.Int, error) {
	return _Validatorset.Contract.NumberOfValidators(&_Validatorset.CallOpts)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(address validatorPoolAddress, address systemRewardAddress) returns()
func (_Validatorset *ValidatorsetTransactor) Init(opts *bind.TransactOpts, validatorPoolAddress common.Address, systemRewardAddress common.Address) (*types.Transaction, error) {
	return _Validatorset.contract.Transact(opts, "init", validatorPoolAddress, systemRewardAddress)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(address validatorPoolAddress, address systemRewardAddress) returns()
func (_Validatorset *ValidatorsetSession) Init(validatorPoolAddress common.Address, systemRewardAddress common.Address) (*types.Transaction, error) {
	return _Validatorset.Contract.Init(&_Validatorset.TransactOpts, validatorPoolAddress, systemRewardAddress)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(address validatorPoolAddress, address systemRewardAddress) returns()
func (_Validatorset *ValidatorsetTransactorSession) Init(validatorPoolAddress common.Address, systemRewardAddress common.Address) (*types.Transaction, error) {
	return _Validatorset.Contract.Init(&_Validatorset.TransactOpts, validatorPoolAddress, systemRewardAddress)
}

// JailValidator is a paid mutator transaction binding the contract method 0xe589b61e.
//
// Solidity: function jailValidator(address consensusAddress) returns()
func (_Validatorset *ValidatorsetTransactor) JailValidator(opts *bind.TransactOpts, consensusAddress common.Address) (*types.Transaction, error) {
	return _Validatorset.contract.Transact(opts, "jailValidator", consensusAddress)
}

// JailValidator is a paid mutator transaction binding the contract method 0xe589b61e.
//
// Solidity: function jailValidator(address consensusAddress) returns()
func (_Validatorset *ValidatorsetSession) JailValidator(consensusAddress common.Address) (*types.Transaction, error) {
	return _Validatorset.Contract.JailValidator(&_Validatorset.TransactOpts, consensusAddress)
}

// JailValidator is a paid mutator transaction binding the contract method 0xe589b61e.
//
// Solidity: function jailValidator(address consensusAddress) returns()
func (_Validatorset *ValidatorsetTransactorSession) JailValidator(consensusAddress common.Address) (*types.Transaction, error) {
	return _Validatorset.Contract.JailValidator(&_Validatorset.TransactOpts, consensusAddress)
}

// UpdateValidatorSet is a paid mutator transaction binding the contract method 0x67d07b2f.
//
// Solidity: function updateValidatorSet() returns()
func (_Validatorset *ValidatorsetTransactor) UpdateValidatorSet(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validatorset.contract.Transact(opts, "updateValidatorSet")
}

// UpdateValidatorSet is a paid mutator transaction binding the contract method 0x67d07b2f.
//
// Solidity: function updateValidatorSet() returns()
func (_Validatorset *ValidatorsetSession) UpdateValidatorSet() (*types.Transaction, error) {
	return _Validatorset.Contract.UpdateValidatorSet(&_Validatorset.TransactOpts)
}

// UpdateValidatorSet is a paid mutator transaction binding the contract method 0x67d07b2f.
//
// Solidity: function updateValidatorSet() returns()
func (_Validatorset *ValidatorsetTransactorSession) UpdateValidatorSet() (*types.Transaction, error) {
	return _Validatorset.Contract.UpdateValidatorSet(&_Validatorset.TransactOpts)
}

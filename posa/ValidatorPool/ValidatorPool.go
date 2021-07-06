// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package validatorpool

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

// ValidatorpoolABI is the input ABI used to generate the binding from.
const ValidatorpoolABI = "[{\"inputs\":[],\"name\":\"NumberOfValidator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"alreadyInit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"consensusAddress\",\"type\":\"address\"}],\"name\":\"bondValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"consensusAddress\",\"type\":\"address\"}],\"name\":\"getTotalPower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"consensusAddress\",\"type\":\"address\"}],\"name\":\"getTotalPowerExcludeUnbonding\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"BKCValidatorSetAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"StakePoolAddress\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"consensusAddress\",\"type\":\"address\"}],\"name\":\"jailValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"consensusAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"operation\",\"type\":\"bool\"}],\"name\":\"rePositionValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registerValidator\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeJailValidatorFromQueue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeRemovingValidatorFromQueue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeUnBondingValidatorFromQueue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"consensusAddress\",\"type\":\"address\"}],\"name\":\"unBondValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validatorJailQueue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorRemove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validatorRemoveQueue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorTopUp\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validatorUnBondQueue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"consensusAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"enumIBond.BondStatus\",\"name\":\"bondStatus\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isJail\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validatorsMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Validatorpool is an auto generated Go binding around an Ethereum contract.
type Validatorpool struct {
	ValidatorpoolCaller     // Read-only binding to the contract
	ValidatorpoolTransactor // Write-only binding to the contract
	ValidatorpoolFilterer   // Log filterer for contract events
}

// ValidatorpoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorpoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorpoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorpoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorpoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorpoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorpoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorpoolSession struct {
	Contract     *Validatorpool    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorpoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorpoolCallerSession struct {
	Contract *ValidatorpoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ValidatorpoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorpoolTransactorSession struct {
	Contract     *ValidatorpoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ValidatorpoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorpoolRaw struct {
	Contract *Validatorpool // Generic contract binding to access the raw methods on
}

// ValidatorpoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorpoolCallerRaw struct {
	Contract *ValidatorpoolCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorpoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorpoolTransactorRaw struct {
	Contract *ValidatorpoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorpool creates a new instance of Validatorpool, bound to a specific deployed contract.
func NewValidatorpool(address common.Address, backend bind.ContractBackend) (*Validatorpool, error) {
	contract, err := bindValidatorpool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Validatorpool{ValidatorpoolCaller: ValidatorpoolCaller{contract: contract}, ValidatorpoolTransactor: ValidatorpoolTransactor{contract: contract}, ValidatorpoolFilterer: ValidatorpoolFilterer{contract: contract}}, nil
}

// NewValidatorpoolCaller creates a new read-only instance of Validatorpool, bound to a specific deployed contract.
func NewValidatorpoolCaller(address common.Address, caller bind.ContractCaller) (*ValidatorpoolCaller, error) {
	contract, err := bindValidatorpool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorpoolCaller{contract: contract}, nil
}

// NewValidatorpoolTransactor creates a new write-only instance of Validatorpool, bound to a specific deployed contract.
func NewValidatorpoolTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorpoolTransactor, error) {
	contract, err := bindValidatorpool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorpoolTransactor{contract: contract}, nil
}

// NewValidatorpoolFilterer creates a new log filterer instance of Validatorpool, bound to a specific deployed contract.
func NewValidatorpoolFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorpoolFilterer, error) {
	contract, err := bindValidatorpool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorpoolFilterer{contract: contract}, nil
}

// bindValidatorpool binds a generic wrapper to an already deployed contract.
func bindValidatorpool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorpoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validatorpool *ValidatorpoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Validatorpool.Contract.ValidatorpoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validatorpool *ValidatorpoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validatorpool.Contract.ValidatorpoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validatorpool *ValidatorpoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validatorpool.Contract.ValidatorpoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validatorpool *ValidatorpoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Validatorpool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validatorpool *ValidatorpoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validatorpool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validatorpool *ValidatorpoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validatorpool.Contract.contract.Transact(opts, method, params...)
}

// NumberOfValidator is a free data retrieval call binding the contract method 0x0c42f631.
//
// Solidity: function NumberOfValidator() view returns(uint256)
func (_Validatorpool *ValidatorpoolCaller) NumberOfValidator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validatorpool.contract.Call(opts, &out, "NumberOfValidator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberOfValidator is a free data retrieval call binding the contract method 0x0c42f631.
//
// Solidity: function NumberOfValidator() view returns(uint256)
func (_Validatorpool *ValidatorpoolSession) NumberOfValidator() (*big.Int, error) {
	return _Validatorpool.Contract.NumberOfValidator(&_Validatorpool.CallOpts)
}

// NumberOfValidator is a free data retrieval call binding the contract method 0x0c42f631.
//
// Solidity: function NumberOfValidator() view returns(uint256)
func (_Validatorpool *ValidatorpoolCallerSession) NumberOfValidator() (*big.Int, error) {
	return _Validatorpool.Contract.NumberOfValidator(&_Validatorpool.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_Validatorpool *ValidatorpoolCaller) AlreadyInit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Validatorpool.contract.Call(opts, &out, "alreadyInit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_Validatorpool *ValidatorpoolSession) AlreadyInit() (bool, error) {
	return _Validatorpool.Contract.AlreadyInit(&_Validatorpool.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_Validatorpool *ValidatorpoolCallerSession) AlreadyInit() (bool, error) {
	return _Validatorpool.Contract.AlreadyInit(&_Validatorpool.CallOpts)
}

// GetTotalPower is a free data retrieval call binding the contract method 0xd7dc81cd.
//
// Solidity: function getTotalPower(address consensusAddress) view returns(uint256)
func (_Validatorpool *ValidatorpoolCaller) GetTotalPower(opts *bind.CallOpts, consensusAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Validatorpool.contract.Call(opts, &out, "getTotalPower", consensusAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPower is a free data retrieval call binding the contract method 0xd7dc81cd.
//
// Solidity: function getTotalPower(address consensusAddress) view returns(uint256)
func (_Validatorpool *ValidatorpoolSession) GetTotalPower(consensusAddress common.Address) (*big.Int, error) {
	return _Validatorpool.Contract.GetTotalPower(&_Validatorpool.CallOpts, consensusAddress)
}

// GetTotalPower is a free data retrieval call binding the contract method 0xd7dc81cd.
//
// Solidity: function getTotalPower(address consensusAddress) view returns(uint256)
func (_Validatorpool *ValidatorpoolCallerSession) GetTotalPower(consensusAddress common.Address) (*big.Int, error) {
	return _Validatorpool.Contract.GetTotalPower(&_Validatorpool.CallOpts, consensusAddress)
}

// GetTotalPowerExcludeUnbonding is a free data retrieval call binding the contract method 0x322923b9.
//
// Solidity: function getTotalPowerExcludeUnbonding(address consensusAddress) view returns(uint256)
func (_Validatorpool *ValidatorpoolCaller) GetTotalPowerExcludeUnbonding(opts *bind.CallOpts, consensusAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Validatorpool.contract.Call(opts, &out, "getTotalPowerExcludeUnbonding", consensusAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPowerExcludeUnbonding is a free data retrieval call binding the contract method 0x322923b9.
//
// Solidity: function getTotalPowerExcludeUnbonding(address consensusAddress) view returns(uint256)
func (_Validatorpool *ValidatorpoolSession) GetTotalPowerExcludeUnbonding(consensusAddress common.Address) (*big.Int, error) {
	return _Validatorpool.Contract.GetTotalPowerExcludeUnbonding(&_Validatorpool.CallOpts, consensusAddress)
}

// GetTotalPowerExcludeUnbonding is a free data retrieval call binding the contract method 0x322923b9.
//
// Solidity: function getTotalPowerExcludeUnbonding(address consensusAddress) view returns(uint256)
func (_Validatorpool *ValidatorpoolCallerSession) GetTotalPowerExcludeUnbonding(consensusAddress common.Address) (*big.Int, error) {
	return _Validatorpool.Contract.GetTotalPowerExcludeUnbonding(&_Validatorpool.CallOpts, consensusAddress)
}

// ValidatorJailQueue is a free data retrieval call binding the contract method 0x727b62d8.
//
// Solidity: function validatorJailQueue(address ) view returns(uint256)
func (_Validatorpool *ValidatorpoolCaller) ValidatorJailQueue(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Validatorpool.contract.Call(opts, &out, "validatorJailQueue", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorJailQueue is a free data retrieval call binding the contract method 0x727b62d8.
//
// Solidity: function validatorJailQueue(address ) view returns(uint256)
func (_Validatorpool *ValidatorpoolSession) ValidatorJailQueue(arg0 common.Address) (*big.Int, error) {
	return _Validatorpool.Contract.ValidatorJailQueue(&_Validatorpool.CallOpts, arg0)
}

// ValidatorJailQueue is a free data retrieval call binding the contract method 0x727b62d8.
//
// Solidity: function validatorJailQueue(address ) view returns(uint256)
func (_Validatorpool *ValidatorpoolCallerSession) ValidatorJailQueue(arg0 common.Address) (*big.Int, error) {
	return _Validatorpool.Contract.ValidatorJailQueue(&_Validatorpool.CallOpts, arg0)
}

// ValidatorRemoveQueue is a free data retrieval call binding the contract method 0x50ef3309.
//
// Solidity: function validatorRemoveQueue(address ) view returns(uint256)
func (_Validatorpool *ValidatorpoolCaller) ValidatorRemoveQueue(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Validatorpool.contract.Call(opts, &out, "validatorRemoveQueue", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorRemoveQueue is a free data retrieval call binding the contract method 0x50ef3309.
//
// Solidity: function validatorRemoveQueue(address ) view returns(uint256)
func (_Validatorpool *ValidatorpoolSession) ValidatorRemoveQueue(arg0 common.Address) (*big.Int, error) {
	return _Validatorpool.Contract.ValidatorRemoveQueue(&_Validatorpool.CallOpts, arg0)
}

// ValidatorRemoveQueue is a free data retrieval call binding the contract method 0x50ef3309.
//
// Solidity: function validatorRemoveQueue(address ) view returns(uint256)
func (_Validatorpool *ValidatorpoolCallerSession) ValidatorRemoveQueue(arg0 common.Address) (*big.Int, error) {
	return _Validatorpool.Contract.ValidatorRemoveQueue(&_Validatorpool.CallOpts, arg0)
}

// ValidatorUnBondQueue is a free data retrieval call binding the contract method 0x099296a5.
//
// Solidity: function validatorUnBondQueue(address ) view returns(uint256)
func (_Validatorpool *ValidatorpoolCaller) ValidatorUnBondQueue(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Validatorpool.contract.Call(opts, &out, "validatorUnBondQueue", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorUnBondQueue is a free data retrieval call binding the contract method 0x099296a5.
//
// Solidity: function validatorUnBondQueue(address ) view returns(uint256)
func (_Validatorpool *ValidatorpoolSession) ValidatorUnBondQueue(arg0 common.Address) (*big.Int, error) {
	return _Validatorpool.Contract.ValidatorUnBondQueue(&_Validatorpool.CallOpts, arg0)
}

// ValidatorUnBondQueue is a free data retrieval call binding the contract method 0x099296a5.
//
// Solidity: function validatorUnBondQueue(address ) view returns(uint256)
func (_Validatorpool *ValidatorpoolCallerSession) ValidatorUnBondQueue(arg0 common.Address) (*big.Int, error) {
	return _Validatorpool.Contract.ValidatorUnBondQueue(&_Validatorpool.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 ) view returns(address consensusAddress, uint256 stakeAmount, uint8 bondStatus, bool isJail)
func (_Validatorpool *ValidatorpoolCaller) Validators(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ConsensusAddress common.Address
	StakeAmount      *big.Int
	BondStatus       uint8
	IsJail           bool
}, error) {
	var out []interface{}
	err := _Validatorpool.contract.Call(opts, &out, "validators", arg0)

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

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 ) view returns(address consensusAddress, uint256 stakeAmount, uint8 bondStatus, bool isJail)
func (_Validatorpool *ValidatorpoolSession) Validators(arg0 *big.Int) (struct {
	ConsensusAddress common.Address
	StakeAmount      *big.Int
	BondStatus       uint8
	IsJail           bool
}, error) {
	return _Validatorpool.Contract.Validators(&_Validatorpool.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators(uint256 ) view returns(address consensusAddress, uint256 stakeAmount, uint8 bondStatus, bool isJail)
func (_Validatorpool *ValidatorpoolCallerSession) Validators(arg0 *big.Int) (struct {
	ConsensusAddress common.Address
	StakeAmount      *big.Int
	BondStatus       uint8
	IsJail           bool
}, error) {
	return _Validatorpool.Contract.Validators(&_Validatorpool.CallOpts, arg0)
}

// ValidatorsMap is a free data retrieval call binding the contract method 0xb6338917.
//
// Solidity: function validatorsMap(address ) view returns(uint256)
func (_Validatorpool *ValidatorpoolCaller) ValidatorsMap(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Validatorpool.contract.Call(opts, &out, "validatorsMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorsMap is a free data retrieval call binding the contract method 0xb6338917.
//
// Solidity: function validatorsMap(address ) view returns(uint256)
func (_Validatorpool *ValidatorpoolSession) ValidatorsMap(arg0 common.Address) (*big.Int, error) {
	return _Validatorpool.Contract.ValidatorsMap(&_Validatorpool.CallOpts, arg0)
}

// ValidatorsMap is a free data retrieval call binding the contract method 0xb6338917.
//
// Solidity: function validatorsMap(address ) view returns(uint256)
func (_Validatorpool *ValidatorpoolCallerSession) ValidatorsMap(arg0 common.Address) (*big.Int, error) {
	return _Validatorpool.Contract.ValidatorsMap(&_Validatorpool.CallOpts, arg0)
}

// BondValidator is a paid mutator transaction binding the contract method 0x360d80ec.
//
// Solidity: function bondValidator(address consensusAddress) returns()
func (_Validatorpool *ValidatorpoolTransactor) BondValidator(opts *bind.TransactOpts, consensusAddress common.Address) (*types.Transaction, error) {
	return _Validatorpool.contract.Transact(opts, "bondValidator", consensusAddress)
}

// BondValidator is a paid mutator transaction binding the contract method 0x360d80ec.
//
// Solidity: function bondValidator(address consensusAddress) returns()
func (_Validatorpool *ValidatorpoolSession) BondValidator(consensusAddress common.Address) (*types.Transaction, error) {
	return _Validatorpool.Contract.BondValidator(&_Validatorpool.TransactOpts, consensusAddress)
}

// BondValidator is a paid mutator transaction binding the contract method 0x360d80ec.
//
// Solidity: function bondValidator(address consensusAddress) returns()
func (_Validatorpool *ValidatorpoolTransactorSession) BondValidator(consensusAddress common.Address) (*types.Transaction, error) {
	return _Validatorpool.Contract.BondValidator(&_Validatorpool.TransactOpts, consensusAddress)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(address BKCValidatorSetAddress, address StakePoolAddress) payable returns()
func (_Validatorpool *ValidatorpoolTransactor) Init(opts *bind.TransactOpts, BKCValidatorSetAddress common.Address, StakePoolAddress common.Address) (*types.Transaction, error) {
	return _Validatorpool.contract.Transact(opts, "init", BKCValidatorSetAddress, StakePoolAddress)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(address BKCValidatorSetAddress, address StakePoolAddress) payable returns()
func (_Validatorpool *ValidatorpoolSession) Init(BKCValidatorSetAddress common.Address, StakePoolAddress common.Address) (*types.Transaction, error) {
	return _Validatorpool.Contract.Init(&_Validatorpool.TransactOpts, BKCValidatorSetAddress, StakePoolAddress)
}

// Init is a paid mutator transaction binding the contract method 0xf09a4016.
//
// Solidity: function init(address BKCValidatorSetAddress, address StakePoolAddress) payable returns()
func (_Validatorpool *ValidatorpoolTransactorSession) Init(BKCValidatorSetAddress common.Address, StakePoolAddress common.Address) (*types.Transaction, error) {
	return _Validatorpool.Contract.Init(&_Validatorpool.TransactOpts, BKCValidatorSetAddress, StakePoolAddress)
}

// JailValidator is a paid mutator transaction binding the contract method 0xe589b61e.
//
// Solidity: function jailValidator(address consensusAddress) returns()
func (_Validatorpool *ValidatorpoolTransactor) JailValidator(opts *bind.TransactOpts, consensusAddress common.Address) (*types.Transaction, error) {
	return _Validatorpool.contract.Transact(opts, "jailValidator", consensusAddress)
}

// JailValidator is a paid mutator transaction binding the contract method 0xe589b61e.
//
// Solidity: function jailValidator(address consensusAddress) returns()
func (_Validatorpool *ValidatorpoolSession) JailValidator(consensusAddress common.Address) (*types.Transaction, error) {
	return _Validatorpool.Contract.JailValidator(&_Validatorpool.TransactOpts, consensusAddress)
}

// JailValidator is a paid mutator transaction binding the contract method 0xe589b61e.
//
// Solidity: function jailValidator(address consensusAddress) returns()
func (_Validatorpool *ValidatorpoolTransactorSession) JailValidator(consensusAddress common.Address) (*types.Transaction, error) {
	return _Validatorpool.Contract.JailValidator(&_Validatorpool.TransactOpts, consensusAddress)
}

// RePositionValidator is a paid mutator transaction binding the contract method 0x3004d42a.
//
// Solidity: function rePositionValidator(address consensusAddress, bool operation) returns()
func (_Validatorpool *ValidatorpoolTransactor) RePositionValidator(opts *bind.TransactOpts, consensusAddress common.Address, operation bool) (*types.Transaction, error) {
	return _Validatorpool.contract.Transact(opts, "rePositionValidator", consensusAddress, operation)
}

// RePositionValidator is a paid mutator transaction binding the contract method 0x3004d42a.
//
// Solidity: function rePositionValidator(address consensusAddress, bool operation) returns()
func (_Validatorpool *ValidatorpoolSession) RePositionValidator(consensusAddress common.Address, operation bool) (*types.Transaction, error) {
	return _Validatorpool.Contract.RePositionValidator(&_Validatorpool.TransactOpts, consensusAddress, operation)
}

// RePositionValidator is a paid mutator transaction binding the contract method 0x3004d42a.
//
// Solidity: function rePositionValidator(address consensusAddress, bool operation) returns()
func (_Validatorpool *ValidatorpoolTransactorSession) RePositionValidator(consensusAddress common.Address, operation bool) (*types.Transaction, error) {
	return _Validatorpool.Contract.RePositionValidator(&_Validatorpool.TransactOpts, consensusAddress, operation)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0xbcc6587f.
//
// Solidity: function registerValidator() payable returns()
func (_Validatorpool *ValidatorpoolTransactor) RegisterValidator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validatorpool.contract.Transact(opts, "registerValidator")
}

// RegisterValidator is a paid mutator transaction binding the contract method 0xbcc6587f.
//
// Solidity: function registerValidator() payable returns()
func (_Validatorpool *ValidatorpoolSession) RegisterValidator() (*types.Transaction, error) {
	return _Validatorpool.Contract.RegisterValidator(&_Validatorpool.TransactOpts)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0xbcc6587f.
//
// Solidity: function registerValidator() payable returns()
func (_Validatorpool *ValidatorpoolTransactorSession) RegisterValidator() (*types.Transaction, error) {
	return _Validatorpool.Contract.RegisterValidator(&_Validatorpool.TransactOpts)
}

// RemoveJailValidatorFromQueue is a paid mutator transaction binding the contract method 0x45c5cdba.
//
// Solidity: function removeJailValidatorFromQueue() returns()
func (_Validatorpool *ValidatorpoolTransactor) RemoveJailValidatorFromQueue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validatorpool.contract.Transact(opts, "removeJailValidatorFromQueue")
}

// RemoveJailValidatorFromQueue is a paid mutator transaction binding the contract method 0x45c5cdba.
//
// Solidity: function removeJailValidatorFromQueue() returns()
func (_Validatorpool *ValidatorpoolSession) RemoveJailValidatorFromQueue() (*types.Transaction, error) {
	return _Validatorpool.Contract.RemoveJailValidatorFromQueue(&_Validatorpool.TransactOpts)
}

// RemoveJailValidatorFromQueue is a paid mutator transaction binding the contract method 0x45c5cdba.
//
// Solidity: function removeJailValidatorFromQueue() returns()
func (_Validatorpool *ValidatorpoolTransactorSession) RemoveJailValidatorFromQueue() (*types.Transaction, error) {
	return _Validatorpool.Contract.RemoveJailValidatorFromQueue(&_Validatorpool.TransactOpts)
}

// RemoveRemovingValidatorFromQueue is a paid mutator transaction binding the contract method 0x7ab05c49.
//
// Solidity: function removeRemovingValidatorFromQueue() returns()
func (_Validatorpool *ValidatorpoolTransactor) RemoveRemovingValidatorFromQueue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validatorpool.contract.Transact(opts, "removeRemovingValidatorFromQueue")
}

// RemoveRemovingValidatorFromQueue is a paid mutator transaction binding the contract method 0x7ab05c49.
//
// Solidity: function removeRemovingValidatorFromQueue() returns()
func (_Validatorpool *ValidatorpoolSession) RemoveRemovingValidatorFromQueue() (*types.Transaction, error) {
	return _Validatorpool.Contract.RemoveRemovingValidatorFromQueue(&_Validatorpool.TransactOpts)
}

// RemoveRemovingValidatorFromQueue is a paid mutator transaction binding the contract method 0x7ab05c49.
//
// Solidity: function removeRemovingValidatorFromQueue() returns()
func (_Validatorpool *ValidatorpoolTransactorSession) RemoveRemovingValidatorFromQueue() (*types.Transaction, error) {
	return _Validatorpool.Contract.RemoveRemovingValidatorFromQueue(&_Validatorpool.TransactOpts)
}

// RemoveUnBondingValidatorFromQueue is a paid mutator transaction binding the contract method 0x77ef7680.
//
// Solidity: function removeUnBondingValidatorFromQueue() returns()
func (_Validatorpool *ValidatorpoolTransactor) RemoveUnBondingValidatorFromQueue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validatorpool.contract.Transact(opts, "removeUnBondingValidatorFromQueue")
}

// RemoveUnBondingValidatorFromQueue is a paid mutator transaction binding the contract method 0x77ef7680.
//
// Solidity: function removeUnBondingValidatorFromQueue() returns()
func (_Validatorpool *ValidatorpoolSession) RemoveUnBondingValidatorFromQueue() (*types.Transaction, error) {
	return _Validatorpool.Contract.RemoveUnBondingValidatorFromQueue(&_Validatorpool.TransactOpts)
}

// RemoveUnBondingValidatorFromQueue is a paid mutator transaction binding the contract method 0x77ef7680.
//
// Solidity: function removeUnBondingValidatorFromQueue() returns()
func (_Validatorpool *ValidatorpoolTransactorSession) RemoveUnBondingValidatorFromQueue() (*types.Transaction, error) {
	return _Validatorpool.Contract.RemoveUnBondingValidatorFromQueue(&_Validatorpool.TransactOpts)
}

// UnBondValidator is a paid mutator transaction binding the contract method 0x41eeacdb.
//
// Solidity: function unBondValidator(address consensusAddress) returns()
func (_Validatorpool *ValidatorpoolTransactor) UnBondValidator(opts *bind.TransactOpts, consensusAddress common.Address) (*types.Transaction, error) {
	return _Validatorpool.contract.Transact(opts, "unBondValidator", consensusAddress)
}

// UnBondValidator is a paid mutator transaction binding the contract method 0x41eeacdb.
//
// Solidity: function unBondValidator(address consensusAddress) returns()
func (_Validatorpool *ValidatorpoolSession) UnBondValidator(consensusAddress common.Address) (*types.Transaction, error) {
	return _Validatorpool.Contract.UnBondValidator(&_Validatorpool.TransactOpts, consensusAddress)
}

// UnBondValidator is a paid mutator transaction binding the contract method 0x41eeacdb.
//
// Solidity: function unBondValidator(address consensusAddress) returns()
func (_Validatorpool *ValidatorpoolTransactorSession) UnBondValidator(consensusAddress common.Address) (*types.Transaction, error) {
	return _Validatorpool.Contract.UnBondValidator(&_Validatorpool.TransactOpts, consensusAddress)
}

// ValidatorRemove is a paid mutator transaction binding the contract method 0x33a92e35.
//
// Solidity: function validatorRemove() returns()
func (_Validatorpool *ValidatorpoolTransactor) ValidatorRemove(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validatorpool.contract.Transact(opts, "validatorRemove")
}

// ValidatorRemove is a paid mutator transaction binding the contract method 0x33a92e35.
//
// Solidity: function validatorRemove() returns()
func (_Validatorpool *ValidatorpoolSession) ValidatorRemove() (*types.Transaction, error) {
	return _Validatorpool.Contract.ValidatorRemove(&_Validatorpool.TransactOpts)
}

// ValidatorRemove is a paid mutator transaction binding the contract method 0x33a92e35.
//
// Solidity: function validatorRemove() returns()
func (_Validatorpool *ValidatorpoolTransactorSession) ValidatorRemove() (*types.Transaction, error) {
	return _Validatorpool.Contract.ValidatorRemove(&_Validatorpool.TransactOpts)
}

// ValidatorTopUp is a paid mutator transaction binding the contract method 0x33fd2e05.
//
// Solidity: function validatorTopUp() payable returns()
func (_Validatorpool *ValidatorpoolTransactor) ValidatorTopUp(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validatorpool.contract.Transact(opts, "validatorTopUp")
}

// ValidatorTopUp is a paid mutator transaction binding the contract method 0x33fd2e05.
//
// Solidity: function validatorTopUp() payable returns()
func (_Validatorpool *ValidatorpoolSession) ValidatorTopUp() (*types.Transaction, error) {
	return _Validatorpool.Contract.ValidatorTopUp(&_Validatorpool.TransactOpts)
}

// ValidatorTopUp is a paid mutator transaction binding the contract method 0x33fd2e05.
//
// Solidity: function validatorTopUp() payable returns()
func (_Validatorpool *ValidatorpoolTransactorSession) ValidatorTopUp() (*types.Transaction, error) {
	return _Validatorpool.Contract.ValidatorTopUp(&_Validatorpool.TransactOpts)
}

// WithdrawFund is a paid mutator transaction binding the contract method 0x0cee1725.
//
// Solidity: function withdrawFund(uint256 amount) returns()
func (_Validatorpool *ValidatorpoolTransactor) WithdrawFund(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Validatorpool.contract.Transact(opts, "withdrawFund", amount)
}

// WithdrawFund is a paid mutator transaction binding the contract method 0x0cee1725.
//
// Solidity: function withdrawFund(uint256 amount) returns()
func (_Validatorpool *ValidatorpoolSession) WithdrawFund(amount *big.Int) (*types.Transaction, error) {
	return _Validatorpool.Contract.WithdrawFund(&_Validatorpool.TransactOpts, amount)
}

// WithdrawFund is a paid mutator transaction binding the contract method 0x0cee1725.
//
// Solidity: function withdrawFund(uint256 amount) returns()
func (_Validatorpool *ValidatorpoolTransactorSession) WithdrawFund(amount *big.Int) (*types.Transaction, error) {
	return _Validatorpool.Contract.WithdrawFund(&_Validatorpool.TransactOpts, amount)
}

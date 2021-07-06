// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stakepool

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

// StakepoolABI is the input ABI used to generate the binding from.
const StakepoolABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"UserUnDelegateQueue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ValidatorDelegation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalDelegation\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"alreadyInit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorConsensus\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"consensusAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"getDelegationAmoountOfEach\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"consensusAddress\",\"type\":\"address\"}],\"name\":\"getDelegators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"consensusAddress\",\"type\":\"address\"}],\"name\":\"getTotalDelegation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"consensusAddress\",\"type\":\"address\"}],\"name\":\"getTotalDelegationExcludeUnbonding\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"getUnbondingQueueAmount\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"getUnbondingQueueTime\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"getUnbondingQueueValidator\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"consensusAddress\",\"type\":\"address\"}],\"name\":\"getUserDelegationBondedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"consensusAddress\",\"type\":\"address\"}],\"name\":\"getUserDelegationBondedAmountCallable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"consensusAddress\",\"type\":\"address\"}],\"name\":\"getUserDelegationUnbondingAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"consensusAddress\",\"type\":\"address\"}],\"name\":\"getUserDelegationUnbondingAmountCallable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"consensusAddress\",\"type\":\"address\"}],\"name\":\"getUserTotalDelegationAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"consensusAddress\",\"type\":\"address\"}],\"name\":\"getUserTotalDelegationAmountCallable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorPoolAddress\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"removeDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeUnbondingUserFromUnbondingQueue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorConsensus\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"undelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Stakepool is an auto generated Go binding around an Ethereum contract.
type Stakepool struct {
	StakepoolCaller     // Read-only binding to the contract
	StakepoolTransactor // Write-only binding to the contract
	StakepoolFilterer   // Log filterer for contract events
}

// StakepoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakepoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakepoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakepoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakepoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakepoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakepoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakepoolSession struct {
	Contract     *Stakepool        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakepoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakepoolCallerSession struct {
	Contract *StakepoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// StakepoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakepoolTransactorSession struct {
	Contract     *StakepoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// StakepoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakepoolRaw struct {
	Contract *Stakepool // Generic contract binding to access the raw methods on
}

// StakepoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakepoolCallerRaw struct {
	Contract *StakepoolCaller // Generic read-only contract binding to access the raw methods on
}

// StakepoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakepoolTransactorRaw struct {
	Contract *StakepoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakepool creates a new instance of Stakepool, bound to a specific deployed contract.
func NewStakepool(address common.Address, backend bind.ContractBackend) (*Stakepool, error) {
	contract, err := bindStakepool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stakepool{StakepoolCaller: StakepoolCaller{contract: contract}, StakepoolTransactor: StakepoolTransactor{contract: contract}, StakepoolFilterer: StakepoolFilterer{contract: contract}}, nil
}

// NewStakepoolCaller creates a new read-only instance of Stakepool, bound to a specific deployed contract.
func NewStakepoolCaller(address common.Address, caller bind.ContractCaller) (*StakepoolCaller, error) {
	contract, err := bindStakepool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakepoolCaller{contract: contract}, nil
}

// NewStakepoolTransactor creates a new write-only instance of Stakepool, bound to a specific deployed contract.
func NewStakepoolTransactor(address common.Address, transactor bind.ContractTransactor) (*StakepoolTransactor, error) {
	contract, err := bindStakepool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakepoolTransactor{contract: contract}, nil
}

// NewStakepoolFilterer creates a new log filterer instance of Stakepool, bound to a specific deployed contract.
func NewStakepoolFilterer(address common.Address, filterer bind.ContractFilterer) (*StakepoolFilterer, error) {
	contract, err := bindStakepool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakepoolFilterer{contract: contract}, nil
}

// bindStakepool binds a generic wrapper to an already deployed contract.
func bindStakepool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakepoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stakepool *StakepoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stakepool.Contract.StakepoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stakepool *StakepoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakepool.Contract.StakepoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stakepool *StakepoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stakepool.Contract.StakepoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stakepool *StakepoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stakepool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stakepool *StakepoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakepool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stakepool *StakepoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stakepool.Contract.contract.Transact(opts, method, params...)
}

// UserUnDelegateQueue is a free data retrieval call binding the contract method 0x0717bdbb.
//
// Solidity: function UserUnDelegateQueue(address , uint256 ) view returns(uint256 amount, uint256 time, address validator)
func (_Stakepool *StakepoolCaller) UserUnDelegateQueue(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Amount    *big.Int
	Time      *big.Int
	Validator common.Address
}, error) {
	var out []interface{}
	err := _Stakepool.contract.Call(opts, &out, "UserUnDelegateQueue", arg0, arg1)

	outstruct := new(struct {
		Amount    *big.Int
		Time      *big.Int
		Validator common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Time = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Validator = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// UserUnDelegateQueue is a free data retrieval call binding the contract method 0x0717bdbb.
//
// Solidity: function UserUnDelegateQueue(address , uint256 ) view returns(uint256 amount, uint256 time, address validator)
func (_Stakepool *StakepoolSession) UserUnDelegateQueue(arg0 common.Address, arg1 *big.Int) (struct {
	Amount    *big.Int
	Time      *big.Int
	Validator common.Address
}, error) {
	return _Stakepool.Contract.UserUnDelegateQueue(&_Stakepool.CallOpts, arg0, arg1)
}

// UserUnDelegateQueue is a free data retrieval call binding the contract method 0x0717bdbb.
//
// Solidity: function UserUnDelegateQueue(address , uint256 ) view returns(uint256 amount, uint256 time, address validator)
func (_Stakepool *StakepoolCallerSession) UserUnDelegateQueue(arg0 common.Address, arg1 *big.Int) (struct {
	Amount    *big.Int
	Time      *big.Int
	Validator common.Address
}, error) {
	return _Stakepool.Contract.UserUnDelegateQueue(&_Stakepool.CallOpts, arg0, arg1)
}

// ValidatorDelegation is a free data retrieval call binding the contract method 0x0c4432d9.
//
// Solidity: function ValidatorDelegation(address ) view returns(uint256 totalDelegation)
func (_Stakepool *StakepoolCaller) ValidatorDelegation(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stakepool.contract.Call(opts, &out, "ValidatorDelegation", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorDelegation is a free data retrieval call binding the contract method 0x0c4432d9.
//
// Solidity: function ValidatorDelegation(address ) view returns(uint256 totalDelegation)
func (_Stakepool *StakepoolSession) ValidatorDelegation(arg0 common.Address) (*big.Int, error) {
	return _Stakepool.Contract.ValidatorDelegation(&_Stakepool.CallOpts, arg0)
}

// ValidatorDelegation is a free data retrieval call binding the contract method 0x0c4432d9.
//
// Solidity: function ValidatorDelegation(address ) view returns(uint256 totalDelegation)
func (_Stakepool *StakepoolCallerSession) ValidatorDelegation(arg0 common.Address) (*big.Int, error) {
	return _Stakepool.Contract.ValidatorDelegation(&_Stakepool.CallOpts, arg0)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_Stakepool *StakepoolCaller) AlreadyInit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Stakepool.contract.Call(opts, &out, "alreadyInit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_Stakepool *StakepoolSession) AlreadyInit() (bool, error) {
	return _Stakepool.Contract.AlreadyInit(&_Stakepool.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_Stakepool *StakepoolCallerSession) AlreadyInit() (bool, error) {
	return _Stakepool.Contract.AlreadyInit(&_Stakepool.CallOpts)
}

// GetDelegationAmoountOfEach is a free data retrieval call binding the contract method 0xff70b10e.
//
// Solidity: function getDelegationAmoountOfEach(address consensusAddress, address delegator) view returns(uint256)
func (_Stakepool *StakepoolCaller) GetDelegationAmoountOfEach(opts *bind.CallOpts, consensusAddress common.Address, delegator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stakepool.contract.Call(opts, &out, "getDelegationAmoountOfEach", consensusAddress, delegator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDelegationAmoountOfEach is a free data retrieval call binding the contract method 0xff70b10e.
//
// Solidity: function getDelegationAmoountOfEach(address consensusAddress, address delegator) view returns(uint256)
func (_Stakepool *StakepoolSession) GetDelegationAmoountOfEach(consensusAddress common.Address, delegator common.Address) (*big.Int, error) {
	return _Stakepool.Contract.GetDelegationAmoountOfEach(&_Stakepool.CallOpts, consensusAddress, delegator)
}

// GetDelegationAmoountOfEach is a free data retrieval call binding the contract method 0xff70b10e.
//
// Solidity: function getDelegationAmoountOfEach(address consensusAddress, address delegator) view returns(uint256)
func (_Stakepool *StakepoolCallerSession) GetDelegationAmoountOfEach(consensusAddress common.Address, delegator common.Address) (*big.Int, error) {
	return _Stakepool.Contract.GetDelegationAmoountOfEach(&_Stakepool.CallOpts, consensusAddress, delegator)
}

// GetDelegators is a free data retrieval call binding the contract method 0x68e76346.
//
// Solidity: function getDelegators(address consensusAddress) view returns(address[])
func (_Stakepool *StakepoolCaller) GetDelegators(opts *bind.CallOpts, consensusAddress common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Stakepool.contract.Call(opts, &out, "getDelegators", consensusAddress)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetDelegators is a free data retrieval call binding the contract method 0x68e76346.
//
// Solidity: function getDelegators(address consensusAddress) view returns(address[])
func (_Stakepool *StakepoolSession) GetDelegators(consensusAddress common.Address) ([]common.Address, error) {
	return _Stakepool.Contract.GetDelegators(&_Stakepool.CallOpts, consensusAddress)
}

// GetDelegators is a free data retrieval call binding the contract method 0x68e76346.
//
// Solidity: function getDelegators(address consensusAddress) view returns(address[])
func (_Stakepool *StakepoolCallerSession) GetDelegators(consensusAddress common.Address) ([]common.Address, error) {
	return _Stakepool.Contract.GetDelegators(&_Stakepool.CallOpts, consensusAddress)
}

// GetTotalDelegation is a free data retrieval call binding the contract method 0xfc5e7e09.
//
// Solidity: function getTotalDelegation(address consensusAddress) view returns(uint256)
func (_Stakepool *StakepoolCaller) GetTotalDelegation(opts *bind.CallOpts, consensusAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stakepool.contract.Call(opts, &out, "getTotalDelegation", consensusAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalDelegation is a free data retrieval call binding the contract method 0xfc5e7e09.
//
// Solidity: function getTotalDelegation(address consensusAddress) view returns(uint256)
func (_Stakepool *StakepoolSession) GetTotalDelegation(consensusAddress common.Address) (*big.Int, error) {
	return _Stakepool.Contract.GetTotalDelegation(&_Stakepool.CallOpts, consensusAddress)
}

// GetTotalDelegation is a free data retrieval call binding the contract method 0xfc5e7e09.
//
// Solidity: function getTotalDelegation(address consensusAddress) view returns(uint256)
func (_Stakepool *StakepoolCallerSession) GetTotalDelegation(consensusAddress common.Address) (*big.Int, error) {
	return _Stakepool.Contract.GetTotalDelegation(&_Stakepool.CallOpts, consensusAddress)
}

// GetTotalDelegationExcludeUnbonding is a free data retrieval call binding the contract method 0xfdbd463d.
//
// Solidity: function getTotalDelegationExcludeUnbonding(address consensusAddress) view returns(uint256)
func (_Stakepool *StakepoolCaller) GetTotalDelegationExcludeUnbonding(opts *bind.CallOpts, consensusAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stakepool.contract.Call(opts, &out, "getTotalDelegationExcludeUnbonding", consensusAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalDelegationExcludeUnbonding is a free data retrieval call binding the contract method 0xfdbd463d.
//
// Solidity: function getTotalDelegationExcludeUnbonding(address consensusAddress) view returns(uint256)
func (_Stakepool *StakepoolSession) GetTotalDelegationExcludeUnbonding(consensusAddress common.Address) (*big.Int, error) {
	return _Stakepool.Contract.GetTotalDelegationExcludeUnbonding(&_Stakepool.CallOpts, consensusAddress)
}

// GetTotalDelegationExcludeUnbonding is a free data retrieval call binding the contract method 0xfdbd463d.
//
// Solidity: function getTotalDelegationExcludeUnbonding(address consensusAddress) view returns(uint256)
func (_Stakepool *StakepoolCallerSession) GetTotalDelegationExcludeUnbonding(consensusAddress common.Address) (*big.Int, error) {
	return _Stakepool.Contract.GetTotalDelegationExcludeUnbonding(&_Stakepool.CallOpts, consensusAddress)
}

// GetUserDelegationBondedAmount is a free data retrieval call binding the contract method 0x08ec40d0.
//
// Solidity: function getUserDelegationBondedAmount(address consensusAddress) view returns(uint256)
func (_Stakepool *StakepoolCaller) GetUserDelegationBondedAmount(opts *bind.CallOpts, consensusAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stakepool.contract.Call(opts, &out, "getUserDelegationBondedAmount", consensusAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserDelegationBondedAmount is a free data retrieval call binding the contract method 0x08ec40d0.
//
// Solidity: function getUserDelegationBondedAmount(address consensusAddress) view returns(uint256)
func (_Stakepool *StakepoolSession) GetUserDelegationBondedAmount(consensusAddress common.Address) (*big.Int, error) {
	return _Stakepool.Contract.GetUserDelegationBondedAmount(&_Stakepool.CallOpts, consensusAddress)
}

// GetUserDelegationBondedAmount is a free data retrieval call binding the contract method 0x08ec40d0.
//
// Solidity: function getUserDelegationBondedAmount(address consensusAddress) view returns(uint256)
func (_Stakepool *StakepoolCallerSession) GetUserDelegationBondedAmount(consensusAddress common.Address) (*big.Int, error) {
	return _Stakepool.Contract.GetUserDelegationBondedAmount(&_Stakepool.CallOpts, consensusAddress)
}

// GetUserDelegationBondedAmountCallable is a free data retrieval call binding the contract method 0xf9353fe0.
//
// Solidity: function getUserDelegationBondedAmountCallable(address delegator, address consensusAddress) view returns(uint256)
func (_Stakepool *StakepoolCaller) GetUserDelegationBondedAmountCallable(opts *bind.CallOpts, delegator common.Address, consensusAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stakepool.contract.Call(opts, &out, "getUserDelegationBondedAmountCallable", delegator, consensusAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserDelegationBondedAmountCallable is a free data retrieval call binding the contract method 0xf9353fe0.
//
// Solidity: function getUserDelegationBondedAmountCallable(address delegator, address consensusAddress) view returns(uint256)
func (_Stakepool *StakepoolSession) GetUserDelegationBondedAmountCallable(delegator common.Address, consensusAddress common.Address) (*big.Int, error) {
	return _Stakepool.Contract.GetUserDelegationBondedAmountCallable(&_Stakepool.CallOpts, delegator, consensusAddress)
}

// GetUserDelegationBondedAmountCallable is a free data retrieval call binding the contract method 0xf9353fe0.
//
// Solidity: function getUserDelegationBondedAmountCallable(address delegator, address consensusAddress) view returns(uint256)
func (_Stakepool *StakepoolCallerSession) GetUserDelegationBondedAmountCallable(delegator common.Address, consensusAddress common.Address) (*big.Int, error) {
	return _Stakepool.Contract.GetUserDelegationBondedAmountCallable(&_Stakepool.CallOpts, delegator, consensusAddress)
}

// GetUserDelegationUnbondingAmount is a free data retrieval call binding the contract method 0x1802f311.
//
// Solidity: function getUserDelegationUnbondingAmount(address consensusAddress) view returns(uint256)
func (_Stakepool *StakepoolCaller) GetUserDelegationUnbondingAmount(opts *bind.CallOpts, consensusAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stakepool.contract.Call(opts, &out, "getUserDelegationUnbondingAmount", consensusAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserDelegationUnbondingAmount is a free data retrieval call binding the contract method 0x1802f311.
//
// Solidity: function getUserDelegationUnbondingAmount(address consensusAddress) view returns(uint256)
func (_Stakepool *StakepoolSession) GetUserDelegationUnbondingAmount(consensusAddress common.Address) (*big.Int, error) {
	return _Stakepool.Contract.GetUserDelegationUnbondingAmount(&_Stakepool.CallOpts, consensusAddress)
}

// GetUserDelegationUnbondingAmount is a free data retrieval call binding the contract method 0x1802f311.
//
// Solidity: function getUserDelegationUnbondingAmount(address consensusAddress) view returns(uint256)
func (_Stakepool *StakepoolCallerSession) GetUserDelegationUnbondingAmount(consensusAddress common.Address) (*big.Int, error) {
	return _Stakepool.Contract.GetUserDelegationUnbondingAmount(&_Stakepool.CallOpts, consensusAddress)
}

// GetUserDelegationUnbondingAmountCallable is a free data retrieval call binding the contract method 0x99e116ac.
//
// Solidity: function getUserDelegationUnbondingAmountCallable(address delegator, address consensusAddress) view returns(uint256)
func (_Stakepool *StakepoolCaller) GetUserDelegationUnbondingAmountCallable(opts *bind.CallOpts, delegator common.Address, consensusAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stakepool.contract.Call(opts, &out, "getUserDelegationUnbondingAmountCallable", delegator, consensusAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserDelegationUnbondingAmountCallable is a free data retrieval call binding the contract method 0x99e116ac.
//
// Solidity: function getUserDelegationUnbondingAmountCallable(address delegator, address consensusAddress) view returns(uint256)
func (_Stakepool *StakepoolSession) GetUserDelegationUnbondingAmountCallable(delegator common.Address, consensusAddress common.Address) (*big.Int, error) {
	return _Stakepool.Contract.GetUserDelegationUnbondingAmountCallable(&_Stakepool.CallOpts, delegator, consensusAddress)
}

// GetUserDelegationUnbondingAmountCallable is a free data retrieval call binding the contract method 0x99e116ac.
//
// Solidity: function getUserDelegationUnbondingAmountCallable(address delegator, address consensusAddress) view returns(uint256)
func (_Stakepool *StakepoolCallerSession) GetUserDelegationUnbondingAmountCallable(delegator common.Address, consensusAddress common.Address) (*big.Int, error) {
	return _Stakepool.Contract.GetUserDelegationUnbondingAmountCallable(&_Stakepool.CallOpts, delegator, consensusAddress)
}

// GetUserTotalDelegationAmount is a free data retrieval call binding the contract method 0xff9ea2f2.
//
// Solidity: function getUserTotalDelegationAmount(address consensusAddress) view returns(uint256)
func (_Stakepool *StakepoolCaller) GetUserTotalDelegationAmount(opts *bind.CallOpts, consensusAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stakepool.contract.Call(opts, &out, "getUserTotalDelegationAmount", consensusAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserTotalDelegationAmount is a free data retrieval call binding the contract method 0xff9ea2f2.
//
// Solidity: function getUserTotalDelegationAmount(address consensusAddress) view returns(uint256)
func (_Stakepool *StakepoolSession) GetUserTotalDelegationAmount(consensusAddress common.Address) (*big.Int, error) {
	return _Stakepool.Contract.GetUserTotalDelegationAmount(&_Stakepool.CallOpts, consensusAddress)
}

// GetUserTotalDelegationAmount is a free data retrieval call binding the contract method 0xff9ea2f2.
//
// Solidity: function getUserTotalDelegationAmount(address consensusAddress) view returns(uint256)
func (_Stakepool *StakepoolCallerSession) GetUserTotalDelegationAmount(consensusAddress common.Address) (*big.Int, error) {
	return _Stakepool.Contract.GetUserTotalDelegationAmount(&_Stakepool.CallOpts, consensusAddress)
}

// GetUserTotalDelegationAmountCallable is a free data retrieval call binding the contract method 0x35044142.
//
// Solidity: function getUserTotalDelegationAmountCallable(address delegator, address consensusAddress) view returns(uint256)
func (_Stakepool *StakepoolCaller) GetUserTotalDelegationAmountCallable(opts *bind.CallOpts, delegator common.Address, consensusAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stakepool.contract.Call(opts, &out, "getUserTotalDelegationAmountCallable", delegator, consensusAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserTotalDelegationAmountCallable is a free data retrieval call binding the contract method 0x35044142.
//
// Solidity: function getUserTotalDelegationAmountCallable(address delegator, address consensusAddress) view returns(uint256)
func (_Stakepool *StakepoolSession) GetUserTotalDelegationAmountCallable(delegator common.Address, consensusAddress common.Address) (*big.Int, error) {
	return _Stakepool.Contract.GetUserTotalDelegationAmountCallable(&_Stakepool.CallOpts, delegator, consensusAddress)
}

// GetUserTotalDelegationAmountCallable is a free data retrieval call binding the contract method 0x35044142.
//
// Solidity: function getUserTotalDelegationAmountCallable(address delegator, address consensusAddress) view returns(uint256)
func (_Stakepool *StakepoolCallerSession) GetUserTotalDelegationAmountCallable(delegator common.Address, consensusAddress common.Address) (*big.Int, error) {
	return _Stakepool.Contract.GetUserTotalDelegationAmountCallable(&_Stakepool.CallOpts, delegator, consensusAddress)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address validatorConsensus) payable returns()
func (_Stakepool *StakepoolTransactor) Delegate(opts *bind.TransactOpts, validatorConsensus common.Address) (*types.Transaction, error) {
	return _Stakepool.contract.Transact(opts, "delegate", validatorConsensus)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address validatorConsensus) payable returns()
func (_Stakepool *StakepoolSession) Delegate(validatorConsensus common.Address) (*types.Transaction, error) {
	return _Stakepool.Contract.Delegate(&_Stakepool.TransactOpts, validatorConsensus)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address validatorConsensus) payable returns()
func (_Stakepool *StakepoolTransactorSession) Delegate(validatorConsensus common.Address) (*types.Transaction, error) {
	return _Stakepool.Contract.Delegate(&_Stakepool.TransactOpts, validatorConsensus)
}

// GetUnbondingQueueAmount is a paid mutator transaction binding the contract method 0x16f23fdc.
//
// Solidity: function getUnbondingQueueAmount(address delegator) returns(uint256[])
func (_Stakepool *StakepoolTransactor) GetUnbondingQueueAmount(opts *bind.TransactOpts, delegator common.Address) (*types.Transaction, error) {
	return _Stakepool.contract.Transact(opts, "getUnbondingQueueAmount", delegator)
}

// GetUnbondingQueueAmount is a paid mutator transaction binding the contract method 0x16f23fdc.
//
// Solidity: function getUnbondingQueueAmount(address delegator) returns(uint256[])
func (_Stakepool *StakepoolSession) GetUnbondingQueueAmount(delegator common.Address) (*types.Transaction, error) {
	return _Stakepool.Contract.GetUnbondingQueueAmount(&_Stakepool.TransactOpts, delegator)
}

// GetUnbondingQueueAmount is a paid mutator transaction binding the contract method 0x16f23fdc.
//
// Solidity: function getUnbondingQueueAmount(address delegator) returns(uint256[])
func (_Stakepool *StakepoolTransactorSession) GetUnbondingQueueAmount(delegator common.Address) (*types.Transaction, error) {
	return _Stakepool.Contract.GetUnbondingQueueAmount(&_Stakepool.TransactOpts, delegator)
}

// GetUnbondingQueueTime is a paid mutator transaction binding the contract method 0xda334f4d.
//
// Solidity: function getUnbondingQueueTime(address delegator) returns(uint256[])
func (_Stakepool *StakepoolTransactor) GetUnbondingQueueTime(opts *bind.TransactOpts, delegator common.Address) (*types.Transaction, error) {
	return _Stakepool.contract.Transact(opts, "getUnbondingQueueTime", delegator)
}

// GetUnbondingQueueTime is a paid mutator transaction binding the contract method 0xda334f4d.
//
// Solidity: function getUnbondingQueueTime(address delegator) returns(uint256[])
func (_Stakepool *StakepoolSession) GetUnbondingQueueTime(delegator common.Address) (*types.Transaction, error) {
	return _Stakepool.Contract.GetUnbondingQueueTime(&_Stakepool.TransactOpts, delegator)
}

// GetUnbondingQueueTime is a paid mutator transaction binding the contract method 0xda334f4d.
//
// Solidity: function getUnbondingQueueTime(address delegator) returns(uint256[])
func (_Stakepool *StakepoolTransactorSession) GetUnbondingQueueTime(delegator common.Address) (*types.Transaction, error) {
	return _Stakepool.Contract.GetUnbondingQueueTime(&_Stakepool.TransactOpts, delegator)
}

// GetUnbondingQueueValidator is a paid mutator transaction binding the contract method 0x2f2d2cd9.
//
// Solidity: function getUnbondingQueueValidator(address delegator) returns(address[])
func (_Stakepool *StakepoolTransactor) GetUnbondingQueueValidator(opts *bind.TransactOpts, delegator common.Address) (*types.Transaction, error) {
	return _Stakepool.contract.Transact(opts, "getUnbondingQueueValidator", delegator)
}

// GetUnbondingQueueValidator is a paid mutator transaction binding the contract method 0x2f2d2cd9.
//
// Solidity: function getUnbondingQueueValidator(address delegator) returns(address[])
func (_Stakepool *StakepoolSession) GetUnbondingQueueValidator(delegator common.Address) (*types.Transaction, error) {
	return _Stakepool.Contract.GetUnbondingQueueValidator(&_Stakepool.TransactOpts, delegator)
}

// GetUnbondingQueueValidator is a paid mutator transaction binding the contract method 0x2f2d2cd9.
//
// Solidity: function getUnbondingQueueValidator(address delegator) returns(address[])
func (_Stakepool *StakepoolTransactorSession) GetUnbondingQueueValidator(delegator common.Address) (*types.Transaction, error) {
	return _Stakepool.Contract.GetUnbondingQueueValidator(&_Stakepool.TransactOpts, delegator)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address validatorPoolAddress) returns()
func (_Stakepool *StakepoolTransactor) Init(opts *bind.TransactOpts, validatorPoolAddress common.Address) (*types.Transaction, error) {
	return _Stakepool.contract.Transact(opts, "init", validatorPoolAddress)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address validatorPoolAddress) returns()
func (_Stakepool *StakepoolSession) Init(validatorPoolAddress common.Address) (*types.Transaction, error) {
	return _Stakepool.Contract.Init(&_Stakepool.TransactOpts, validatorPoolAddress)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address validatorPoolAddress) returns()
func (_Stakepool *StakepoolTransactorSession) Init(validatorPoolAddress common.Address) (*types.Transaction, error) {
	return _Stakepool.Contract.Init(&_Stakepool.TransactOpts, validatorPoolAddress)
}

// RemoveDelegation is a paid mutator transaction binding the contract method 0x3dff07d2.
//
// Solidity: function removeDelegation(address delegator, address validator) returns()
func (_Stakepool *StakepoolTransactor) RemoveDelegation(opts *bind.TransactOpts, delegator common.Address, validator common.Address) (*types.Transaction, error) {
	return _Stakepool.contract.Transact(opts, "removeDelegation", delegator, validator)
}

// RemoveDelegation is a paid mutator transaction binding the contract method 0x3dff07d2.
//
// Solidity: function removeDelegation(address delegator, address validator) returns()
func (_Stakepool *StakepoolSession) RemoveDelegation(delegator common.Address, validator common.Address) (*types.Transaction, error) {
	return _Stakepool.Contract.RemoveDelegation(&_Stakepool.TransactOpts, delegator, validator)
}

// RemoveDelegation is a paid mutator transaction binding the contract method 0x3dff07d2.
//
// Solidity: function removeDelegation(address delegator, address validator) returns()
func (_Stakepool *StakepoolTransactorSession) RemoveDelegation(delegator common.Address, validator common.Address) (*types.Transaction, error) {
	return _Stakepool.Contract.RemoveDelegation(&_Stakepool.TransactOpts, delegator, validator)
}

// RemoveUnbondingUserFromUnbondingQueue is a paid mutator transaction binding the contract method 0x9dbcd939.
//
// Solidity: function removeUnbondingUserFromUnbondingQueue() returns()
func (_Stakepool *StakepoolTransactor) RemoveUnbondingUserFromUnbondingQueue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakepool.contract.Transact(opts, "removeUnbondingUserFromUnbondingQueue")
}

// RemoveUnbondingUserFromUnbondingQueue is a paid mutator transaction binding the contract method 0x9dbcd939.
//
// Solidity: function removeUnbondingUserFromUnbondingQueue() returns()
func (_Stakepool *StakepoolSession) RemoveUnbondingUserFromUnbondingQueue() (*types.Transaction, error) {
	return _Stakepool.Contract.RemoveUnbondingUserFromUnbondingQueue(&_Stakepool.TransactOpts)
}

// RemoveUnbondingUserFromUnbondingQueue is a paid mutator transaction binding the contract method 0x9dbcd939.
//
// Solidity: function removeUnbondingUserFromUnbondingQueue() returns()
func (_Stakepool *StakepoolTransactorSession) RemoveUnbondingUserFromUnbondingQueue() (*types.Transaction, error) {
	return _Stakepool.Contract.RemoveUnbondingUserFromUnbondingQueue(&_Stakepool.TransactOpts)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address validatorConsensus, uint256 amount) returns()
func (_Stakepool *StakepoolTransactor) Undelegate(opts *bind.TransactOpts, validatorConsensus common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stakepool.contract.Transact(opts, "undelegate", validatorConsensus, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address validatorConsensus, uint256 amount) returns()
func (_Stakepool *StakepoolSession) Undelegate(validatorConsensus common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stakepool.Contract.Undelegate(&_Stakepool.TransactOpts, validatorConsensus, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address validatorConsensus, uint256 amount) returns()
func (_Stakepool *StakepoolTransactorSession) Undelegate(validatorConsensus common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stakepool.Contract.Undelegate(&_Stakepool.TransactOpts, validatorConsensus, amount)
}

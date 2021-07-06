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

// ValidatorsetBin is the compiled bytecode used for deploying new contracts.
var ValidatorsetBin = "0x6080604052600a60015534801561001557600080fd5b5061211f806100256000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c8063a78abc1611610066578063a78abc1614610111578063ad3c9da61461012f578063b7ab4db51461015f578063e589b61e1461017d578063f09a40161461019957610093565b80633197cbb61461009857806367d07b2f146100b65780636969a25c146100c057806374de1669146100f3575b600080fd5b6100a06101b5565b6040516100ad9190611e2b565b60405180910390f35b6100be6101bb565b005b6100da60048036038101906100d59190611b73565b610ef2565b6040516100ea9493929190611d69565b60405180910390f35b6100fb610f6c565b6040516101089190611e2b565b60405180910390f35b610119610f72565b6040516101269190611dd0565b60405180910390f35b61014960048036038101906101449190611aab565b610f83565b6040516101569190611e2b565b60405180910390f35b610167610f9b565b6040516101749190611dae565b60405180910390f35b61019760048036038101906101929190611aab565b611149565b005b6101b360048036038101906101ae9190611ad4565b611289565b005b60085481565b6000600290506002600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630c42f6316040518163ffffffff1660e01b815260040160206040518083038186803b15801561022b57600080fd5b505afa15801561023f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102639190611b9c565b101561030c57600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630c42f6316040518163ffffffff1660e01b815260040160206040518083038186803b1580156102d157600080fd5b505afa1580156102e5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103099190611b9c565b90505b60005b60028054905081101561056157600060028281548110610358577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600052602060002090600302016040518060800160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820160009054906101000a900460ff16600281111561041c577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6002811115610454577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b81526020016002820160019054906101000a900460ff1615151515815250509050600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166341eeacdb82600001516040518263ffffffff1660e01b81526004016104d49190611d4e565b600060405180830381600087803b1580156104ee57600080fd5b505af1158015610502573d6000803e3d6000fd5b5050505060036000826000015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000905550808061055990611f87565b91505061030f565b5060026000610570919061195d565b6006600061057e9190611981565b806004819055506000600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630c42f6316040518163ffffffff1660e01b815260040160206040518083038186803b1580156105ef57600080fd5b505afa158015610603573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106279190611b9c565b905060005b818110156107ff576000600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166335aa2e44836040518263ffffffff1660e01b81526004016106919190611e2b565b60806040518083038186803b1580156106a957600080fd5b505afa1580156106bd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106e19190611b10565b50505090506000600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663322923b9836040518263ffffffff1660e01b81526004016107439190611d4e565b60206040518083038186803b15801561075b57600080fd5b505afa15801561076f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107939190611b9c565b9050600681908060018154018082558091505060019003906000526020600020016000909190919091505560006005600085815260200190815260200160002060006101000a81548160ff021916908315150217905550505080806107f790611f87565b91505061062c565b5060005b600454811015610e635760008060005b84811015610a6b57600080600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166335aa2e44846040518263ffffffff1660e01b81526004016108799190611e2b565b60806040518083038186803b15801561089157600080fd5b505afa1580156108a5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108c99190611b10565b9350505091506000600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166350ef3309846040518263ffffffff1660e01b815260040161092c9190611d4e565b60206040518083038186803b15801561094457600080fd5b505afa158015610958573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061097c9190611b9c565b905085600685815481106109b9577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600052602060002001541180156109ef57506005600085815260200190815260200160002060009054906101000a900460ff16155b80156109f9575081155b8015610a055750600081145b15610a555760068481548110610a44577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020015495508394505b5050508080610a6390611f87565b915050610813565b5060016005600083815260200190815260200160002060006101000a81548160ff021916908315150217905550600080600080600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166335aa2e44866040518263ffffffff1660e01b8152600401610af99190611e2b565b60806040518083038186803b158015610b1157600080fd5b505afa158015610b25573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b499190611b10565b9350935093509350600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663360d80ec856040518263ffffffff1660e01b8152600401610bac9190611d4e565b600060405180830381600087803b158015610bc657600080fd5b505af1158015610bda573d6000803e3d6000fd5b50505050600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166335aa2e44866040518263ffffffff1660e01b8152600401610c399190611e2b565b60806040518083038186803b158015610c5157600080fd5b505afa158015610c65573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c899190611b10565b909192509091505080925050600260405180608001604052808673ffffffffffffffffffffffffffffffffffffffff168152602001858152602001846002811115610cfd577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8152602001831515815250908060018154018082558091505060019003906000526020600020906003020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548160ff02191690836002811115610dd3577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b021790555060608201518160020160016101000a81548160ff0219169083151502179055505050600187610e079190611e90565b600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050505050508080610e5b90611f87565b915050610803565b50600c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638f73c5ae6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015610ece57600080fd5b505af1158015610ee2573d6000803e3d6000fd5b50505050610eee61141c565b5050565b60028181548110610f0257600080fd5b90600052602060002090600302016000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900460ff16908060020160019054906101000a900460ff16905084565b60045481565b60008054906101000a900460ff1681565b60036020528060005260406000206000915090505481565b606060008054906101000a900460ff16610fea576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fe190611deb565b60405180910390fd5b6002805480602002602001604051908101604052809291908181526020016000905b8282101561114057838290600052602060002090600302016040518060800160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820160009054906101000a900460ff1660028111156110da577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6002811115611112577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b81526020016002820160019054906101000a900460ff1615151515815250508152602001906001019061100c565b50505050905090565b60008054906101000a900460ff16611196576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161118d90611deb565b60405180910390fd5b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146111f057600080fd5b6111f981611439565b600c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663348749e0826040518263ffffffff1660e01b81526004016112549190611d4e565b600060405180830381600087803b15801561126e57600080fd5b505af1158015611282573d6000803e3d6000fd5b5050505050565b60008054906101000a900460ff16156112d7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112ce90611e0b565b60405180910390fd5b4260088190555081600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600b60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600c60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506113ea6101bb565b600154426113f89190611e90565b60088190555060016000806101000a81548160ff0219169083151502179055505050565b600154600860008282546114309190611e90565b92505081905550565b60008054906101000a900460ff16611486576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161147d90611deb565b60405180910390fd5b60006001600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546114d49190611ee6565b905060008190505b60016004546114eb9190611ee6565b8110156118535760026001826115019190611e90565b81548110611538577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020906003020160028281548110611580577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600052602060002090600302016000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600182015481600101556002820160009054906101000a900460ff168160020160006101000a81548160ff0219169083600281111561165d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b02179055506002820160019054906101000a900460ff168160020160016101000a81548160ff0219169083151502179055509050506000600282815481106116ce577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600052602060002090600302016040518060800160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820160009054906101000a900460ff166002811115611792577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60028111156117ca577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b81526020016002820160019054906101000a900460ff16151515158152505090506001826117f89190611e90565b60036000836000015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050808061184b90611f87565b9150506114dc565b50600280548061188c577f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6001900381819060005260206000209060030201600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905560018201600090556002820160006101000a81549060ff02191690556002820160016101000a81549060ff021916905550509055600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600090556001600460008282546119529190611ee6565b925050819055505050565b508054600082556003029060005260206000209081019061197e91906119a2565b50565b508054600082559060005260206000209081019061199f9190611a10565b50565b5b80821115611a0c57600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905560018201600090556002820160006101000a81549060ff02191690556002820160016101000a81549060ff0219169055506003016119a3565b5090565b5b80821115611a29576000816000905550600101611a11565b5090565b600081359050611a3c81612094565b92915050565b600081519050611a5181612094565b92915050565b600081519050611a66816120ab565b92915050565b600081519050611a7b816120c2565b92915050565b600081359050611a90816120d2565b92915050565b600081519050611aa5816120d2565b92915050565b600060208284031215611abd57600080fd5b6000611acb84828501611a2d565b91505092915050565b60008060408385031215611ae757600080fd5b6000611af585828601611a2d565b9250506020611b0685828601611a2d565b9150509250929050565b60008060008060808587031215611b2657600080fd5b6000611b3487828801611a42565b9450506020611b4587828801611a96565b9350506040611b5687828801611a6c565b9250506060611b6787828801611a57565b91505092959194509250565b600060208284031215611b8557600080fd5b6000611b9384828501611a81565b91505092915050565b600060208284031215611bae57600080fd5b6000611bbc84828501611a96565b91505092915050565b6000611bd18383611cdb565b60808301905092915050565b611be681611f1a565b82525050565b611bf581611f1a565b82525050565b6000611c0682611e56565b611c108185611e6e565b9350611c1b83611e46565b8060005b83811015611c4c578151611c338882611bc5565b9750611c3e83611e61565b925050600181019050611c1f565b5085935050505092915050565b611c6281611f2c565b82525050565b611c7181611f2c565b82525050565b611c8081611f75565b82525050565b611c8f81611f75565b82525050565b6000611ca2601983611e7f565b9150611cad8261202e565b602082019050919050565b6000611cc5601983611e7f565b9150611cd082612057565b602082019050919050565b608082016000820151611cf16000850182611bdd565b506020820151611d046020850182611d30565b506040820151611d176040850182611c77565b506060820151611d2a6060850182611c59565b50505050565b611d3981611f6b565b82525050565b611d4881611f6b565b82525050565b6000602082019050611d636000830184611bec565b92915050565b6000608082019050611d7e6000830187611bec565b611d8b6020830186611d3f565b611d986040830185611c86565b611da56060830184611c68565b95945050505050565b60006020820190508181036000830152611dc88184611bfb565b905092915050565b6000602082019050611de56000830184611c68565b92915050565b60006020820190508181036000830152611e0481611c95565b9050919050565b60006020820190508181036000830152611e2481611cb8565b9050919050565b6000602082019050611e406000830184611d3f565b92915050565b6000819050602082019050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b6000611e9b82611f6b565b9150611ea683611f6b565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611edb57611eda611fd0565b5b828201905092915050565b6000611ef182611f6b565b9150611efc83611f6b565b925082821015611f0f57611f0e611fd0565b5b828203905092915050565b6000611f2582611f4b565b9050919050565b60008115159050919050565b6000819050611f4682612080565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000611f8082611f38565b9050919050565b6000611f9282611f6b565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611fc557611fc4611fd0565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f74686520636f6e7472616374206e6f7420696e69742079657400000000000000600082015250565b7f74686520636f6e747261637420616c726561647920696e697400000000000000600082015250565b6003811061209157612090611fff565b5b50565b61209d81611f1a565b81146120a857600080fd5b50565b6120b481611f2c565b81146120bf57600080fd5b50565b600381106120cf57600080fd5b50565b6120db81611f6b565b81146120e657600080fd5b5056fea264697066735822122051fb096b396b37c8993e94e0a8721152e88a95d78736a4d9e349ad32311302f764736f6c63430008040033"

// DeployValidatorset deploys a new Ethereum contract, binding an instance of Validatorset to it.
func DeployValidatorset(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Validatorset, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorsetABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ValidatorsetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Validatorset{ValidatorsetCaller: ValidatorsetCaller{contract: contract}, ValidatorsetTransactor: ValidatorsetTransactor{contract: contract}, ValidatorsetFilterer: ValidatorsetFilterer{contract: contract}}, nil
}

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

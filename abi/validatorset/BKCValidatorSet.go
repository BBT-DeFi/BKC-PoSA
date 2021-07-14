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
var ValidatorsetBin = "0x6080604052600a60015534801561001557600080fd5b50611f36806100256000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c8063a78abc1611610066578063a78abc1614610111578063ad3c9da61461012f578063b7ab4db51461015f578063e589b61e1461017d578063f09a40161461019957610093565b80633197cbb61461009857806367d07b2f146100b65780636969a25c146100c057806374de1669146100f3575b600080fd5b6100a06101b5565b6040516100ad9190611bdf565b60405180910390f35b6100be6101bb565b005b6100da60048036038101906100d5919061191f565b610de8565b6040516100ea9493929190611b1d565b60405180910390f35b6100fb610e62565b6040516101089190611bdf565b60405180910390f35b610119610e68565b6040516101269190611b84565b60405180910390f35b6101496004803603810190610144919061184b565b610e79565b6040516101569190611bdf565b60405180910390f35b610167610e91565b6040516101749190611b62565b60405180910390f35b6101976004803603810190610192919061184b565b610ff3565b005b6101b360048036038101906101ae9190611878565b611133565b005b60085481565b6000600290506002600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630c42f6316040518163ffffffff1660e01b815260040160206040518083038186803b15801561022b57600080fd5b505afa15801561023f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610263919061194c565b101561030c57600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630c42f6316040518163ffffffff1660e01b815260040160206040518083038186803b1580156102d157600080fd5b505afa1580156102e5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610309919061194c565b90505b60005b6002805490508110156104ef5760006002828154811061033257610331611e11565b5b90600052602060002090600302016040518060800160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820160009054906101000a900460ff1660028111156103d0576103cf611db3565b5b60028111156103e2576103e1611db3565b5b81526020016002820160019054906101000a900460ff1615151515815250509050600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166341eeacdb82600001516040518263ffffffff1660e01b81526004016104629190611b02565b600060405180830381600087803b15801561047c57600080fd5b505af1158015610490573d6000803e3d6000fd5b5050505060036000826000015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600090555080806104e790611d3b565b91505061030f565b50600260006104fe91906116fd565b6006600061050c9190611721565b806004819055506000600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630c42f6316040518163ffffffff1660e01b815260040160206040518083038186803b15801561057d57600080fd5b505afa158015610591573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105b5919061194c565b905060005b8181101561078d576000600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166335aa2e44836040518263ffffffff1660e01b815260040161061f9190611bdf565b60806040518083038186803b15801561063757600080fd5b505afa15801561064b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061066f91906118b8565b50505090506000600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663322923b9836040518263ffffffff1660e01b81526004016106d19190611b02565b60206040518083038186803b1580156106e957600080fd5b505afa1580156106fd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610721919061194c565b9050600681908060018154018082558091505060019003906000526020600020016000909190919091505560006005600085815260200190815260200160002060006101000a81548160ff0219169083151502179055505050808061078590611d3b565b9150506105ba565b5060005b600454811015610d595760008060005b848110156109ad57600080600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166335aa2e44846040518263ffffffff1660e01b81526004016108079190611bdf565b60806040518083038186803b15801561081f57600080fd5b505afa158015610833573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061085791906118b8565b9350505091506000600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166350ef3309846040518263ffffffff1660e01b81526004016108ba9190611b02565b60206040518083038186803b1580156108d257600080fd5b505afa1580156108e6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061090a919061194c565b9050856006858154811061092157610920611e11565b5b906000526020600020015411801561095757506005600085815260200190815260200160002060009054906101000a900460ff16155b8015610961575081155b801561096d5750600081145b15610997576006848154811061098657610985611e11565b5b906000526020600020015495508394505b50505080806109a590611d3b565b9150506107a1565b5060016005600083815260200190815260200160002060006101000a81548160ff021916908315150217905550600080600080600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166335aa2e44866040518263ffffffff1660e01b8152600401610a3b9190611bdf565b60806040518083038186803b158015610a5357600080fd5b505afa158015610a67573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a8b91906118b8565b9350935093509350600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663360d80ec856040518263ffffffff1660e01b8152600401610aee9190611b02565b600060405180830381600087803b158015610b0857600080fd5b505af1158015610b1c573d6000803e3d6000fd5b50505050600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166335aa2e44866040518263ffffffff1660e01b8152600401610b7b9190611bdf565b60806040518083038186803b158015610b9357600080fd5b505afa158015610ba7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bcb91906118b8565b909192509091505080925050600260405180608001604052808673ffffffffffffffffffffffffffffffffffffffff168152602001858152602001846002811115610c1957610c18611db3565b5b8152602001831515815250908060018154018082558091505060019003906000526020600020906003020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020160006101000a81548160ff02191690836002811115610cc957610cc8611db3565b5b021790555060608201518160020160016101000a81548160ff0219169083151502179055505050600187610cfd9190611c44565b600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050505050508080610d5190611d3b565b915050610791565b50600c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638f73c5ae6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015610dc457600080fd5b505af1158015610dd8573d6000803e3d6000fd5b50505050610de46112c6565b5050565b60028181548110610df857600080fd5b90600052602060002090600302016000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020160009054906101000a900460ff16908060020160019054906101000a900460ff16905084565b60045481565b60008054906101000a900460ff1681565b60036020528060005260406000206000915090505481565b606060008054906101000a900460ff16610ee0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ed790611b9f565b60405180910390fd5b6002805480602002602001604051908101604052809291908181526020016000905b82821015610fea57838290600052602060002090600302016040518060800160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820160009054906101000a900460ff166002811115610faa57610fa9611db3565b5b6002811115610fbc57610fbb611db3565b5b81526020016002820160019054906101000a900460ff16151515158152505081526020019060010190610f02565b50505050905090565b60008054906101000a900460ff16611040576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161103790611b9f565b60405180910390fd5b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461109a57600080fd5b6110a3816112e3565b600c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663348749e0826040518263ffffffff1660e01b81526004016110fe9190611b02565b600060405180830381600087803b15801561111857600080fd5b505af115801561112c573d6000803e3d6000fd5b5050505050565b60008054906101000a900460ff1615611181576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161117890611bbf565b60405180910390fd5b4260088190555081600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600b60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600c60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506112946101bb565b600154426112a29190611c44565b60088190555060016000806101000a81548160ff0219169083151502179055505050565b600154600860008282546112da9190611c44565b92505081905550565b60008054906101000a900460ff16611330576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161132790611b9f565b60405180910390fd5b60006001600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461137e9190611c9a565b905060008190505b60016004546113959190611c9a565b8110156116195760026001826113ab9190611c44565b815481106113bc576113bb611e11565b5b9060005260206000209060030201600282815481106113de576113dd611e11565b5b90600052602060002090600302016000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600182015481600101556002820160009054906101000a900460ff168160020160006101000a81548160ff0219169083600281111561149557611494611db3565b5b02179055506002820160019054906101000a900460ff168160020160016101000a81548160ff0219169083151502179055509050506000600282815481106114e0576114df611e11565b5b90600052602060002090600302016040518060800160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182015481526020016002820160009054906101000a900460ff16600281111561157e5761157d611db3565b5b60028111156115905761158f611db3565b5b81526020016002820160019054906101000a900460ff16151515158152505090506001826115be9190611c44565b60036000836000015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050808061161190611d3b565b915050611386565b50600280548061162c5761162b611de2565b5b6001900381819060005260206000209060030201600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905560018201600090556002820160006101000a81549060ff02191690556002820160016101000a81549060ff021916905550509055600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600090556001600460008282546116f29190611c9a565b925050819055505050565b508054600082556003029060005260206000209081019061171e9190611742565b50565b508054600082559060005260206000209081019061173f91906117b0565b50565b5b808211156117ac57600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905560018201600090556002820160006101000a81549060ff02191690556002820160016101000a81549060ff021916905550600301611743565b5090565b5b808211156117c95760008160009055506001016117b1565b5090565b6000813590506117dc81611eab565b92915050565b6000815190506117f181611eab565b92915050565b60008151905061180681611ec2565b92915050565b60008151905061181b81611ed9565b92915050565b60008135905061183081611ee9565b92915050565b60008151905061184581611ee9565b92915050565b60006020828403121561186157611860611e40565b5b600061186f848285016117cd565b91505092915050565b6000806040838503121561188f5761188e611e40565b5b600061189d858286016117cd565b92505060206118ae858286016117cd565b9150509250929050565b600080600080608085870312156118d2576118d1611e40565b5b60006118e0878288016117e2565b94505060206118f187828801611836565b93505060406119028782880161180c565b9250506060611913878288016117f7565b91505092959194509250565b60006020828403121561193557611934611e40565b5b600061194384828501611821565b91505092915050565b60006020828403121561196257611961611e40565b5b600061197084828501611836565b91505092915050565b60006119858383611a8f565b60808301905092915050565b61199a81611cce565b82525050565b6119a981611cce565b82525050565b60006119ba82611c0a565b6119c48185611c22565b93506119cf83611bfa565b8060005b83811015611a005781516119e78882611979565b97506119f283611c15565b9250506001810190506119d3565b5085935050505092915050565b611a1681611ce0565b82525050565b611a2581611ce0565b82525050565b611a3481611d29565b82525050565b611a4381611d29565b82525050565b6000611a56601983611c33565b9150611a6182611e45565b602082019050919050565b6000611a79601983611c33565b9150611a8482611e6e565b602082019050919050565b608082016000820151611aa56000850182611991565b506020820151611ab86020850182611ae4565b506040820151611acb6040850182611a2b565b506060820151611ade6060850182611a0d565b50505050565b611aed81611d1f565b82525050565b611afc81611d1f565b82525050565b6000602082019050611b1760008301846119a0565b92915050565b6000608082019050611b3260008301876119a0565b611b3f6020830186611af3565b611b4c6040830185611a3a565b611b596060830184611a1c565b95945050505050565b60006020820190508181036000830152611b7c81846119af565b905092915050565b6000602082019050611b996000830184611a1c565b92915050565b60006020820190508181036000830152611bb881611a49565b9050919050565b60006020820190508181036000830152611bd881611a6c565b9050919050565b6000602082019050611bf46000830184611af3565b92915050565b6000819050602082019050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b6000611c4f82611d1f565b9150611c5a83611d1f565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611c8f57611c8e611d84565b5b828201905092915050565b6000611ca582611d1f565b9150611cb083611d1f565b925082821015611cc357611cc2611d84565b5b828203905092915050565b6000611cd982611cff565b9050919050565b60008115159050919050565b6000819050611cfa82611e97565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000611d3482611cec565b9050919050565b6000611d4682611d1f565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611d7957611d78611d84565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600080fd5b7f74686520636f6e7472616374206e6f7420696e69742079657400000000000000600082015250565b7f74686520636f6e747261637420616c726561647920696e697400000000000000600082015250565b60038110611ea857611ea7611db3565b5b50565b611eb481611cce565b8114611ebf57600080fd5b50565b611ecb81611ce0565b8114611ed657600080fd5b50565b60038110611ee657600080fd5b50565b611ef281611d1f565b8114611efd57600080fd5b5056fea2646970667358221220f631151ba60b5d9b6995c8d423c9ca5154d40a0ec7f6eaab4fe287b2eaf55b6064736f6c63430008060033"

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

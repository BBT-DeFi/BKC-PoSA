// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package kbd

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

// KbdABI is the input ABI used to generate the binding from.
const KbdABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fund\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sell\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"this_balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// KbdBin is the compiled bytecode used for deploying new contracts.
var KbdBin = "0x60806040526113886006557360bb6c1b4a0f5b1ea820be6c762384982b8a658c600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055503480156200006c57600080fd5b506040518060400160405280600581526020017f53616b6f6e0000000000000000000000000000000000000000000000000000008152506040518060400160405280600381526020017f57494e00000000000000000000000000000000000000000000000000000000008152508160039080519060200190620000f192919062000322565b5080600490805190602001906200010a92919062000322565b5050506200014f33620001226200019660201b60201c565b60ff16600a62000133919062000512565b620f42406200014391906200064f565b6200019f60201b60201c565b33600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555062000784565b60006012905090565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141562000212576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000209906200040a565b60405180910390fd5b62000226600083836200031860201b60201c565b80600260008282546200023a91906200045a565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546200029191906200045a565b925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051620002f891906200042c565b60405180910390a362000314600083836200031d60201b60201c565b5050565b505050565b505050565b8280546200033090620006ba565b90600052602060002090601f016020900481019282620003545760008555620003a0565b82601f106200036f57805160ff1916838001178555620003a0565b82800160010185558215620003a0579182015b828111156200039f57825182559160200191906001019062000382565b5b509050620003af9190620003b3565b5090565b5b80821115620003ce576000816000905550600101620003b4565b5090565b6000620003e1601f8362000449565b9150620003ee826200075b565b602082019050919050565b6200040481620006b0565b82525050565b600060208201905081810360008301526200042581620003d2565b9050919050565b6000602082019050620004436000830184620003f9565b92915050565b600082825260208201905092915050565b60006200046782620006b0565b91506200047483620006b0565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115620004ac57620004ab620006f0565b5b828201905092915050565b6000808291508390505b60018511156200050957808604811115620004e157620004e0620006f0565b5b6001851615620004f15780820291505b808102905062000501856200074e565b9450620004c1565b94509492505050565b60006200051f82620006b0565b91506200052c83620006b0565b92506200055b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff848462000563565b905092915050565b60008262000575576001905062000648565b8162000585576000905062000648565b81600181146200059e5760028114620005a957620005df565b600191505062000648565b60ff841115620005be57620005bd620006f0565b5b8360020a915084821115620005d857620005d7620006f0565b5b5062000648565b5060208310610133831016604e8410600b8410161715620006195782820a905083811115620006135762000612620006f0565b5b62000648565b620006288484846001620004b7565b92509050818404811115620006425762000641620006f0565b5b81810290505b9392505050565b60006200065c82620006b0565b91506200066983620006b0565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615620006a557620006a4620006f0565b5b828202905092915050565b6000819050919050565b60006002820490506001821680620006d357607f821691505b60208210811415620006ea57620006e96200071f565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60008160011c9050919050565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b6122ff80620007946000396000f3fe6080604052600436106101095760003560e01c806395d89b4111610095578063a9059cbb11610064578063a9059cbb14610372578063b60d4288146103af578063dd62ed3e146103b9578063e3faecc2146103f6578063e4849b321461042157610109565b806395d89b41146102d75780639dc29fac14610302578063a457c2d71461032b578063a6f2ae3a1461036857610109565b8063313ce567116100dc578063313ce567146101de57806339509351146102095780633ba0b9a91461024657806340c10f191461027157806370a082311461029a57610109565b806306fdde031461010e578063095ea7b31461013957806318160ddd1461017657806323b872dd146101a1575b600080fd5b34801561011a57600080fd5b5061012361044a565b60405161013091906118e3565b60405180910390f35b34801561014557600080fd5b50610160600480360381019061015b919061160b565b6104dc565b60405161016d91906118c8565b60405180910390f35b34801561018257600080fd5b5061018b6104fa565b6040516101989190611ac5565b60405180910390f35b3480156101ad57600080fd5b506101c860048036038101906101c391906115b8565b610504565b6040516101d591906118c8565b60405180910390f35b3480156101ea57600080fd5b506101f36105fc565b6040516102009190611ae0565b60405180910390f35b34801561021557600080fd5b50610230600480360381019061022b919061160b565b610605565b60405161023d91906118c8565b60405180910390f35b34801561025257600080fd5b5061025b6106b1565b6040516102689190611ac5565b60405180910390f35b34801561027d57600080fd5b506102986004803603810190610293919061160b565b6106b7565b005b3480156102a657600080fd5b506102c160048036038101906102bc919061154b565b610776565b6040516102ce9190611ac5565b60405180910390f35b3480156102e357600080fd5b506102ec6107be565b6040516102f991906118e3565b60405180910390f35b34801561030e57600080fd5b506103296004803603810190610324919061160b565b610850565b005b34801561033757600080fd5b50610352600480360381019061034d919061160b565b61090f565b60405161035f91906118c8565b60405180910390f35b6103706109fa565b005b34801561037e57600080fd5b506103996004803603810190610394919061160b565b610afe565b6040516103a691906118c8565b60405180910390f35b6103b7610b1c565b005b3480156103c557600080fd5b506103e060048036038101906103db9190611578565b610b61565b6040516103ed9190611ac5565b60405180910390f35b34801561040257600080fd5b5061040b610be8565b6040516104189190611ac5565b60405180910390f35b34801561042d57600080fd5b506104486004803603810190610443919061164b565b610c80565b005b60606003805461045990611e25565b80601f016020809104026020016040519081016040528092919081815260200182805461048590611e25565b80156104d25780601f106104a7576101008083540402835291602001916104d2565b820191906000526020600020905b8154815290600101906020018083116104b557829003601f168201915b5050505050905090565b60006104f06104e9610d68565b8484610d70565b6001905092915050565b6000600254905090565b6000610511848484610f3b565b6000600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600061055c610d68565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050828110156105dc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105d3906119e5565b60405180910390fd5b6105f0856105e8610d68565b858403610d70565b60019150509392505050565b60006012905090565b60006106a7610612610d68565b848460016000610620610d68565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546106a29190611b17565b610d70565b6001905092915050565b60065481565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610747576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161073e90611965565b60405180910390fd5b610772826107536105fc565b60ff16600a6107629190611bf1565b8361076d9190611d0f565b6111bc565b5050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6060600480546107cd90611e25565b80601f01602080910402602001604051908101604052809291908181526020018280546107f990611e25565b80156108465780601f1061081b57610100808354040283529160200191610846565b820191906000526020600020905b81548152906001019060200180831161082957829003601f168201915b5050505050905090565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108e0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108d790611965565b60405180910390fd5b61090b826108ec6105fc565b60ff16600a6108fb9190611bf1565b836109069190611d0f565b61131c565b5050565b6000806001600061091e610d68565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050828110156109db576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109d290611a85565b60405180910390fd5b6109ef6109e6610d68565b85858403610d70565b600191505092915050565b60003411610a3d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a34906119c5565b60405180910390fd5b600060065434610a4d9190611d0f565b905080610a7b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16610776565b1115610abc57610aad600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168261131c565b610ab733826111bc565b610af3565b610ac633826111bc565b610af2600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16826111bc565b5b610afb6114f3565b50565b6000610b12610b0b610d68565b8484610f3b565b6001905092915050565b60003411610b5f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b5690611a45565b60405180910390fd5b565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610c7a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c7190611965565b60405180910390fd5b47905090565b80610c8a33610776565b1015610ccb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cc290611985565b60405180910390fd5b600060065482610cdb9190611b6d565b9050600033905081471015610cef57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166108fc839081150290604051600060405180830381858888f19350505050158015610d35573d6000803e3d6000fd5b50610d62600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1684610afe565b50505050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610de0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dd790611a65565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610e50576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e4790611945565b60405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92583604051610f2e9190611ac5565b60405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610fab576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fa290611a25565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561101b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161101290611905565b60405180910390fd5b611026838383611517565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050818110156110ac576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110a3906119a5565b60405180910390fd5b8181036000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461113f9190611b17565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516111a39190611ac5565b60405180910390a36111b684848461151c565b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561122c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161122390611aa5565b60405180910390fd5b61123860008383611517565b806002600082825461124a9190611b17565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461129f9190611b17565b925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516113049190611ac5565b60405180910390a36113186000838361151c565b5050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561138c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161138390611a05565b60405180910390fd5b61139882600083611517565b60008060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490508181101561141e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161141590611925565b60405180910390fd5b8181036000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600260008282546114759190611d69565b92505081905550600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516114da9190611ac5565b60405180910390a36114ee8360008461151c565b505050565b600a6006546115029190611b6d565b60065461150f9190611b17565b600681905550565b505050565b505050565b6000813590506115308161229b565b92915050565b600081359050611545816122b2565b92915050565b60006020828403121561156157611560611ee4565b5b600061156f84828501611521565b91505092915050565b6000806040838503121561158f5761158e611ee4565b5b600061159d85828601611521565b92505060206115ae85828601611521565b9150509250929050565b6000806000606084860312156115d1576115d0611ee4565b5b60006115df86828701611521565b93505060206115f086828701611521565b925050604061160186828701611536565b9150509250925092565b6000806040838503121561162257611621611ee4565b5b600061163085828601611521565b925050602061164185828601611536565b9150509250929050565b60006020828403121561166157611660611ee4565b5b600061166f84828501611536565b91505092915050565b61168181611daf565b82525050565b600061169282611afb565b61169c8185611b06565b93506116ac818560208601611df2565b6116b581611ee9565b840191505092915050565b60006116cd602383611b06565b91506116d882611f07565b604082019050919050565b60006116f0602283611b06565b91506116fb82611f56565b604082019050919050565b6000611713602283611b06565b915061171e82611fa5565b604082019050919050565b6000611736600a83611b06565b915061174182611ff4565b602082019050919050565b6000611759601a83611b06565b91506117648261201d565b602082019050919050565b600061177c602683611b06565b915061178782612046565b604082019050919050565b600061179f601b83611b06565b91506117aa82612095565b602082019050919050565b60006117c2602883611b06565b91506117cd826120be565b604082019050919050565b60006117e5602183611b06565b91506117f08261210d565b604082019050919050565b6000611808602583611b06565b91506118138261215c565b604082019050919050565b600061182b601883611b06565b9150611836826121ab565b602082019050919050565b600061184e602483611b06565b9150611859826121d4565b604082019050919050565b6000611871602583611b06565b915061187c82612223565b604082019050919050565b6000611894601f83611b06565b915061189f82612272565b602082019050919050565b6118b381611ddb565b82525050565b6118c281611de5565b82525050565b60006020820190506118dd6000830184611678565b92915050565b600060208201905081810360008301526118fd8184611687565b905092915050565b6000602082019050818103600083015261191e816116c0565b9050919050565b6000602082019050818103600083015261193e816116e3565b9050919050565b6000602082019050818103600083015261195e81611706565b9050919050565b6000602082019050818103600083015261197e81611729565b9050919050565b6000602082019050818103600083015261199e8161174c565b9050919050565b600060208201905081810360008301526119be8161176f565b9050919050565b600060208201905081810360008301526119de81611792565b9050919050565b600060208201905081810360008301526119fe816117b5565b9050919050565b60006020820190508181036000830152611a1e816117d8565b9050919050565b60006020820190508181036000830152611a3e816117fb565b9050919050565b60006020820190508181036000830152611a5e8161181e565b9050919050565b60006020820190508181036000830152611a7e81611841565b9050919050565b60006020820190508181036000830152611a9e81611864565b9050919050565b60006020820190508181036000830152611abe81611887565b9050919050565b6000602082019050611ada60008301846118aa565b92915050565b6000602082019050611af560008301846118b9565b92915050565b600081519050919050565b600082825260208201905092915050565b6000611b2282611ddb565b9150611b2d83611ddb565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611b6257611b61611e57565b5b828201905092915050565b6000611b7882611ddb565b9150611b8383611ddb565b925082611b9357611b92611e86565b5b828204905092915050565b6000808291508390505b6001851115611be857808604811115611bc457611bc3611e57565b5b6001851615611bd35780820291505b8081029050611be185611efa565b9450611ba8565b94509492505050565b6000611bfc82611ddb565b9150611c0783611ddb565b9250611c347fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484611c3c565b905092915050565b600082611c4c5760019050611d08565b81611c5a5760009050611d08565b8160018114611c705760028114611c7a57611ca9565b6001915050611d08565b60ff841115611c8c57611c8b611e57565b5b8360020a915084821115611ca357611ca2611e57565b5b50611d08565b5060208310610133831016604e8410600b8410161715611cde5782820a905083811115611cd957611cd8611e57565b5b611d08565b611ceb8484846001611b9e565b92509050818404811115611d0257611d01611e57565b5b81810290505b9392505050565b6000611d1a82611ddb565b9150611d2583611ddb565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611d5e57611d5d611e57565b5b828202905092915050565b6000611d7482611ddb565b9150611d7f83611ddb565b925082821015611d9257611d91611e57565b5b828203905092915050565b6000611da882611dbb565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60005b83811015611e10578082015181840152602081019050611df5565b83811115611e1f576000848401525b50505050565b60006002820490506001821680611e3d57607f821691505b60208210811415611e5157611e50611eb5565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600080fd5b6000601f19601f8301169050919050565b60008160011c9050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60008201527f6365000000000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b7f61646d696e206f6e6c7900000000000000000000000000000000000000000000600082015250565b7f6e6f7420656e6f7567682062616c616e636520746f2073656c6c000000000000600082015250565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206260008201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b7f63616e27742062757920616e792077697468207a65726f204554480000000000600082015250565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206160008201527f6c6c6f77616e6365000000000000000000000000000000000000000000000000602082015250565b7f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360008201527f7300000000000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b7f63616e27742066756e642077697468203020616d6f756e740000000000000000600082015250565b7f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b6122a481611d9d565b81146122af57600080fd5b50565b6122bb81611ddb565b81146122c657600080fd5b5056fea26469706673582212202dcebad1ec3d3e5f18a1515afc1cc4e57519136fdaaeff4701a5b756adeaff0364736f6c63430008050033"

// DeployKbd deploys a new Ethereum contract, binding an instance of Kbd to it.
func DeployKbd(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Kbd, error) {
	parsed, err := abi.JSON(strings.NewReader(KbdABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KbdBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Kbd{KbdCaller: KbdCaller{contract: contract}, KbdTransactor: KbdTransactor{contract: contract}, KbdFilterer: KbdFilterer{contract: contract}}, nil
}

// Kbd is an auto generated Go binding around an Ethereum contract.
type Kbd struct {
	KbdCaller     // Read-only binding to the contract
	KbdTransactor // Write-only binding to the contract
	KbdFilterer   // Log filterer for contract events
}

// KbdCaller is an auto generated read-only Go binding around an Ethereum contract.
type KbdCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KbdTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KbdTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KbdFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KbdFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KbdSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KbdSession struct {
	Contract     *Kbd              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KbdCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KbdCallerSession struct {
	Contract *KbdCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// KbdTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KbdTransactorSession struct {
	Contract     *KbdTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KbdRaw is an auto generated low-level Go binding around an Ethereum contract.
type KbdRaw struct {
	Contract *Kbd // Generic contract binding to access the raw methods on
}

// KbdCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KbdCallerRaw struct {
	Contract *KbdCaller // Generic read-only contract binding to access the raw methods on
}

// KbdTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KbdTransactorRaw struct {
	Contract *KbdTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKbd creates a new instance of Kbd, bound to a specific deployed contract.
func NewKbd(address common.Address, backend bind.ContractBackend) (*Kbd, error) {
	contract, err := bindKbd(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Kbd{KbdCaller: KbdCaller{contract: contract}, KbdTransactor: KbdTransactor{contract: contract}, KbdFilterer: KbdFilterer{contract: contract}}, nil
}

// NewKbdCaller creates a new read-only instance of Kbd, bound to a specific deployed contract.
func NewKbdCaller(address common.Address, caller bind.ContractCaller) (*KbdCaller, error) {
	contract, err := bindKbd(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KbdCaller{contract: contract}, nil
}

// NewKbdTransactor creates a new write-only instance of Kbd, bound to a specific deployed contract.
func NewKbdTransactor(address common.Address, transactor bind.ContractTransactor) (*KbdTransactor, error) {
	contract, err := bindKbd(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KbdTransactor{contract: contract}, nil
}

// NewKbdFilterer creates a new log filterer instance of Kbd, bound to a specific deployed contract.
func NewKbdFilterer(address common.Address, filterer bind.ContractFilterer) (*KbdFilterer, error) {
	contract, err := bindKbd(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KbdFilterer{contract: contract}, nil
}

// bindKbd binds a generic wrapper to an already deployed contract.
func bindKbd(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KbdABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kbd *KbdRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Kbd.Contract.KbdCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kbd *KbdRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kbd.Contract.KbdTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kbd *KbdRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kbd.Contract.KbdTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kbd *KbdCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Kbd.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kbd *KbdTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kbd.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kbd *KbdTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kbd.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Kbd *KbdCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Kbd.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Kbd *KbdSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Kbd.Contract.Allowance(&_Kbd.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Kbd *KbdCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Kbd.Contract.Allowance(&_Kbd.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Kbd *KbdCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Kbd.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Kbd *KbdSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Kbd.Contract.BalanceOf(&_Kbd.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Kbd *KbdCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Kbd.Contract.BalanceOf(&_Kbd.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Kbd *KbdCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Kbd.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Kbd *KbdSession) Decimals() (uint8, error) {
	return _Kbd.Contract.Decimals(&_Kbd.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Kbd *KbdCallerSession) Decimals() (uint8, error) {
	return _Kbd.Contract.Decimals(&_Kbd.CallOpts)
}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() view returns(uint256)
func (_Kbd *KbdCaller) ExchangeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Kbd.contract.Call(opts, &out, "exchangeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() view returns(uint256)
func (_Kbd *KbdSession) ExchangeRate() (*big.Int, error) {
	return _Kbd.Contract.ExchangeRate(&_Kbd.CallOpts)
}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() view returns(uint256)
func (_Kbd *KbdCallerSession) ExchangeRate() (*big.Int, error) {
	return _Kbd.Contract.ExchangeRate(&_Kbd.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Kbd *KbdCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Kbd.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Kbd *KbdSession) Name() (string, error) {
	return _Kbd.Contract.Name(&_Kbd.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Kbd *KbdCallerSession) Name() (string, error) {
	return _Kbd.Contract.Name(&_Kbd.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Kbd *KbdCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Kbd.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Kbd *KbdSession) Symbol() (string, error) {
	return _Kbd.Contract.Symbol(&_Kbd.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Kbd *KbdCallerSession) Symbol() (string, error) {
	return _Kbd.Contract.Symbol(&_Kbd.CallOpts)
}

// ThisBalance is a free data retrieval call binding the contract method 0xe3faecc2.
//
// Solidity: function this_balance() view returns(uint256)
func (_Kbd *KbdCaller) ThisBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Kbd.contract.Call(opts, &out, "this_balance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ThisBalance is a free data retrieval call binding the contract method 0xe3faecc2.
//
// Solidity: function this_balance() view returns(uint256)
func (_Kbd *KbdSession) ThisBalance() (*big.Int, error) {
	return _Kbd.Contract.ThisBalance(&_Kbd.CallOpts)
}

// ThisBalance is a free data retrieval call binding the contract method 0xe3faecc2.
//
// Solidity: function this_balance() view returns(uint256)
func (_Kbd *KbdCallerSession) ThisBalance() (*big.Int, error) {
	return _Kbd.Contract.ThisBalance(&_Kbd.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Kbd *KbdCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Kbd.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Kbd *KbdSession) TotalSupply() (*big.Int, error) {
	return _Kbd.Contract.TotalSupply(&_Kbd.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Kbd *KbdCallerSession) TotalSupply() (*big.Int, error) {
	return _Kbd.Contract.TotalSupply(&_Kbd.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Kbd *KbdTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kbd.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Kbd *KbdSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kbd.Contract.Approve(&_Kbd.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Kbd *KbdTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kbd.Contract.Approve(&_Kbd.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_Kbd *KbdTransactor) Burn(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kbd.contract.Transact(opts, "burn", account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_Kbd *KbdSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kbd.Contract.Burn(&_Kbd.TransactOpts, account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_Kbd *KbdTransactorSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kbd.Contract.Burn(&_Kbd.TransactOpts, account, amount)
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() payable returns()
func (_Kbd *KbdTransactor) Buy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kbd.contract.Transact(opts, "buy")
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() payable returns()
func (_Kbd *KbdSession) Buy() (*types.Transaction, error) {
	return _Kbd.Contract.Buy(&_Kbd.TransactOpts)
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() payable returns()
func (_Kbd *KbdTransactorSession) Buy() (*types.Transaction, error) {
	return _Kbd.Contract.Buy(&_Kbd.TransactOpts)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Kbd *KbdTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Kbd.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Kbd *KbdSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Kbd.Contract.DecreaseAllowance(&_Kbd.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Kbd *KbdTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Kbd.Contract.DecreaseAllowance(&_Kbd.TransactOpts, spender, subtractedValue)
}

// Fund is a paid mutator transaction binding the contract method 0xb60d4288.
//
// Solidity: function fund() payable returns()
func (_Kbd *KbdTransactor) Fund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kbd.contract.Transact(opts, "fund")
}

// Fund is a paid mutator transaction binding the contract method 0xb60d4288.
//
// Solidity: function fund() payable returns()
func (_Kbd *KbdSession) Fund() (*types.Transaction, error) {
	return _Kbd.Contract.Fund(&_Kbd.TransactOpts)
}

// Fund is a paid mutator transaction binding the contract method 0xb60d4288.
//
// Solidity: function fund() payable returns()
func (_Kbd *KbdTransactorSession) Fund() (*types.Transaction, error) {
	return _Kbd.Contract.Fund(&_Kbd.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Kbd *KbdTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Kbd.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Kbd *KbdSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Kbd.Contract.IncreaseAllowance(&_Kbd.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Kbd *KbdTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Kbd.Contract.IncreaseAllowance(&_Kbd.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_Kbd *KbdTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kbd.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_Kbd *KbdSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kbd.Contract.Mint(&_Kbd.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_Kbd *KbdTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kbd.Contract.Mint(&_Kbd.TransactOpts, account, amount)
}

// Sell is a paid mutator transaction binding the contract method 0xe4849b32.
//
// Solidity: function sell(uint256 amount) returns()
func (_Kbd *KbdTransactor) Sell(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Kbd.contract.Transact(opts, "sell", amount)
}

// Sell is a paid mutator transaction binding the contract method 0xe4849b32.
//
// Solidity: function sell(uint256 amount) returns()
func (_Kbd *KbdSession) Sell(amount *big.Int) (*types.Transaction, error) {
	return _Kbd.Contract.Sell(&_Kbd.TransactOpts, amount)
}

// Sell is a paid mutator transaction binding the contract method 0xe4849b32.
//
// Solidity: function sell(uint256 amount) returns()
func (_Kbd *KbdTransactorSession) Sell(amount *big.Int) (*types.Transaction, error) {
	return _Kbd.Contract.Sell(&_Kbd.TransactOpts, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Kbd *KbdTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kbd.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Kbd *KbdSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kbd.Contract.Transfer(&_Kbd.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Kbd *KbdTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kbd.Contract.Transfer(&_Kbd.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Kbd *KbdTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kbd.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Kbd *KbdSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kbd.Contract.TransferFrom(&_Kbd.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Kbd *KbdTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kbd.Contract.TransferFrom(&_Kbd.TransactOpts, sender, recipient, amount)
}

// KbdApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Kbd contract.
type KbdApprovalIterator struct {
	Event *KbdApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KbdApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KbdApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KbdApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KbdApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KbdApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KbdApproval represents a Approval event raised by the Kbd contract.
type KbdApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Kbd *KbdFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*KbdApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Kbd.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &KbdApprovalIterator{contract: _Kbd.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Kbd *KbdFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *KbdApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Kbd.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KbdApproval)
				if err := _Kbd.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Kbd *KbdFilterer) ParseApproval(log types.Log) (*KbdApproval, error) {
	event := new(KbdApproval)
	if err := _Kbd.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KbdTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Kbd contract.
type KbdTransferIterator struct {
	Event *KbdTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KbdTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KbdTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KbdTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KbdTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KbdTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KbdTransfer represents a Transfer event raised by the Kbd contract.
type KbdTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Kbd *KbdFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*KbdTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Kbd.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &KbdTransferIterator{contract: _Kbd.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Kbd *KbdFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *KbdTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Kbd.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KbdTransfer)
				if err := _Kbd.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Kbd *KbdFilterer) ParseTransfer(log types.Log) (*KbdTransfer, error) {
	event := new(KbdTransfer)
	if err := _Kbd.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

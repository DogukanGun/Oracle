// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package oracle

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// WalletAsset is an auto generated low-level Go binding around an user-defined struct.
type WalletAsset struct {
	Sign   string
	Amount *big.Int
}

// WalletMetaData contains all meta data concerning the Wallet contract.
var WalletMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sign\",\"type\":\"string\"}],\"name\":\"addNewAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str1\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"str2\",\"type\":\"string\"}],\"name\":\"compareStrings\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"passwordSaver\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"password\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"username\",\"type\":\"string\"}],\"name\":\"createUserAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sign\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddresss\",\"type\":\"address\"}],\"name\":\"getUserAssets\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"sign\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structWallet.Asset[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sign\",\"type\":\"string\"}],\"name\":\"isAssetExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userWallets\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"passwordSaver\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"password\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"assetCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sign\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506119f1806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063765ff0611161005b578063765ff061146101135780638571faff14610143578063bed34bba1461015f578063f79433381461018f57610088565b806320f02ec71461008d578063445f414a146100a957806345512594146100c557806363e6ffdd146100e1575b600080fd5b6100a760048036038101906100a29190610efb565b6101bf565b005b6100c360048036038101906100be9190611012565b61031a565b005b6100df60048036038101906100da9190610efb565b6104ad565b005b6100fb60048036038101906100f69190611086565b610608565b60405161010a93929190611141565b60405180910390f35b61012d60048036038101906101289190611086565b610742565b60405161013a9190611300565b60405180910390f35b61015d60048036038101906101589190611322565b610933565b005b610179600480360381019061017491906113c9565b610ad2565b604051610186919061145c565b60405180910390f35b6101a960048036038101906101a49190611477565b610b2b565b6040516101b6919061145c565b60405180910390f35b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060005b8160020154811015610313576102ba826003016000838152602001908152602001600020600001805461023690611515565b80601f016020809104026020016040519081016040528092919081815260200182805461026290611515565b80156102af5780601f10610284576101008083540402835291602001916102af565b820191906000526020600020905b81548152906001019060200180831161029257829003601f168201915b505050505085610ad2565b156103005782826003016000838152602001908152602001600020600201546102e39190611575565b826003016000838152602001908152602001600020600201819055505b808061030b906115a9565b915050610204565b5050505050565b60008060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506103ab8585858080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505084610b2b565b6104a6576103b7610cda565b84848080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050816000018190525060008160400181815250508082600301600084600201548152602001908152602001600020600082015181600001908161043e919061179d565b5060208201518160010160006101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff160217905550604082015181600201559050506001826002015461049c919061186f565b8260020181905550505b5050505050565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060005b8160020154811015610601576105a8826003016000838152602001908152602001600020600001805461052490611515565b80601f016020809104026020016040519081016040528092919081815260200182805461055090611515565b801561059d5780601f106105725761010080835404028352916020019161059d565b820191906000526020600020905b81548152906001019060200180831161058057829003601f168201915b505050505085610ad2565b156105ee5782826003016000838152602001908152602001600020600201546105d1919061186f565b826003016000838152602001908152602001600020600201819055505b80806105f9906115a9565b9150506104f2565b5050505050565b600060205280600052604060002060009150905080600001805461062b90611515565b80601f016020809104026020016040519081016040528092919081815260200182805461065790611515565b80156106a45780601f10610679576101008083540402835291602001916106a4565b820191906000526020600020905b81548152906001019060200180831161068757829003601f168201915b5050505050908060010180546106b990611515565b80601f01602080910402602001604051908101604052809291908181526020018280546106e590611515565b80156107325780601f1061070757610100808354040283529160200191610732565b820191906000526020600020905b81548152906001019060200180831161071557829003601f168201915b5050505050908060020154905083565b606060008060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000816002015467ffffffffffffffff8111156107a6576107a5610d9a565b5b6040519080825280602002602001820160405280156107df57816020015b6107cc610cda565b8152602001906001900390816107c45790505b50905060005b82600201548110156109285782600301600082815260200190815260200160002060405180606001604052908160008201805461082190611515565b80601f016020809104026020016040519081016040528092919081815260200182805461084d90611515565b801561089a5780601f1061086f5761010080835404028352916020019161089a565b820191906000526020600020905b81548152906001019060200180831161087d57829003601f168201915b505050505081526020016001820160009054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff16815260200160028201548152505082828151811061090a576109096118a3565b5b60200260200101819052508080610920906115a9565b9150506107e5565b508092505050919050565b600061093e82610cbb565b905060008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050610a2781600101805461099490611515565b80601f01602080910402602001604051908101604052809291908181526020018280546109c090611515565b8015610a0d5780601f106109e257610100808354040283529160200191610a0d565b820191906000526020600020905b8154815290600101906020018083116109f057829003601f168201915b505050505060405180602001604052806000815250610ad2565b610a66576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a5d9061191e565b60405180910390fd5b83816001019081610a77919061179d565b5084816000019081610a89919061179d565b50600081600201819055507f234cf33b32239d80b54161e2396c80cdeaf4d34161300e54d8bc01eb7c0ea55382604051610ac3919061194d565b60405180910390a15050505050565b600081604051602001610ae591906119a4565b6040516020818303038152906040528051906020012083604051602001610b0c91906119a4565b6040516020818303038152906040528051906020012014905092915050565b6000806000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060005b8160020154811015610cad57610c278260030160008381526020019081526020016000206000018054610ba390611515565b80601f0160208091040260200160405190810160405280929190818152602001828054610bcf90611515565b8015610c1c5780601f10610bf157610100808354040283529160200191610c1c565b820191906000526020600020905b815481529060010190602001808311610bff57829003601f168201915b505050505086610ad2565b8015610c8a575081600301600082815260200190815260200160002060010160009054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff16846fffffffffffffffffffffffffffffffff16145b15610c9a57600192505050610cb4565b8080610ca5906115a9565b915050610b71565b5060009150505b9392505050565b6000808290506000818051906020012060001c90508092505050919050565b60405180606001604052806060815260200160006fffffffffffffffffffffffffffffffff168152602001600081525090565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610d4c82610d21565b9050919050565b610d5c81610d41565b8114610d6757600080fd5b50565b600081359050610d7981610d53565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610dd282610d89565b810181811067ffffffffffffffff82111715610df157610df0610d9a565b5b80604052505050565b6000610e04610d0d565b9050610e108282610dc9565b919050565b600067ffffffffffffffff821115610e3057610e2f610d9a565b5b610e3982610d89565b9050602081019050919050565b82818337600083830152505050565b6000610e68610e6384610e15565b610dfa565b905082815260208101848484011115610e8457610e83610d84565b5b610e8f848285610e46565b509392505050565b600082601f830112610eac57610eab610d7f565b5b8135610ebc848260208601610e55565b91505092915050565b6000819050919050565b610ed881610ec5565b8114610ee357600080fd5b50565b600081359050610ef581610ecf565b92915050565b600080600060608486031215610f1457610f13610d17565b5b6000610f2286828701610d6a565b935050602084013567ffffffffffffffff811115610f4357610f42610d1c565b5b610f4f86828701610e97565b9250506040610f6086828701610ee6565b9150509250925092565b600080fd5b600080fd5b60008083601f840112610f8a57610f89610d7f565b5b8235905067ffffffffffffffff811115610fa757610fa6610f6a565b5b602083019150836001820283011115610fc357610fc2610f6f565b5b9250929050565b60006fffffffffffffffffffffffffffffffff82169050919050565b610fef81610fca565b8114610ffa57600080fd5b50565b60008135905061100c81610fe6565b92915050565b6000806000806060858703121561102c5761102b610d17565b5b600061103a87828801610d6a565b945050602085013567ffffffffffffffff81111561105b5761105a610d1c565b5b61106787828801610f74565b9350935050604061107a87828801610ffd565b91505092959194509250565b60006020828403121561109c5761109b610d17565b5b60006110aa84828501610d6a565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156110ed5780820151818401526020810190506110d2565b60008484015250505050565b6000611104826110b3565b61110e81856110be565b935061111e8185602086016110cf565b61112781610d89565b840191505092915050565b61113b81610ec5565b82525050565b6000606082019050818103600083015261115b81866110f9565b9050818103602083015261116f81856110f9565b905061117e6040830184611132565b949350505050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600082825260208201905092915050565b60006111ce826110b3565b6111d881856111b2565b93506111e88185602086016110cf565b6111f181610d89565b840191505092915050565b61120581610fca565b82525050565b61121481610ec5565b82525050565b6000606083016000830151848203600086015261123782826111c3565b915050602083015161124c60208601826111fc565b50604083015161125f604086018261120b565b508091505092915050565b6000611276838361121a565b905092915050565b6000602082019050919050565b600061129682611186565b6112a08185611191565b9350836020820285016112b2856111a2565b8060005b858110156112ee57848403895281516112cf858261126a565b94506112da8361127e565b925060208a019950506001810190506112b6565b50829750879550505050505092915050565b6000602082019050818103600083015261131a818461128b565b905092915050565b60008060006060848603121561133b5761133a610d17565b5b600084013567ffffffffffffffff81111561135957611358610d1c565b5b61136586828701610e97565b935050602084013567ffffffffffffffff81111561138657611385610d1c565b5b61139286828701610e97565b925050604084013567ffffffffffffffff8111156113b3576113b2610d1c565b5b6113bf86828701610e97565b9150509250925092565b600080604083850312156113e0576113df610d17565b5b600083013567ffffffffffffffff8111156113fe576113fd610d1c565b5b61140a85828601610e97565b925050602083013567ffffffffffffffff81111561142b5761142a610d1c565b5b61143785828601610e97565b9150509250929050565b60008115159050919050565b61145681611441565b82525050565b6000602082019050611471600083018461144d565b92915050565b6000806000606084860312156114905761148f610d17565b5b600061149e86828701610d6a565b935050602084013567ffffffffffffffff8111156114bf576114be610d1c565b5b6114cb86828701610e97565b92505060406114dc86828701610ffd565b9150509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061152d57607f821691505b6020821081036115405761153f6114e6565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061158082610ec5565b915061158b83610ec5565b92508282039050818111156115a3576115a2611546565b5b92915050565b60006115b482610ec5565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036115e6576115e5611546565b5b600182019050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026116537fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611616565b61165d8683611616565b95508019841693508086168417925050509392505050565b6000819050919050565b600061169a61169561169084610ec5565b611675565b610ec5565b9050919050565b6000819050919050565b6116b48361167f565b6116c86116c0826116a1565b848454611623565b825550505050565b600090565b6116dd6116d0565b6116e88184846116ab565b505050565b5b8181101561170c576117016000826116d5565b6001810190506116ee565b5050565b601f82111561175157611722816115f1565b61172b84611606565b8101602085101561173a578190505b61174e61174685611606565b8301826116ed565b50505b505050565b600082821c905092915050565b600061177460001984600802611756565b1980831691505092915050565b600061178d8383611763565b9150826002028217905092915050565b6117a6826110b3565b67ffffffffffffffff8111156117bf576117be610d9a565b5b6117c98254611515565b6117d4828285611710565b600060209050601f83116001811461180757600084156117f5578287015190505b6117ff8582611781565b865550611867565b601f198416611815866115f1565b60005b8281101561183d57848901518255600182019150602085019450602081019050611818565b8683101561185a5784890151611856601f891682611763565b8355505b6001600288020188555050505b505050505050565b600061187a82610ec5565b915061188583610ec5565b925082820190508082111561189d5761189c611546565b5b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f557365722069732063726561746564206265666f726500000000000000000000600082015250565b60006119086016836110be565b9150611913826118d2565b602082019050919050565b60006020820190508181036000830152611937816118fb565b9050919050565b61194781610d41565b82525050565b6000602082019050611962600083018461193e565b92915050565b600081905092915050565b600061197e826110b3565b6119888185611968565b93506119988185602086016110cf565b80840191505092915050565b60006119b08284611973565b91508190509291505056fea2646970667358221220247080255faf858c31103f924dc1171b69d29a321ffe07e153450865ab5dd93164736f6c63430008130033",
}

// WalletABI is the input ABI used to generate the binding from.
// Deprecated: Use WalletMetaData.ABI instead.
var WalletABI = WalletMetaData.ABI

// WalletBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WalletMetaData.Bin instead.
var WalletBin = WalletMetaData.Bin

// DeployWallet deploys a new Ethereum contract, binding an instance of Wallet to it.
func DeployWallet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Wallet, error) {
	parsed, err := WalletMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WalletBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Wallet{WalletCaller: WalletCaller{contract: contract}, WalletTransactor: WalletTransactor{contract: contract}, WalletFilterer: WalletFilterer{contract: contract}}, nil
}

// Wallet is an auto generated Go binding around an Ethereum contract.
type Wallet struct {
	WalletCaller     // Read-only binding to the contract
	WalletTransactor // Write-only binding to the contract
	WalletFilterer   // Log filterer for contract events
}

// WalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletSession struct {
	Contract     *Wallet           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletCallerSession struct {
	Contract *WalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// WalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletTransactorSession struct {
	Contract     *WalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletRaw struct {
	Contract *Wallet // Generic contract binding to access the raw methods on
}

// WalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletCallerRaw struct {
	Contract *WalletCaller // Generic read-only contract binding to access the raw methods on
}

// WalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletTransactorRaw struct {
	Contract *WalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWallet creates a new instance of Wallet, bound to a specific deployed contract.
func NewWallet(address common.Address, backend bind.ContractBackend) (*Wallet, error) {
	contract, err := bindWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Wallet{WalletCaller: WalletCaller{contract: contract}, WalletTransactor: WalletTransactor{contract: contract}, WalletFilterer: WalletFilterer{contract: contract}}, nil
}

// NewWalletCaller creates a new read-only instance of Wallet, bound to a specific deployed contract.
func NewWalletCaller(address common.Address, caller bind.ContractCaller) (*WalletCaller, error) {
	contract, err := bindWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletCaller{contract: contract}, nil
}

// NewWalletTransactor creates a new write-only instance of Wallet, bound to a specific deployed contract.
func NewWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletTransactor, error) {
	contract, err := bindWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletTransactor{contract: contract}, nil
}

// NewWalletFilterer creates a new log filterer instance of Wallet, bound to a specific deployed contract.
func NewWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletFilterer, error) {
	contract, err := bindWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletFilterer{contract: contract}, nil
}

// bindWallet binds a generic wrapper to an already deployed contract.
func bindWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WalletMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wallet *WalletRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Wallet.Contract.WalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wallet *WalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.Contract.WalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wallet *WalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wallet.Contract.WalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wallet *WalletCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Wallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wallet *WalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wallet *WalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wallet.Contract.contract.Transact(opts, method, params...)
}

// CompareStrings is a free data retrieval call binding the contract method 0xbed34bba.
//
// Solidity: function compareStrings(string str1, string str2) pure returns(bool)
func (_Wallet *WalletCaller) CompareStrings(opts *bind.CallOpts, str1 string, str2 string) (bool, error) {
	var out []interface{}
	err := _Wallet.contract.Call(opts, &out, "compareStrings", str1, str2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CompareStrings is a free data retrieval call binding the contract method 0xbed34bba.
//
// Solidity: function compareStrings(string str1, string str2) pure returns(bool)
func (_Wallet *WalletSession) CompareStrings(str1 string, str2 string) (bool, error) {
	return _Wallet.Contract.CompareStrings(&_Wallet.CallOpts, str1, str2)
}

// CompareStrings is a free data retrieval call binding the contract method 0xbed34bba.
//
// Solidity: function compareStrings(string str1, string str2) pure returns(bool)
func (_Wallet *WalletCallerSession) CompareStrings(str1 string, str2 string) (bool, error) {
	return _Wallet.Contract.CompareStrings(&_Wallet.CallOpts, str1, str2)
}

// GetUserAssets is a free data retrieval call binding the contract method 0x765ff061.
//
// Solidity: function getUserAssets(address userAddresss) view returns((string,uint256)[])
func (_Wallet *WalletCaller) GetUserAssets(opts *bind.CallOpts, userAddresss common.Address) ([]WalletAsset, error) {
	var out []interface{}
	err := _Wallet.contract.Call(opts, &out, "getUserAssets", userAddresss)

	if err != nil {
		return *new([]WalletAsset), err
	}

	out0 := *abi.ConvertType(out[0], new([]WalletAsset)).(*[]WalletAsset)

	return out0, err

}

// GetUserAssets is a free data retrieval call binding the contract method 0x765ff061.
//
// Solidity: function getUserAssets(address userAddresss) view returns((string,uint256)[])
func (_Wallet *WalletSession) GetUserAssets(userAddresss common.Address) ([]WalletAsset, error) {
	return _Wallet.Contract.GetUserAssets(&_Wallet.CallOpts, userAddresss)
}

// GetUserAssets is a free data retrieval call binding the contract method 0x765ff061.
//
// Solidity: function getUserAssets(address userAddresss) view returns((string,uint256)[])
func (_Wallet *WalletCallerSession) GetUserAssets(userAddresss common.Address) ([]WalletAsset, error) {
	return _Wallet.Contract.GetUserAssets(&_Wallet.CallOpts, userAddresss)
}

// IsAssetExist is a free data retrieval call binding the contract method 0x42bcaff7.
//
// Solidity: function isAssetExist(address userAddress, string sign) view returns(bool)
func (_Wallet *WalletCaller) IsAssetExist(opts *bind.CallOpts, userAddress common.Address, sign string) (bool, error) {
	var out []interface{}
	err := _Wallet.contract.Call(opts, &out, "isAssetExist", userAddress, sign)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAssetExist is a free data retrieval call binding the contract method 0x42bcaff7.
//
// Solidity: function isAssetExist(address userAddress, string sign) view returns(bool)
func (_Wallet *WalletSession) IsAssetExist(userAddress common.Address, sign string) (bool, error) {
	return _Wallet.Contract.IsAssetExist(&_Wallet.CallOpts, userAddress, sign)
}

// IsAssetExist is a free data retrieval call binding the contract method 0x42bcaff7.
//
// Solidity: function isAssetExist(address userAddress, string sign) view returns(bool)
func (_Wallet *WalletCallerSession) IsAssetExist(userAddress common.Address, sign string) (bool, error) {
	return _Wallet.Contract.IsAssetExist(&_Wallet.CallOpts, userAddress, sign)
}

// UserWallets is a free data retrieval call binding the contract method 0x63e6ffdd.
//
// Solidity: function userWallets(address ) view returns(string passwordSaver, string password, uint256 assetCount)
func (_Wallet *WalletCaller) UserWallets(opts *bind.CallOpts, arg0 common.Address) (struct {
	PasswordSaver string
	Password      string
	AssetCount    *big.Int
}, error) {
	var out []interface{}
	err := _Wallet.contract.Call(opts, &out, "userWallets", arg0)

	outstruct := new(struct {
		PasswordSaver string
		Password      string
		AssetCount    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PasswordSaver = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Password = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.AssetCount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserWallets is a free data retrieval call binding the contract method 0x63e6ffdd.
//
// Solidity: function userWallets(address ) view returns(string passwordSaver, string password, uint256 assetCount)
func (_Wallet *WalletSession) UserWallets(arg0 common.Address) (struct {
	PasswordSaver string
	Password      string
	AssetCount    *big.Int
}, error) {
	return _Wallet.Contract.UserWallets(&_Wallet.CallOpts, arg0)
}

// UserWallets is a free data retrieval call binding the contract method 0x63e6ffdd.
//
// Solidity: function userWallets(address ) view returns(string passwordSaver, string password, uint256 assetCount)
func (_Wallet *WalletCallerSession) UserWallets(arg0 common.Address) (struct {
	PasswordSaver string
	Password      string
	AssetCount    *big.Int
}, error) {
	return _Wallet.Contract.UserWallets(&_Wallet.CallOpts, arg0)
}

// AddNewAsset is a paid mutator transaction binding the contract method 0xe98b9469.
//
// Solidity: function addNewAsset(address userAddress, string sign) returns()
func (_Wallet *WalletTransactor) AddNewAsset(opts *bind.TransactOpts, userAddress common.Address, sign string) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "addNewAsset", userAddress, sign)
}

// AddNewAsset is a paid mutator transaction binding the contract method 0xe98b9469.
//
// Solidity: function addNewAsset(address userAddress, string sign) returns()
func (_Wallet *WalletSession) AddNewAsset(userAddress common.Address, sign string) (*types.Transaction, error) {
	return _Wallet.Contract.AddNewAsset(&_Wallet.TransactOpts, userAddress, sign)
}

// AddNewAsset is a paid mutator transaction binding the contract method 0xe98b9469.
//
// Solidity: function addNewAsset(address userAddress, string sign) returns()
func (_Wallet *WalletTransactorSession) AddNewAsset(userAddress common.Address, sign string) (*types.Transaction, error) {
	return _Wallet.Contract.AddNewAsset(&_Wallet.TransactOpts, userAddress, sign)
}

// CreateUserAddress is a paid mutator transaction binding the contract method 0x8571faff.
//
// Solidity: function createUserAddress(string passwordSaver, string password, string username) returns(address)
func (_Wallet *WalletTransactor) CreateUserAddress(opts *bind.TransactOpts, passwordSaver string, password string, username string) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "createUserAddress", passwordSaver, password, username)
}

// CreateUserAddress is a paid mutator transaction binding the contract method 0x8571faff.
//
// Solidity: function createUserAddress(string passwordSaver, string password, string username) returns(address)
func (_Wallet *WalletSession) CreateUserAddress(passwordSaver string, password string, username string) (*types.Transaction, error) {
	return _Wallet.Contract.CreateUserAddress(&_Wallet.TransactOpts, passwordSaver, password, username)
}

// CreateUserAddress is a paid mutator transaction binding the contract method 0x8571faff.
//
// Solidity: function createUserAddress(string passwordSaver, string password, string username) returns(address)
func (_Wallet *WalletTransactorSession) CreateUserAddress(passwordSaver string, password string, username string) (*types.Transaction, error) {
	return _Wallet.Contract.CreateUserAddress(&_Wallet.TransactOpts, passwordSaver, password, username)
}

// DepositAsset is a paid mutator transaction binding the contract method 0x45512594.
//
// Solidity: function depositAsset(address userAddress, string sign, uint256 amount) returns()
func (_Wallet *WalletTransactor) DepositAsset(opts *bind.TransactOpts, userAddress common.Address, sign string, amount *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "depositAsset", userAddress, sign, amount)
}

// DepositAsset is a paid mutator transaction binding the contract method 0x45512594.
//
// Solidity: function depositAsset(address userAddress, string sign, uint256 amount) returns()
func (_Wallet *WalletSession) DepositAsset(userAddress common.Address, sign string, amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.DepositAsset(&_Wallet.TransactOpts, userAddress, sign, amount)
}

// DepositAsset is a paid mutator transaction binding the contract method 0x45512594.
//
// Solidity: function depositAsset(address userAddress, string sign, uint256 amount) returns()
func (_Wallet *WalletTransactorSession) DepositAsset(userAddress common.Address, sign string, amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.DepositAsset(&_Wallet.TransactOpts, userAddress, sign, amount)
}

// WithdrawAsset is a paid mutator transaction binding the contract method 0x20f02ec7.
//
// Solidity: function withdrawAsset(address userAddress, string sign, uint256 amount) returns()
func (_Wallet *WalletTransactor) WithdrawAsset(opts *bind.TransactOpts, userAddress common.Address, sign string, amount *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "withdrawAsset", userAddress, sign, amount)
}

// WithdrawAsset is a paid mutator transaction binding the contract method 0x20f02ec7.
//
// Solidity: function withdrawAsset(address userAddress, string sign, uint256 amount) returns()
func (_Wallet *WalletSession) WithdrawAsset(userAddress common.Address, sign string, amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.WithdrawAsset(&_Wallet.TransactOpts, userAddress, sign, amount)
}

// WithdrawAsset is a paid mutator transaction binding the contract method 0x20f02ec7.
//
// Solidity: function withdrawAsset(address userAddress, string sign, uint256 amount) returns()
func (_Wallet *WalletTransactorSession) WithdrawAsset(userAddress common.Address, sign string, amount *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.WithdrawAsset(&_Wallet.TransactOpts, userAddress, sign, amount)
}

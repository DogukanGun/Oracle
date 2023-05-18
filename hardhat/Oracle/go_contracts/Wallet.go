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
}

// WalletABI is the input ABI used to generate the binding from.
// Deprecated: Use WalletMetaData.ABI instead.
var WalletABI = WalletMetaData.ABI

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

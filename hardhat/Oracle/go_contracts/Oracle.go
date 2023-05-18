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

// OracleMetaData contains all meta data concerning the Oracle contract.
var OracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"interestRate\",\"type\":\"uint256\"}],\"name\":\"decreaseTotalBarrowedLimitOf\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLendingPoolAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"interestRate\",\"type\":\"uint256\"}],\"name\":\"increaseTotalBarrowedLimitOf\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"interestRate\",\"type\":\"uint256\"}],\"name\":\"setInterestRate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lendingPoolAddress\",\"type\":\"address\"}],\"name\":\"setLendingPoolAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506107db806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80635f36a7421461005c5780635f84f3021461007a5780637ca5df52146100aa578063c4713aa4146100da578063ed3f827a1461010a575b600080fd5b610064610126565b6040516100719190610578565b60405180910390f35b610094600480360381019061008f91906105ce565b61014f565b6040516100a19190610616565b60405180910390f35b6100c460048036038101906100bf919061065d565b610284565b6040516100d19190610616565b60405180910390f35b6100f460048036038101906100ef919061065d565b6103bc565b6040516101019190610616565b60405180910390f35b610124600480360381019061011f919061069d565b6104f4565b005b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000808260405160240161016391906106d9565b6040516020818303038152906040527f5f84f302000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050905060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168260405161022a9190610765565b6000604051808303816000865af19150503d8060008114610267576040519150601f19603f3d011682016040523d82523d6000602084013e61026c565b606091505b505090508061027a57600080fd5b8092505050919050565b600080838360405160240161029a92919061077c565b6040516020818303038152906040527f7ca5df52000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050905060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16826040516103619190610765565b6000604051808303816000865af19150503d806000811461039e576040519150601f19603f3d011682016040523d82523d6000602084013e6103a3565b606091505b50509050806103b157600080fd5b809250505092915050565b60008083836040516024016103d292919061077c565b6040516020818303038152906040527fc4713aa4000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050905060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16826040516104999190610765565b6000604051808303816000865af19150503d80600081146104d6576040519150601f19603f3d011682016040523d82523d6000602084013e6104db565b606091505b50509050806104e957600080fd5b809250505092915050565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061056282610537565b9050919050565b61057281610557565b82525050565b600060208201905061058d6000830184610569565b92915050565b600080fd5b6000819050919050565b6105ab81610598565b81146105b657600080fd5b50565b6000813590506105c8816105a2565b92915050565b6000602082840312156105e4576105e3610593565b5b60006105f2848285016105b9565b91505092915050565b60008115159050919050565b610610816105fb565b82525050565b600060208201905061062b6000830184610607565b92915050565b61063a81610557565b811461064557600080fd5b50565b60008135905061065781610631565b92915050565b6000806040838503121561067457610673610593565b5b600061068285828601610648565b9250506020610693858286016105b9565b9150509250929050565b6000602082840312156106b3576106b2610593565b5b60006106c184828501610648565b91505092915050565b6106d381610598565b82525050565b60006020820190506106ee60008301846106ca565b92915050565b600081519050919050565b600081905092915050565b60005b8381101561072857808201518184015260208101905061070d565b60008484015250505050565b600061073f826106f4565b61074981856106ff565b935061075981856020860161070a565b80840191505092915050565b60006107718284610734565b915081905092915050565b60006040820190506107916000830185610569565b61079e60208301846106ca565b939250505056fea264697066735822122065110003b99c904d2d0a546a9a09ab4243ea44ece0c4986e3fead4c809a129b464736f6c63430008130033",
}

// OracleABI is the input ABI used to generate the binding from.
// Deprecated: Use OracleMetaData.ABI instead.
var OracleABI = OracleMetaData.ABI

// OracleBin is the compiled bytecode used for deploying new go_contracts.
// Deprecated: Use OracleMetaData.Bin instead.
var OracleBin = OracleMetaData.Bin

// DeployOracle deploys a new Ethereum contract, binding an instance of Oracle to it.
func DeployOracle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Oracle, error) {
	parsed, err := OracleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OracleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// Oracle is an auto generated Go binding around an Ethereum contract.
type Oracle struct {
	OracleCaller     // Read-only binding to the contract
	OracleTransactor // Write-only binding to the contract
	OracleFilterer   // Log filterer for contract events
}

// OracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleSession struct {
	Contract     *Oracle           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleCallerSession struct {
	Contract *OracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleTransactorSession struct {
	Contract     *OracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleRaw struct {
	Contract *Oracle // Generic contract binding to access the raw methods on
}

// OracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleCallerRaw struct {
	Contract *OracleCaller // Generic read-only contract binding to access the raw methods on
}

// OracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleTransactorRaw struct {
	Contract *OracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracle creates a new instance of Oracle, bound to a specific deployed contract.
func NewOracle(address common.Address, backend bind.ContractBackend) (*Oracle, error) {
	contract, err := bindOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// NewOracleCaller creates a new read-only instance of Oracle, bound to a specific deployed contract.
func NewOracleCaller(address common.Address, caller bind.ContractCaller) (*OracleCaller, error) {
	contract, err := bindOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleCaller{contract: contract}, nil
}

// NewOracleTransactor creates a new write-only instance of Oracle, bound to a specific deployed contract.
func NewOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleTransactor, error) {
	contract, err := bindOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleTransactor{contract: contract}, nil
}

// NewOracleFilterer creates a new log filterer instance of Oracle, bound to a specific deployed contract.
func NewOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleFilterer, error) {
	contract, err := bindOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleFilterer{contract: contract}, nil
}

// bindOracle binds a generic wrapper to an already deployed contract.
func bindOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.OracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transact(opts, method, params...)
}

// GetLendingPoolAddress is a free data retrieval call binding the contract method 0x5f36a742.
//
// Solidity: function getLendingPoolAddress() view returns(address)
func (_Oracle *OracleCaller) GetLendingPoolAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getLendingPoolAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLendingPoolAddress is a free data retrieval call binding the contract method 0x5f36a742.
//
// Solidity: function getLendingPoolAddress() view returns(address)
func (_Oracle *OracleSession) GetLendingPoolAddress() (common.Address, error) {
	return _Oracle.Contract.GetLendingPoolAddress(&_Oracle.CallOpts)
}

// GetLendingPoolAddress is a free data retrieval call binding the contract method 0x5f36a742.
//
// Solidity: function getLendingPoolAddress() view returns(address)
func (_Oracle *OracleCallerSession) GetLendingPoolAddress() (common.Address, error) {
	return _Oracle.Contract.GetLendingPoolAddress(&_Oracle.CallOpts)
}

// DecreaseTotalBarrowedLimitOf is a paid mutator transaction binding the contract method 0x7ca5df52.
//
// Solidity: function decreaseTotalBarrowedLimitOf(address user, uint256 interestRate) returns(bool)
func (_Oracle *OracleTransactor) DecreaseTotalBarrowedLimitOf(opts *bind.TransactOpts, user common.Address, interestRate *big.Int) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "decreaseTotalBarrowedLimitOf", user, interestRate)
}

// DecreaseTotalBarrowedLimitOf is a paid mutator transaction binding the contract method 0x7ca5df52.
//
// Solidity: function decreaseTotalBarrowedLimitOf(address user, uint256 interestRate) returns(bool)
func (_Oracle *OracleSession) DecreaseTotalBarrowedLimitOf(user common.Address, interestRate *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.DecreaseTotalBarrowedLimitOf(&_Oracle.TransactOpts, user, interestRate)
}

// DecreaseTotalBarrowedLimitOf is a paid mutator transaction binding the contract method 0x7ca5df52.
//
// Solidity: function decreaseTotalBarrowedLimitOf(address user, uint256 interestRate) returns(bool)
func (_Oracle *OracleTransactorSession) DecreaseTotalBarrowedLimitOf(user common.Address, interestRate *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.DecreaseTotalBarrowedLimitOf(&_Oracle.TransactOpts, user, interestRate)
}

// IncreaseTotalBarrowedLimitOf is a paid mutator transaction binding the contract method 0xc4713aa4.
//
// Solidity: function increaseTotalBarrowedLimitOf(address user, uint256 interestRate) returns(bool)
func (_Oracle *OracleTransactor) IncreaseTotalBarrowedLimitOf(opts *bind.TransactOpts, user common.Address, interestRate *big.Int) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "increaseTotalBarrowedLimitOf", user, interestRate)
}

// IncreaseTotalBarrowedLimitOf is a paid mutator transaction binding the contract method 0xc4713aa4.
//
// Solidity: function increaseTotalBarrowedLimitOf(address user, uint256 interestRate) returns(bool)
func (_Oracle *OracleSession) IncreaseTotalBarrowedLimitOf(user common.Address, interestRate *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.IncreaseTotalBarrowedLimitOf(&_Oracle.TransactOpts, user, interestRate)
}

// IncreaseTotalBarrowedLimitOf is a paid mutator transaction binding the contract method 0xc4713aa4.
//
// Solidity: function increaseTotalBarrowedLimitOf(address user, uint256 interestRate) returns(bool)
func (_Oracle *OracleTransactorSession) IncreaseTotalBarrowedLimitOf(user common.Address, interestRate *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.IncreaseTotalBarrowedLimitOf(&_Oracle.TransactOpts, user, interestRate)
}

// SetInterestRate is a paid mutator transaction binding the contract method 0x5f84f302.
//
// Solidity: function setInterestRate(uint256 interestRate) returns(bool)
func (_Oracle *OracleTransactor) SetInterestRate(opts *bind.TransactOpts, interestRate *big.Int) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "setInterestRate", interestRate)
}

// SetInterestRate is a paid mutator transaction binding the contract method 0x5f84f302.
//
// Solidity: function setInterestRate(uint256 interestRate) returns(bool)
func (_Oracle *OracleSession) SetInterestRate(interestRate *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.SetInterestRate(&_Oracle.TransactOpts, interestRate)
}

// SetInterestRate is a paid mutator transaction binding the contract method 0x5f84f302.
//
// Solidity: function setInterestRate(uint256 interestRate) returns(bool)
func (_Oracle *OracleTransactorSession) SetInterestRate(interestRate *big.Int) (*types.Transaction, error) {
	return _Oracle.Contract.SetInterestRate(&_Oracle.TransactOpts, interestRate)
}

// SetLendingPoolAddress is a paid mutator transaction binding the contract method 0xed3f827a.
//
// Solidity: function setLendingPoolAddress(address _lendingPoolAddress) returns()
func (_Oracle *OracleTransactor) SetLendingPoolAddress(opts *bind.TransactOpts, _lendingPoolAddress common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "setLendingPoolAddress", _lendingPoolAddress)
}

// SetLendingPoolAddress is a paid mutator transaction binding the contract method 0xed3f827a.
//
// Solidity: function setLendingPoolAddress(address _lendingPoolAddress) returns()
func (_Oracle *OracleSession) SetLendingPoolAddress(_lendingPoolAddress common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.SetLendingPoolAddress(&_Oracle.TransactOpts, _lendingPoolAddress)
}

// SetLendingPoolAddress is a paid mutator transaction binding the contract method 0xed3f827a.
//
// Solidity: function setLendingPoolAddress(address _lendingPoolAddress) returns()
func (_Oracle *OracleTransactorSession) SetLendingPoolAddress(_lendingPoolAddress common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.SetLendingPoolAddress(&_Oracle.TransactOpts, _lendingPoolAddress)
}

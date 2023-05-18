package main

import (
	"Oracle/contracts"
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
	_ "log"
	"math/big"
	"strings"
)

func runContract() {
	lendingPoolAddress, err, contractInstance, _publicAddress, res := createFunctionRequirements()
	resultOfSet, err := contractInstance.SetLendingPoolAddress(res, common.HexToAddress(lendingPoolAddress))
	if err != nil {
		// Handle error
		fmt.Println(err)
	}
	fmt.Println(resultOfSet)
	result, err := contractInstance.GetLendingPoolAddress(&bind.CallOpts{
		From:    _publicAddress,
		Context: context.Background(),
	})
	if err != nil {
		// Handle error
		fmt.Println(err)
	}
	fmt.Println(result)
}

func createFunctionRequirements() (string, error, *oracle.Oracle, common.Address, *bind.TransactOpts) {
	contractAddress := "0x5FbDB2315678afecb367f032d93F642f64180aa3"
	lendingPoolAddress := "0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512"
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		// Handle error
	}

	address := common.HexToAddress(contractAddress)
	abiFile, err := ioutil.ReadFile("Oracle.abi")
	contractAbi, err := abi.JSON(strings.NewReader(string(abiFile)))
	if err != nil {
		// Handle error
		fmt.Println(err)
	}

	fmt.Println(address)
	fmt.Println(client)
	fmt.Println(contractAbi)

	contractInstance, err := oracle.NewOracle(address, client)
	if err != nil {
		// Handle error
	}

	fmt.Println(contractInstance)
	privateKey := "0x701b615bbdfb9de65240bc28bd21bbc0d996645a3dd57e7b12bc2bdf6f192c82"
	_, err = GetAddressFromPrivateKey(privateKey)
	if err != nil {
		// Handle error
	}
	_privateKey, _, _publicAddress, _ := GenerateKeypairFromPrivateKeyHex(privateKey)
	res, _ := BuildTransactionOptions(client, _publicAddress, _privateKey, 300000)
	fmt.Println(res)
	return lendingPoolAddress, err, contractInstance, _publicAddress, res
}

func BuildTransactionOptions(client *ethclient.Client, fromAddress common.Address, prvKey *ecdsa.PrivateKey, gasLimit uint64) (*bind.TransactOpts, error) {

	// Retrieve the chainID
	chainID, cIDErr := client.ChainID(context.Background())

	if cIDErr != nil {
		return nil, cIDErr
	}

	// Retrieve suggested gas price
	gasPrice, gErr := client.SuggestGasPrice(context.Background())

	if gErr != nil {
		return nil, gErr
	}

	// Create empty options object
	txOptions, txOptErr := bind.NewKeyedTransactorWithChainID(prvKey, chainID)

	if txOptErr != nil {
		return nil, txOptErr
	}

	txOptions.Value = big.NewInt(0) // in wei
	txOptions.GasLimit = gasLimit   // in units
	txOptions.GasPrice = gasPrice

	return txOptions, nil
}

func GenerateKeypairFromPrivateKeyHex(privateKeyHex string) (*ecdsa.PrivateKey, *ecdsa.PublicKey, common.Address, error) {

	log.Println("Generating the keypair...")

	// If hex string has "0x" at the start discard it
	if privateKeyHex[:2] == "0x" {
		privateKeyHex = privateKeyHex[2:]
	}

	// Convert hex key to a private key object
	privateKey, privateKeyErr := crypto.HexToECDSA(privateKeyHex)

	if privateKeyErr != nil {
		return nil, nil, common.Address{}, privateKeyErr
	}

	// Generate the public key using the private key
	publicKey := privateKey.Public()

	// Cast crypto.Publickey object to the ecdsa.PublicKey object
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		return nil, nil, common.Address{}, errors.New("error casting public key to ECDSA")
	}

	// Convert publickey to a address
	pubkeyAsAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return privateKey, publicKeyECDSA, pubkeyAsAddress, nil
}

func GetAddressFromPrivateKey(privateKey string) (common.Address, error) {
	privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return common.Address{}, err
	}

	publicKey := privateKeyECDSA.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, fmt.Errorf("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return fromAddress, nil
}

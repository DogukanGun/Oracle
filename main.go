package main

import (
	"Oracle/contracts"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	_ "log"
	"strings"
)

func main() {
	contractAddress := "0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512"
	lendingPoolAddress := "0x5FbDB2315678afecb367f032d93F642f64180aa3"
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
	privateKey := "0xde9be858da4a475276426320d5e9262ecfc3ba460bfac56360bfa6c4c28b4ee0"
	fromAddress, err := GetAddressFromPrivateKey(privateKey)
	if err != nil {
		// Handle error
	}
	signer, err := getSigner(privateKey)
	if err != nil {
		// Handle error
	}

	resultOfSet, err := contractInstance.SetLendingPoolAddress(&bind.TransactOpts{
		From:   fromAddress,
		Value:  nil,
		Signer: signer,
	}, common.HexToAddress(lendingPoolAddress))
	if err != nil {
		// Handle error
		fmt.Println(err)
	}
	fmt.Println(resultOfSet)
	result, err := contractInstance.GetLendingPoolAddress(&bind.CallOpts{
		From: fromAddress,
	})
	if err != nil {
		// Handle error
	}
	fmt.Println(result)
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

func getSigner(privateKeyStr string) (bind.SignerFn, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return nil, err
	}

	signer := bind.NewKeyedTransactor(privateKey)
	return signer.Signer, nil
}

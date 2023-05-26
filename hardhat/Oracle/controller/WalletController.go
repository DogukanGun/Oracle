package controller

import (
	"Oracle/contractcaller"
	"Oracle/data"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WalletController struct{}

func (h *WalletController) CreateUserAddress(c *gin.Context) {
	var requestBody data.CreateUserAddress
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	walletAddress := "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"
	userHash := requestBody.UserHash
	username := requestBody.Username
	password := requestBody.Password
	err, contractInstance, _, res, client, _ := contractcaller.CreateFunctionRequirementsForWallet(walletAddress, "0x47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a")
	resultOfSet, err := contractInstance.CreateUserAddress(res, userHash, password, username)
	receipt, err := client.TransactionReceipt(context.Background(), resultOfSet.Hash())
	fmt.Println(receipt.Status)
	fmt.Println(receipt.Logs[0].Topics)
	fmt.Println(common.BytesToAddress(receipt.Logs[0].Data))
	if err != nil {
		// Handle the error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(resultOfSet.Data())

	c.JSON(200, gin.H{"message": fmt.Sprintf("Res =")})
}

func (h *WalletController) GetUserAssets(c *gin.Context) {
	userAddress := c.Param("userAddress")
	walletAddress := "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"
	err, contractInstance, _publicAddress, _, _, _ := contractcaller.CreateFunctionRequirementsForWallet(walletAddress, "0x47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a")
	resultOfSet, err := contractInstance.GetUserAssets(&bind.CallOpts{
		From:    _publicAddress,
		Context: context.Background(),
	}, common.HexToAddress(userAddress))
	if err != nil {
		// Handle error
		fmt.Println(err)
	}
	fmt.Println(resultOfSet)

	c.JSON(200, gin.H{"message": fmt.Sprintf("Res = %s", resultOfSet)})
}

func (h *WalletController) AddNewAsset(c *gin.Context) {

}

func (h *WalletController) Deposit(c *gin.Context) {

}

func (h *WalletController) Withdraw(c *gin.Context) {

}

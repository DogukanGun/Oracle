package controller

import (
	"Oracle/contractcaller"
	"Oracle/data"
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"

	// Import the Gin library
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// HelloWorldController will hold the methods to the
type OracleController struct{}

// Default controller handles returning the hello world JSON response
func (h *OracleController) SetLendingAddress(c *gin.Context) {
	var requestBody data.SetLendingAddress
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	lendingPoolAddress := requestBody.LendingPoolAddress
	oracleAddress := requestBody.OracleAddress
	fmt.Println(requestBody)
	err, contractInstance, _, res := contractcaller.CreateFunctionRequirementsForOracle(oracleAddress, "0x47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a")
	resultOfSet, err := contractInstance.SetLendingPoolAddress(res, common.HexToAddress(lendingPoolAddress))
	if err != nil {
		// Handle error
		fmt.Println(err)
	}
	fmt.Println(resultOfSet)

	c.JSON(200, gin.H{"message": fmt.Sprintf("Res = %s", resultOfSet)})
}

func (h *OracleController) GetLendingAddress(c *gin.Context) {
	oracleAddress := c.Param("oracleAddress")
	err, contractInstance, _publicAddress, _ := contractcaller.CreateFunctionRequirementsForOracle(oracleAddress, "0x47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a")
	result, err := contractInstance.GetLendingPoolAddress(&bind.CallOpts{
		From:    _publicAddress,
		Context: context.Background(),
	})
	if err != nil {
		// Handle error
		fmt.Println(err)
	}
	fmt.Println(result)
	c.JSON(200, fmt.Sprintf("Res = %s", result))
}

func (h *OracleController) SetInterestRate(c *gin.Context) {
	var requestBody data.SetInterestRate
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	oracleAddress := requestBody.OracleAddress
	var interestRateNumber big.Int
	errForNumberConv, ok := interestRateNumber.SetString(requestBody.InterestRate, 10)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": errForNumberConv})
		return
	}
	_, contractInstance, _, res := contractcaller.CreateFunctionRequirementsForOracle(oracleAddress, "0x47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a")
	result, _ := contractInstance.SetInterestRate(res, &interestRateNumber)
	// Access the request body fields
	fmt.Println(requestBody)
	c.JSON(200, gin.H{"message": fmt.Sprintf("The result is %s", result)})
}

func (h *OracleController) IncreaseTotalBarrowedLimitOf(c *gin.Context) {
	userAddress := c.Param("userId")
	var requestBody data.SetInterestRate
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	oracleAddress := requestBody.OracleAddress
	var interestRateNumber big.Int
	errForNumberConv, ok := interestRateNumber.SetString(requestBody.InterestRate, 10)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": errForNumberConv})
		return
	}
	_, contractInstance, _, res := contractcaller.CreateFunctionRequirementsForOracle(oracleAddress, "0x47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a")
	result, _ := contractInstance.IncreaseTotalBarrowedLimitOf(res, common.HexToAddress(userAddress), &interestRateNumber)
	c.JSON(200, gin.H{"message": fmt.Sprintf("The result is %s", result)})
}

func (h *OracleController) DecreaseTotalBarrowedLimitOf(c *gin.Context) {
	userAddress := c.Param("userId")
	var requestBody data.SetInterestRate
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	oracleAddress := requestBody.OracleAddress
	var interestRateNumber big.Int
	errForNumberConv, ok := interestRateNumber.SetString(requestBody.InterestRate, 10)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": errForNumberConv})
		return
	}
	_, contractInstance, _, res := contractcaller.CreateFunctionRequirementsForOracle(oracleAddress, "0x47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a")
	result, _ := contractInstance.DecreaseTotalBarrowedLimitOf(res, common.HexToAddress(userAddress), &interestRateNumber)
	c.JSON(200, gin.H{"message": fmt.Sprintf("The result is %s", result)})
}

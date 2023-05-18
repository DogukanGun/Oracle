package controller

import (
	"Oracle/data"
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

	// Access the request body fields
	param1 := requestBody.LendingPoolAddress
	fmt.Println(requestBody)
	c.JSON(200, gin.H{"message": fmt.Sprintf("The address is %s", param1)})
}

func (h *OracleController) GetLendingAddress(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello world, climate change is real"})
}

func (h *OracleController) SetInterestRate(c *gin.Context) {
	var requestBody data.SetInterestRate
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Access the request body fields
	param1 := requestBody.InterestRate
	fmt.Println(requestBody)
	c.JSON(200, gin.H{"message": fmt.Sprintf("The address is %s", param1)})
}

func (h *OracleController) IncreaseTotalBarrowedLimitOf(c *gin.Context) {
	id := c.Param("userId")
	c.JSON(200, gin.H{"message": fmt.Sprintf("The id of param is %s", id)})
}

func (h *OracleController) DecreaseTotalBarrowedLimitOf(c *gin.Context) {
	id := c.Param("userId")
	c.JSON(200, gin.H{"message": fmt.Sprintf("The id of param is %s", id)})
}

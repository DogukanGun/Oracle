package main

import (
	"Oracle/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	v1 := server.Group("/api/v1")
	{
		// Define the hello controller
		hello := new(controller.OracleController)
		// Define a GET request to call the Default
		// method in controllers/hello.go
		v1.POST("/lending/address", hello.SetLendingAddress)
		v1.GET("/lending/address", hello.GetLendingAddress)
		v1.POST("/interest-rate", hello.SetInterestRate)
		v1.POST("/increase/barrowed-limit/:userId", hello.IncreaseTotalBarrowedLimitOf)
		v1.POST("/decrease/barrowed-limit/:userId", hello.DecreaseTotalBarrowedLimitOf)

	}

	// Handle error response when a route is not defined
	server.NoRoute(func(c *gin.Context) {
		// In gin this is how you return a JSON response
		c.JSON(404, gin.H{"message": "Not found"})
	})
	server.Run()
}

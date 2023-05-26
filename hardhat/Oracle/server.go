package main

import (
	"Oracle/controller"
	"Oracle/subscriber"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	v1 := server.Group("/api/v1")
	{
		oracle := v1.Group("/oracle")
		{
			oracleController := new(controller.OracleController)
			oracle.POST("/lending/address", oracleController.SetLendingAddress)
			oracle.GET("/lending/address/:oracleAddress", oracleController.GetLendingAddress)
			oracle.POST("/interest-rate", oracleController.SetInterestRate)
			oracle.POST("/increase/barrowed-limit/:userId", oracleController.IncreaseTotalBarrowedLimitOf)
			oracle.POST("/decrease/barrowed-limit/:userId", oracleController.DecreaseTotalBarrowedLimitOf)
		}

		wallet := v1.Group("/wallet")
		{
			walletController := new(controller.WalletController)
			wallet.POST("/new", walletController.CreateUserAddress)
			wallet.GET("/:userAddress", walletController.GetUserAssets)
			wallet.POST("/new/asset", walletController.AddNewAsset)
			wallet.POST("/deposit", walletController.Deposit)
			wallet.POST("/withdraw", walletController.Withdraw)
		}

		networkListener := v1.Group("/listener")
		{
			networkListenerController := new(subscriber.NetworkListenerController)
			networkListener.POST("/start", networkListenerController.ListenNetwork)
		}
		//TODO Lending pool routes
	}

	// Handle error response when a route is not defined
	server.NoRoute(func(c *gin.Context) {
		// In gin this is how you return a JSON response
		c.JSON(404, gin.H{"message": "Not found"})
	})
	server.Run()
}

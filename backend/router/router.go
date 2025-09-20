package router

import (
	"blockchain/controller"
	"blockchain/global"
	"blockchain/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitRouter() {

	router := gin.Default()
	router.Use(middleware.Cors())

	RegistRouter(router)

	port := viper.GetString("server.port")
	if port == "" {
		port = "8080"
	}
	router.Run(port)

	global.Logger.Info("Start Server")

}

func RegistRouter(r *gin.Engine) {

	api := r.Group("/api")
	api.POST("/register", controller.Register)
	api.POST("/login", controller.Login)

	auth := api.Group("/auth", middleware.Auth())
	{

		user := auth.Group("/user")
		{
			user.POST("/profile", controller.Profile)
			user.POST("/mine", controller.Mine)
		}

		card := auth.Group("/card")
		{
			card.POST("/mint", controller.Mint)
			card.POST("/destroy", controller.Destroy)
			card.POST("/query", controller.Query)
			card.POST("/sell", controller.Sell)
			card.POST("/cancel", controller.Cancel)
		}

		market := auth.Group("/market")
		{
			market.POST("/query", controller.Market)
			market.POST("/buy", controller.Buy)
		}

	}

}

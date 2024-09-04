package route

import (
	controllers "github.com/msazad/assessment/src/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	/***BASEPATH OF AN API. NOTE:THIS SHOULDN'T BE CHANGED***/
	api := router.Group("/MocktestAPI")

	/***ADD THE ROUTES HERE***/
	api.POST("/saveUsers", controllers.CreateUsers)
	api.GET("/readUsers", controllers.ReadUsers)
	api.PUT("/updateUsers", controllers.UpdateUsers)
	api.DELETE("/deleteUsers", controllers.DeleteUsers)
	api.GET("/cryptocurrencies", controllers.ListCryptocurrencies)
	api.GET("/cryptocurrencies/search", controllers.SearchCryptocurrency)

}

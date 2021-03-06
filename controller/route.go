package controller

import (
	"github.com/expenseledger/web-service/config"
	"github.com/gin-gonic/gin"
)

// Route data structure to hold a path and the corresponding handler
type Route struct {
	Path    string
	Handler func(context *gin.Context)
}

// InitRoutes ...
func InitRoutes() *gin.Engine {
	configs := config.GetConfigs()
	router := gin.Default()

	walletRoute := router.Group("/wallet")
	walletRoute.POST("/create", walletCreate)
	walletRoute.POST("/get", walletGet)
	walletRoute.POST("/delete", walletDelete)
	walletRoute.POST("/list", walletList)
	walletRoute.POST("/listTypes", walletListTypes)
	walletRoute.POST("/init", walletInit)

	categoryRoute := router.Group("/category")
	categoryRoute.POST("/create", categoryCreate)
	categoryRoute.POST("/get", categoryGet)
	categoryRoute.POST("/delete", categoryDelete)
	categoryRoute.POST("/list", categoryList)
	categoryRoute.POST("/init", categoryInit)

	if configs.Mode != "PRODUCTION" {
		walletRoute.POST("/clear", walletClear)
		categoryRoute.POST("/clear", categoryClear)
	}

	return router
}

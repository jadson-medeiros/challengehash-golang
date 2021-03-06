package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jadson-medeiros/challengehash/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {

	main := router.Group("api/v1")
	{
		products := main.Group("products")
		{
			products.GET("/", controllers.ShowProducts)
			products.GET("/:id", controllers.ShowProduct)
			products.POST("/", controllers.CreateProduct)
		}
	}

	return router
}

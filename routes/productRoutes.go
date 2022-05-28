package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/musa/product-api/controllers"
	"github.com/musa/product-api/middleware"
)

func ProductRoutes(routes *gin.Engine) {
	routes.Use(middleware.Authorization())
	routes.GET("/products", controllers.GetProducts())
	routes.GET("/product/:id", controllers.GetProduct())
	routes.DELETE("/product/:id", controllers.DeleteProduct())
	routes.POST("/product", controllers.CreateProduct())
	routes.PUT("/product/:id", controllers.UpdateProduct())
}

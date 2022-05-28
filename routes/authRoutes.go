package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/musa/product-api/controllers"
)

func AuthRoutes(routes *gin.Engine) {
	routes.POST("/auth/register", controllers.Register())
	routes.POST("/auth/login", controllers.Login())
}

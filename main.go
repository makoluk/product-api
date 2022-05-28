package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/musa/product-api/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9001"
	}

	router := gin.New()

	routes.AuthRoutes(router)
	routes.ProductRoutes(router)

	router.Run(":" + port)
}

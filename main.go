package main

import (
	routes "github.com/LeonardoMagnani/auth_go/routes"

	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRouter(router)
	routes.AuthRouter(router)

	router.Run(":" + port)
}

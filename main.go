package main

import (
	"log"

	routes "github.com/LeonardoMagnani/auth_go/routes"

	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	log.SetPrefix("[Auth_Go] ")

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())

	// No-middleware routes
	routes.AuthRouter(router)

	// Middleware routes
	routes.UserRouter(router)

	router.Run(":" + port)
}

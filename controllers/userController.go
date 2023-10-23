package controllers

import (
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Listagem de usuários",
	})
}

func GetUser(c *gin.Context) {
	userID := c.Param("id")
	c.JSON(200, gin.H{
		"message": "Detalhes do usuário",
		"user_id": userID,
	})
}

func Signup(c *gin.Context) {
	// fmt.Println(database.DBInstance())
	c.JSON(200, gin.H{
		"message": "Cadastro de usuário",
	})
}

func Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Login de usuário",
	})
}

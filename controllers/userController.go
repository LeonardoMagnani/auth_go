package controllers

import (
	"database/sql"

	"github.com/LeonardoMagnani/auth_go/database"
	"github.com/LeonardoMagnani/auth_go/helpers"
	"github.com/LeonardoMagnani/auth_go/models"
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

func VerifyUser(c *gin.Context) {
	var verifyUserRequest models.VerifyUserRequest

	if err := c.ShouldBindJSON(&verifyUserRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	hash := helpers.Hash(verifyUserRequest.Code)

	c.JSON(200, gin.H{
		"message": hash,
	})
}

func Signup(c *gin.Context) {
	dbInstance := database.DBInstance()

	var signupRequest models.SignupRequest

	if err := c.ShouldBindJSON(&signupRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var count int

	errSelect := dbInstance.QueryRow("SELECT COUNT(ID) FROM users WHERE Email = ?", signupRequest.Email).Scan(&count)

	if errSelect != nil {
		c.JSON(500, gin.H{"error": errSelect.Error()})
		return
	}

	if count != 0 {
		c.JSON(400, gin.H{"error": "E-mail já cadastrado."})
		return
	}

	hashUserId, _ := helpers.GenerateHash()

	_, err := dbInstance.Exec("INSERT INTO users (First_name, Last_name, Password, Email, Phone, User_id) VALUES (?, ?, ?, ?, ?, ?)",
		signupRequest.FirstName, signupRequest.LastName, helpers.HashFromString(signupRequest.Password), signupRequest.Email, signupRequest.Phone, hashUserId)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	hash, randomNumber := helpers.GenerateHash()

	_, erro := dbInstance.Exec("INSERT INTO verify_token (Token, User_id) VALUES (?, ?)",
		hash, hashUserId)

	if erro != nil {
		c.JSON(500, gin.H{"error": erro.Error()})
		return
	}

	helpers.SendMail(randomNumber, signupRequest.Email)

	defer dbInstance.Close()

	c.JSON(200, gin.H{
		"user":  signupRequest.FirstName + " " + signupRequest.LastName,
		"email": signupRequest.Email,
		"phone": signupRequest.Phone,
	})
}

func Login(c *gin.Context) {
	var loginRequest models.LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	dbInstance := database.DBInstance()

	var result models.AuthJwtModel

	err := dbInstance.QueryRow("SELECT First_name, Password, Email, Phone, User_type, User_id FROM users WHERE Email = ?", loginRequest.Email).Scan(
		&result.First_name,
		&result.Password,
		&result.Email,
		&result.Phone,
		&result.User_type,
		&result.User_id,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(401, gin.H{"error": "Credenciais inválidas"})
			return
		}

		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	hashPassword := helpers.HashFromString(loginRequest.Password)

	if hashPassword != result.Password {
		c.JSON(401, gin.H{"error": "Credenciais inválidas"})
		return
	}

	accessToken, err := helpers.GenerateAccessToken(result)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	refreshToken, err := helpers.GenerateRefreshToken(result)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	c.JSON(200, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

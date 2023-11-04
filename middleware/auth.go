package middleware

import (
	"log"
	"os"
	"time"

	"github.com/LeonardoMagnani/auth_go/database"
	"github.com/LeonardoMagnani/auth_go/helpers"
	"github.com/LeonardoMagnani/auth_go/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")

		if authorization == "" {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		if !isValidTokenSignature(authorization) || !isValidTokenGeneral(authorization) {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func isValidTokenSignature(tokenString string) bool {
	secret := []byte(os.Getenv("SECRET_TOKEN"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if !token.Valid || err != nil {
		log.Println("isValidTokenSignature: !token.Valid || err != nil")

		return false
	}

	return true
}

func isValidTokenGeneral(tokenString string) bool {
	dbInstance := database.DBInstance()

	var result models.TokenEntity

	err := dbInstance.QueryRow("SELECT * FROM access_tokens WHERE Token = ?", tokenString).Scan(
		&result.Id,
		&result.Token,
		&result.User_id,
		&result.Expiration_time,
		&result.Created_at,
		&result.Last_used_at,
		&result.Canceled,
		&result.Canceled_at,
	)

	if err != nil {
		log.Println("isValidTokenGeneral: err != nil")
		log.Println(err)

		return false
	}

	if result.Canceled {
		log.Println("isValidTokenGeneral: result.Canceled")

		return false
	}

	expirationTime, _ := helpers.Timestamp(result.Expiration_time)
	location, _ := time.LoadLocation("America/Sao_Paulo")

	timeNow := time.Now().In(location).Truncate(time.Second)

	if expirationTime.Before(timeNow) {
		log.Println("isValidTokenGeneral: expirationTime.After(time.Now())")

		return false
	}

	return true
}

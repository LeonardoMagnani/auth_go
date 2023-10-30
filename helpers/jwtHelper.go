package helpers

import (
	"os"
	"time"

	"github.com/LeonardoMagnani/auth_go/database"
	"github.com/LeonardoMagnani/auth_go/models"
	"github.com/dgrijalva/jwt-go"
)

func timeLocation() (*time.Location, error) {
	timeLocation, err := time.LoadLocation(os.Getenv("TIME_LOCATION"))
	if err != nil {
		return nil, err
	}
	return timeLocation, nil
}

func GenerateAccessToken(user models.AuthJwtModel) (string, error) {
	location, err := timeLocation()
	if err != nil {
		return "", err
	}

	var secretKey = []byte(os.Getenv("SECRET_TOKEN"))
	var expTime = time.Now().In(location).Add(time.Hour * 1).Unix()

	claims := jwt.MapClaims{
		"user_id":    user.User_id,
		"first_name": user.First_name,
		"email":      user.Email,
		"phone":      user.Phone,
		"user_type":  user.User_type,
		"verified":   user.Verified,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	err = SaveToken(user.User_id, signedToken, expTime)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func GenerateRefreshToken(user models.AuthJwtModel) (string, error) {
	location, err := timeLocation()
	if err != nil {
		return "", err
	}

	var secretKey = []byte(os.Getenv("SECRET_TOKEN"))
	var expTime = time.Now().In(location).Add(time.Hour * 24 * 7).Unix()

	claims := jwt.MapClaims{
		"user_id": user.User_id,
		"exp":     expTime,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	err = SaveRefreshToken(user.User_id, signedToken, expTime)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func SaveToken(userID, accessToken string, expTime int64) error {
	dbInstance := database.DBInstance()
	time := time.Unix(expTime, 0).String()

	_, err := dbInstance.Exec("INSERT INTO access_tokens (User_id, Token, Expiration_time) VALUES (?, ?, ?)", userID, accessToken, time[:19])
	if err != nil {
		return err
	}

	return nil
}

func SaveRefreshToken(userID, accessToken string, expTime int64) error {
	dbInstance := database.DBInstance()
	time := time.Unix(expTime, 0).String()

	_, err := dbInstance.Exec("INSERT INTO refresh_tokens (User_id, Token, Expiration_time) VALUES (?, ?, ?)", userID, accessToken, time[:19])
	if err != nil {
		return err
	}

	return nil
}

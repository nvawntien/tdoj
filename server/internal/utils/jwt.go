package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SECRET_KEY = [2][]byte{[]byte(os.Getenv("ACCESS_JWT_SECRET")), []byte(os.Getenv("REFRESH_JWT_SECRET"))}


func GenerateToken(UserID string, index int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": UserID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte(SECRET_KEY[index]))

	if err != nil {
		return "", err
	}

	return result, nil
}

package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SECRET_KEY = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(UserID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": UserID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte(SECRET_KEY))

	if err != nil {
		return "", err
	}

	return result, nil
}

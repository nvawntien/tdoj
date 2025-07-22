package middleware

import (
	"backend/internal/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr, err := c.Cookie("access_token")

		if err != nil || tokenStr == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.Response{
				Status:  http.StatusUnauthorized,
				Message: "Missing access token",
			})
			return
		}

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(utils.SECRET_KEY[0]), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.Response{
				Status:  http.StatusUnauthorized,
				Message: "Invalid or expire token",
			})
			return
		}

		userID, ok := claims["user_id"].(string)

		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.Response{
				Status:  http.StatusUnauthorized,
				Message: "Cannot get user id from claims",
			})
			return
		}

		c.Set("user_id", userID)
		c.Next()
	}
}

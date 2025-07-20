package routes

import (
	"backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetUpUserRouter(route *gin.RouterGroup, userHandler *handlers.UserHandler) {
	api := route.Group("/user")
	{
		api.GET("/", userHandler.Welcome)
		api.POST("/register", userHandler.Register)
	}
}

package routes

import (
	"backend/internal/handlers"
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetUpUserRouter(route *gin.RouterGroup, userHandler *handlers.UserHandler) {
	api := route.Group("/user")
	{
		api.GET("/", userHandler.Welcome)
		api.POST("/register", userHandler.Register)
		api.POST("/login", userHandler.Login)
		api.GET("/logout", userHandler.Logout)
		api.GET("/profile", middleware.AuthMiddleWare(), userHandler.GetProfile)
		//api.POST("/forgot-password", userHandler.ForgotPassword)
	}
}

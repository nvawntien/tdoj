package handlers

import (
	services "backend/internal/services/user"
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc services.UserService
}

func NewUserHandler(svc services.UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

func (h *UserHandler) Welcome(c *gin.Context) {
	fmt.Println("Hello world")
}

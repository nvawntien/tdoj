package handlers

import (
	customErrors "backend/internal/errors"
	"backend/internal/request"
	services "backend/internal/services/user"
	"backend/internal/utils"
	"fmt"
	"net/http"

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

func (h *UserHandler) Register(c *gin.Context) {
	var req request.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	ctx := c.Request.Context()

	user, err := h.svc.Register(ctx, req)

	if err != nil {
		if err == customErrors.UserConflict {
			c.JSON(http.StatusConflict, utils.Response{
				Status:  http.StatusConflict,
				Message: err.Error(),
				Data:    nil,
			})
		} else {
			c.JSON(http.StatusInternalServerError, utils.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
				Data:    nil,
			})
		}

		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Status:  http.StatusOK,
		Message: "Account registration successful, please check email to activate account",
		Data:    user,
	})
}

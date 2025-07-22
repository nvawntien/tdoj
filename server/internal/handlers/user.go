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

func (h *UserHandler) Login(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	ctx := c.Request.Context()

	accessToken, refreshToken, err := h.svc.Login(ctx, req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.SetCookie("access_token", accessToken, 900, "/", "", false, true)
	c.SetCookie("refresh_token", refreshToken, 3600*2, "/tdoj/user/refresh-token", "localhost", false, true)

	c.JSON(http.StatusOK, utils.Response{
		Status:  http.StatusOK,
		Message: "Login succesfully",
		Data:    nil,
	})
}

func (h *UserHandler) Logout(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "", false, true)

	c.JSON(http.StatusOK, utils.Response{
		Status:  http.StatusOK,
		Message: "Logout succesfully",
		Data:    nil,
	})
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, utils.Response{
			Status:  http.StatusUnauthorized,
			Message: "Cannot get user id from context",
			Data:    nil,
		})
		return
	}

	ctx := c.Request.Context()

	user, err := h.svc.GetProfile(ctx, userID.(string))

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Status:  http.StatusOK,
		Message: "Get your profile succesfully",
		Data:    user,
	})
}

package services

import (
	"backend/internal/models"
	"backend/internal/request"
	"context"
)

type UserService interface {
	Register(ctx context.Context, req request.RegisterRequest) (*models.User, error)
	Login(ctx context.Context, req request.LoginRequest) (string, string, error)
	GetProfile(ctx context.Context, userID string) (*models.User, error)
}

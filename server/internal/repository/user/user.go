package repository

import (
	"backend/internal/models"
	"context"
)

type UserRepository interface {
	CheckExistsUserByEmail(ctx context.Context, email string) (bool, error)
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByID(ctx context.Context, ID string) (*models.User, error)
}

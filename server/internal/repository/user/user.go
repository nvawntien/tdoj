package repository

import (
	"backend/internal/models"
	"context"
)

type UserRepository interface {
	CheckExistsUserByEmail(ctx context.Context, email string) (bool, error)
	CreateUser(ctx context.Context, user *models.User) error
}

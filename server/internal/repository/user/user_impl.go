package repository

import (
	customErrors "backend/internal/errors"
	"backend/internal/models"
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type userRepositoryImpl struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

func (r *userRepositoryImpl) CheckExistsUserByEmail(ctx context.Context, email string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`
	err := r.db.GetContext(ctx, &exists, query, email)
	return exists, err
}

func (r *userRepositoryImpl) CreateUser(ctx context.Context, user *models.User) error {
	query := `
		INSERT INTO users (user_id, full_name, email, username, password, rating, created_at, updated_at)
		VALUES (:user_id, :full_name, :email, :username, :password, :rating, :created_at, :updated_at)
	`

	_, err := r.db.NamedExecContext(ctx, query, user)

	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return customErrors.UserConflict
			}
		}	

		return err
	}

	return nil
}

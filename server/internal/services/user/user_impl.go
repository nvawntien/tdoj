package services

import (
	customErrors "backend/internal/errors"
	"backend/internal/models"
	repository "backend/internal/repository/user"
	"backend/internal/request"
	"backend/internal/utils"
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

type userServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userServiceImpl{
		userRepo: userRepo,
	}
}

func (s *userServiceImpl) Register(ctx context.Context, req request.RegisterRequest) (*models.User, error) {
	exists, err := s.userRepo.CheckExistsUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("Check exists user by email failed")
	}

	if exists {
		return nil, customErrors.UserConflict
	}

	hashedPassword, err := utils.HashAndSalt([]byte(req.Password))

	if err != nil {
		return nil, errors.New("Hash password failed")
	}

	user := &models.User{
		UserID:    uuid.NewString(),
		FullName:  req.FullName,
		Email:     req.Email,
		Username:  req.Email,
		Password:  hashedPassword,
		Rating:    0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.userRepo.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userServiceImpl) Login(ctx context.Context, req request.LoginRequest) (string, string, error) {
	user, err := s.userRepo.GetUserByEmail(ctx, req.Email)

	if err != nil {
		return "", "", errors.New("Get user by email failed")
	}
	if user == nil {
		return "", "", customErrors.UserNotFound
	}

	if !utils.ComparePasswords(user.Password, []byte(req.Password)) {
		return "", "", customErrors.PasswordIncorrect
	}

	accessToken, err := utils.GenerateToken(user.UserID, 0)

	if err != nil {
		return "", "", customErrors.ErrGenerateToken
	}

	refreshToken, err := utils.GenerateToken(user.UserID, 1)

	if err != nil {
		return "", "", customErrors.ErrGenerateToken
	}

	return accessToken, refreshToken, nil
}

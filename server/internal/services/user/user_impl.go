package services

import (
	customErrors "backend/internal/errors"
	"backend/internal/models"
	otpRepo "backend/internal/repository/otp"
	userRepo "backend/internal/repository/user"
	"backend/internal/request"
	"backend/internal/utils"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type userServiceImpl struct {
	userRepo userRepo.UserRepository
	otpRepo  otpRepo.OTPRepository
}

func NewUserService(userRepo userRepo.UserRepository, otpRepo otpRepo.OTPRepository) UserService {
	return &userServiceImpl{
		userRepo: userRepo,
		otpRepo:  otpRepo,
	}
}

func (s *userServiceImpl) Register(ctx context.Context, req request.RegisterRequest) (*models.User, error) {
	exists, err := s.userRepo.CheckExistsUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, customErrors.ErrCheckExistsUserByEmail
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
		Username:  req.Username,
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

func (s *userServiceImpl) GetProfile(ctx context.Context, userID string) (*models.User, error) {
	user, err := s.userRepo.GetUserByID(ctx, userID)

	if err != nil {
		return nil, errors.New("Get user by id failed")
	}

	if user == nil {
		return nil, customErrors.UserNotFound
	}

	return user, nil
}

func (s *userServiceImpl) ForgotPassword(ctx context.Context, email string) error {
	exists, err := s.userRepo.CheckExistsUserByEmail(ctx, email)

	if err != nil {
		return customErrors.ErrCheckExistsUserByEmail
	}

	if !exists {
		return customErrors.UserNotFound
	}

	otp, err := utils.GenerateOTP()

	if err != nil {
		return customErrors.ErrGenerateOTP
	}

	if err := s.otpRepo.StoreOTPInRedis(ctx, email, otp); err != nil {
		return err
	}

	subject := "OTP to verify forgot password"
	body := fmt.Sprintf("Your OTP is: %s. OTP has expiry 5 minutes.", otp)

	if err := utils.SendOTPToVerifyForgotPassword(email, subject, body); err != nil {
		return errors.New("Send otp to verify forgot password failed")
	}

	return nil
}

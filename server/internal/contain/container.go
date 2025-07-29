package contain

import (
	"backend/internal/config"
	"backend/internal/database"
	"backend/internal/handlers"
	"backend/internal/redis"
	otpRepo "backend/internal/repository/otp"
	userRepo "backend/internal/repository/user"
	userService "backend/internal/services/user"
)

type Contain struct {
	UserHandler *handlers.UserHandler
}

func NewContainer(cfg *config.Config) (*Contain, error) {
	if err := database.NewConnection(&cfg.Database); err != nil {
		return nil, err
	}

	if err := redis.NewConnection(&cfg.Redis); err != nil {
		return nil, err
	}

	userRepo := userRepo.NewUserRepository(cfg.Database.Db)
	otpRepo := otpRepo.NewOTPRepository(cfg.Redis.Rd)

	userService := userService.NewUserService(userRepo, otpRepo)
	userHandler := handlers.NewUserHandler(userService)

	return &Contain{
		UserHandler: userHandler,
	}, nil
}

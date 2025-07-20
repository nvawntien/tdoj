package contain

import (
	"backend/internal/config"
	"backend/internal/db"
	"backend/internal/handlers"
	userRepo "backend/internal/repository/user"
	userService "backend/internal/services/user"
)

type Contain struct {
	UserHandler *handlers.UserHandler
}

func NewContainer(cfg *config.Config) (*Contain, error) {
	db, err := db.NewConnection(&cfg.Database)

	if err != nil {
		return nil, err
	}

	userRepo := userRepo.NewUserRepository(db)
	userService := userService.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	return &Contain{
		UserHandler: userHandler,
	}, nil
}

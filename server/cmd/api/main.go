package main

import (
	"backend/internal/config"
	"backend/internal/contain"
	"backend/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()

	if err != nil {
		log.Fatal("Unable to load configuration")
	}

	contain, err := contain.NewContainer(cfg)

	if err != nil {
		log.Fatal("Container initialization failed")
	}

	r := gin.Default()

	route := r.Group("/tdoj")
	routes.SetUpUserRouter(route, contain.UserHandler)

	log.Printf("Server is running at port: %s", cfg.Server.Port)
	if err := r.Run(":" + cfg.Server.Port); err != nil {
		log.Fatal("Error running server:", err)
	}
}

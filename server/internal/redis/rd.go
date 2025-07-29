package redis

import (
	"backend/internal/config"
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

func NewConnection(r *config.RedisConfig) error {
	address := fmt.Sprintf("%s:%s", r.Host, r.Port)

	db, err := strconv.Atoi(r.Db)
	if err != nil {
		return fmt.Errorf("invalid Redis DB number: %w", err)
	}

	opt := &redis.Options{
		Addr:     address,
		Password: r.Password,
		DB:       db,
	}

	client := redis.NewClient(opt)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := client.Ping(ctx).Result(); err != nil {
		log.Println("Redis connection failed")
		return err
	}

	log.Println("Connected to Redis:", address)

	r.Rd = client
	return nil
}

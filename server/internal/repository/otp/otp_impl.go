package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type OTPRepositoryImpl struct {
	rd *redis.Client
}

func NewOTPRepository(rd *redis.Client) OTPRepository {
	return &OTPRepositoryImpl{
		rd: rd,
	}
}

func (r *OTPRepositoryImpl) StoreOTPInRedis(ctx context.Context, userIdentifier string, otp string) error {
	const OTPExpiryTime = 5 * time.Minute
	otpKey := fmt.Sprintf("OTP:%s", userIdentifier)

	if err := r.rd.Set(ctx, otpKey, otp, OTPExpiryTime).Err(); err != nil {
		return errors.New("Store otp in redis failed")
	}

	return nil
}

func (r *OTPRepositoryImpl) CheckUserBlockStatus(ctx context.Context, userIdentifier string) error {
	blockKey := fmt.Sprintf("OTP_REQUEST_BLOCK:%s", userIdentifier)

	isBlocked, err := r.rd.Exists(ctx, blockKey).Result()

	if err != nil {
		return errors.New("Checking user block status failed")
	}

	if isBlocked == 1 {
		return errors.New("You are temporarily blocked due to multiple unverified OTP requests.")
	}

	return nil
}

func (r *OTPRepositoryImpl) CheckRateLimit(ctx context.Context, userIdentifier string) error {
	const MaxRequestsPerSecond = 1

	rateLimitKey := fmt.Sprintf("OTP_REQUEST_RATE:%s", userIdentifier)
	rate, err := r.rd.Incr(ctx, rateLimitKey).Result()

	if err != nil {
		return errors.New("Check rate limit failed.")
	}

	if rate > MaxRequestsPerSecond {
		return errors.New("User has exceeded the maximum requests per second")
	}

	r.rd.Expire(ctx, rateLimitKey, time.Second)
	return nil
}

func (r *OTPRepositoryImpl) IncrementOTPRequestCount(ctx context.Context, userIdentifier string) error {
	const (
		MaxUnverifiedOTP = 5
		OTPRequestBlockTime = 10 * time.Second
	)

	requestKey := fmt.Sprintf("OTP_REQUEST_COUNT:%s", userIdentifier)
	requests, err := r.rd.Incr(ctx, requestKey).Result()

	if err != nil {
		return errors.New("Incrementing OTP request count failed")
	}

	if requests >= MaxUnverifiedOTP {
		if err := r.rd.Set(ctx, requestKey, "0", OTPRequestBlockTime).Err(); err != nil {
			return errors.New("Resetting request count failed")
		}
		
		if err := r.rd.Set(ctx, fmt.Sprintf("OTP_REQUEST_BLOCK:%s", userIdentifier), "1", OTPRequestBlockTime).Err(); err != nil {
			return errors.New("Blocking user from incrementing OTP request count failed")
		}
	}

	return nil
}

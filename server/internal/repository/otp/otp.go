package repository

import "context"

type OTPRepository interface {
	StoreOTPInRedis(ctx context.Context, userIdentifier string, otp string) error 
	CheckUserBlockStatus(ctx context.Context, userIdentifier string) error
	CheckRateLimit(ctx context.Context, userIdentifier string) error
	IncrementOTPRequestCount(ctx context.Context, userIdentifier string) error
}
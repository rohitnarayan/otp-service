package service

import (
	"context"

	"github.com/rohitnarayan/otp-service/internal/store"
)

type OTPService interface {
	CreateOTP(ctx context.Context, userID string) (store.OTPModel, error)
}

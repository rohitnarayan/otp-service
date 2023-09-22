package service

import (
	"context"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/pkg/errors"

	"github.com/rohitnarayan/otp-service/internal/config"
	"github.com/rohitnarayan/otp-service/internal/store"
)

type otpService struct {
	store store.OTPStore
}

type OTPService interface {
	CreateOTP(ctx context.Context, userID string) (*store.OTPModel, error)
}

func NewOTPService(store store.OTPStore) OTPService {
	return &otpService{
		store: store,
	}
}

func (s *otpService) CreateOTP(ctx context.Context, userID string) (*store.OTPModel, error) {
	otpValue := generateFixedLengthOTP()
	otpModel := &store.OTPModel{
		Otp:       otpValue,
		UserID:    userID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Status:    "created",
	}

	if err := s.store.InsertOTP(ctx, otpModel); err != nil {
		return nil, errors.Wrap(err, "error inserting OTP in DB")
	}

	return otpModel, nil
}

func generateFixedLengthOTP() string {
	baseValue := int(math.Pow10(config.App.Server.OTPLength))
	divisor := baseValue / 10

	firstDigit := (rand.Intn(9) + 1) * divisor
	restDigits := rand.Intn(divisor)
	value := firstDigit + restDigits

	return strconv.Itoa(value)
}

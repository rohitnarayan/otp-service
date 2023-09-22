package store

import "context"

type OTPStore interface {
	InsertOTP(ctx context.Context, otpModel *OTPModel) error
}

package store

import "time"

type OTPModel struct {
	Otp       string    `db:"otp"`
	UserID    string    `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Status    string    `db:"status"`
}

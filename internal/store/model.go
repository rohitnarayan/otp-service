package store

import "time"

type OTPModel struct {
	Otp       string    `json:"otp" db:"otp"`
	UserID    string    `json:"user_id" db:"user_id"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
	Status    string    `json:"status,omitempty" db:"status"`
}

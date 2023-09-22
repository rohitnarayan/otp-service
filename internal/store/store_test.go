package store

import (
	"context"
	"log"
	"testing"
	"time"
	
	"github.com/stretchr/testify/assert"

	"github.com/rohitnarayan/otp-service/internal/config"
	"github.com/rohitnarayan/otp-service/internal/postgres"
)

func TestOTPStore(t *testing.T) {
	ctx := context.Background()
	config.InitTestConfig()

	db, err := postgres.NewDB(config.App.Database.Postgres)
	if err != nil {
		log.Fatalf("error creating the DB")
	}

	store := NewStore(db, db)

	otpModel := &OTPModel{
		Otp:       "2341",
		UserID:    "rohit",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Status:    "created",
	}

	err = store.InsertOTP(ctx, otpModel)
	assert.Equal(t, nil, err)
}

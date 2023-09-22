package store

import (
	"context"
	"github.com/pkg/errors"
	"github.com/rohitnarayan/otp-service/internal/config"

	"github.com/jmoiron/sqlx"

	"github.com/rohitnarayan/otp-service/internal/postgres"
)

var (
	insertOTPQuery = &postgres.Query{
		Name:  "insert otp",
		Query: "INSERT INTO otp(otp, user_id, created_at, updated_at, status) VALUES($1, $2, $3, $4, $5)",
	}
)

// select * from otp where userid=$1 AND status=created limit 1;
// update otp set status=expired/validated whre

type otpStore struct {
	readDB  *sqlx.DB
	writeDB *sqlx.DB
}

type OTPStore interface {
	InsertOTP(ctx context.Context, otpModel *OTPModel) error
}

func NewStore(readDB *sqlx.DB, writeDB *sqlx.DB) OTPStore {
	return &otpStore{
		readDB:  readDB,
		writeDB: writeDB,
	}
}

func (s *otpStore) InsertOTP(ctx context.Context, otpModel *OTPModel) error {
	insertOTPQuery.Args = getInsertOTPArgs(otpModel)
	if err := postgres.Exec(ctx, s.writeDB, config.App.Database.Postgres.WriteTimeout, insertOTPQuery); err != nil {
		return errors.Wrap(err, "[store.InsertOTP]")
	}

	return nil
}

func getInsertOTPArgs(otpModel *OTPModel) []interface{} {
	return []interface{}{otpModel.Otp, otpModel.UserID, otpModel.CreatedAt, otpModel.UpdatedAt, otpModel.Status}
}

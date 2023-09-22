package postgres

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/rohitnarayan/otp-service/internal/config"
)

type otpStore struct {
	Id        int       `db:"id"`
	Otp       int       `db:"otp"`
	UserID    string    `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Status    string    `db:"status"`
}

func TestDBMethods(t *testing.T) {
	ctx := context.Background()
	query := []*Query{
		{
			Query: "SELECT * FROM otp WHERE id=$1",
			Args:  []interface{}{1},
		},
		{
			Query: "INSERT INTO otp VALUES($1, $2, $3,$4)",
			Args:  []interface{}{4534, "User1", time.Now().UTC(), "created"},
		},
		{
			Query: "SELECT * FROM otp WHERE otp=$1",
			Args:  []interface{}{4534},
		},
		{
			Query: "DELETE FROM employee WHERE otp=$1",
			Args:  []interface{}{4534},
		},
	}

	config.InitTestConfig()
	db, _ := NewDB(config.App.Database.Postgres)

	var otps []*otpStore
	err := Get(ctx, db, config.App.Database.Postgres.ReadTimeout, &otps, query[0])
	assert.Equal(t, err, nil)
	assert.Equal(t, 4534, otps[0].Otp)

	err = Exec(ctx, db, config.App.Database.Postgres.WriteTimeout, query[3])
	assert.Equal(t, nil, err)

	err = Exec(ctx, db, config.App.Database.Postgres.WriteTimeout, query[1])
	assert.Equal(t, nil, err)

	err = Get(ctx, db, config.App.Database.Postgres.ReadTimeout, &otps, query[2])
	assert.Equal(t, err, nil)
	assert.Equal(t, 4534, otps[0].Otp)
}

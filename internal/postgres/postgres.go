package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"

	"github.com/rohitnarayan/otp-service/internal/config"
)

type Query struct {
	Name  string
	Query string
	Args  []interface{}
}

func NewDB(cfg *config.PostgresConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open(cfg.Driver, ConnectionURL(cfg))
	if err != nil {
		return nil, errors.Wrapf(err, "[postgres.NewDB] failed to initialize postgres")
	}

	if err = db.Ping(); err != nil {
		return nil, errors.Wrap(err, "[postgres.NewDB] failed to ping postgres")
	}

	db.SetMaxIdleConns(cfg.MaxIdleConnections)
	db.SetMaxOpenConns(cfg.MaxOpenConnections)

	return db, nil
}

func ConnectionURL(cfg *config.PostgresConfig) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DatabaseName)
}

func Get(ctx context.Context, readDB *sqlx.DB, timeout time.Duration, dest interface{}, q *Query) error {
	if len(q.Query) == 0 {
		return nil
	}

	return withTimeout(ctx, timeout, func(ctx context.Context) error {
		if err := readDB.SelectContext(ctx, dest, q.Query, q.Args...); err != nil {
			return errors.Wrap(err, "[postgres.Get] failed to get result")
		}
		return nil
	})
}

func Exec(ctx context.Context, writeDB *sqlx.DB, timeout time.Duration, q *Query) error {
	if len(q.Query) == 0 {
		return nil
	}

	return withTimeout(ctx, timeout, func(ctx context.Context) error {
		if _, err := writeDB.ExecContext(ctx, q.Query, q.Args...); err != nil {
			return errors.Wrap(err, "[postgres.Exec] failed to execute query")
		}
		return nil
	})
}

func withTimeout(ctx context.Context, timeout time.Duration, op func(ctx context.Context) error) error {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	return op(ctxWithTimeout)
}

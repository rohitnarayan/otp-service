package config

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	InitTestConfig()

	assert.Equal(t, App.Server.Name, "leaderboard")
	assert.Equal(t, App.Server.Port, 8085)
	assert.Equal(t, App.Database.Postgres.DatabaseName, "leaderboard_production")
	assert.Equal(t, App.Database.Postgres.Host, "localhost")
	assert.Equal(t, App.Database.Postgres.Port, 5432)
	assert.Equal(t, App.Database.Postgres.Driver, "postgres")
	assert.Equal(t, App.Database.Postgres.WriteTimeout, time.Millisecond*time.Duration(200))
	assert.Equal(t, App.Database.Postgres.ReadTimeout, time.Millisecond*time.Duration(50))
	assert.Equal(t, App.Database.Postgres.Username, "leaderboard")
	assert.Equal(t, App.Database.Postgres.MaxOpenConnections, 10)
}

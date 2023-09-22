package config

import (
	"time"
)

type Config struct {
	Server   *ServerConfig
	Database *DatabaseConfig
	Logger   *LoggerConfig
}

type LoggerConfig struct {
	Level  string
	Format string
}

type ServerConfig struct {
	Name         string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

type DatabaseConfig struct {
	Postgres *PostgresConfig
}

type PostgresConfig struct {
	connectionURL      string
	Host               string
	Port               int
	Driver             string
	DatabaseName       string
	Username           string
	Password           string
	MaxIdleConnections int
	MaxOpenConnections int
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
}

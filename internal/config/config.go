package config

import (
	"os"

	"github.com/spf13/viper"
)

var App *Config

const (
	environment = "ENVIRONMENT"
)

func Init() {
	if os.Getenv(environment) == "test" {
		viper.SetConfigName("test")
	} else {
		viper.SetConfigName("application")
	}

	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath("./../configs")
	viper.AddConfigPath("./../../configs")
	viper.AddConfigPath("./../../../configs")
	viper.AddConfigPath("./../../../../configs")
	viper.AddConfigPath("./../../../../../configs")

	_ = viper.ReadInConfig()
	viper.AutomaticEnv()

	App = &Config{
		Database: &DatabaseConfig{
			Postgres: &PostgresConfig{
				Host:               getStringOrPanic("DB_HOST"),
				Port:               getIntOrPanic("DB_PORT"),
				DatabaseName:       getStringOrPanic("DB_NAME"),
				Driver:             getStringOrPanic("DB_DRIVER"),
				Username:           getStringOrPanic("DB_USER"),
				Password:           getStringOrPanic("DB_PASSWORD"),
				ReadTimeout:        getDurationInMS("DB_READ_TIMEOUT_IN_MS"),
				WriteTimeout:       getDurationInMS("DB_WRITE_TIMEOUT_IN_MS"),
				MaxOpenConnections: getIntOrPanic("DB_MAX_OPEN_CONNECTIONS"),
				MaxIdleConnections: getIntOrPanic("DB_MAX_IDLE_CONNECTIONS"),
			},
		},
		Server: &ServerConfig{
			Name: getStringOrPanic("APP_NAME"),
			Port: getIntOrPanic("PORT"),
		},
		Logger: &LoggerConfig{
			Level:  getStringOrPanic("LOG_LEVEL"),
			Format: getStringOrPanic("LOG_FORMAT"),
		},
	}
}

func InitTestConfig() {
	_ = os.Setenv("ENVIRONMENT", "test")

	Init()
}

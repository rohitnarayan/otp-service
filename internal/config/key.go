package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

func getStringOrPanic(key string) string {
	mustHaveKey(key)
	return viper.GetString(key)
}

func getIntOrPanic(key string) int {
	mustHaveKey(key)
	return viper.GetInt(key)
}

func getDurationInMS(key string) time.Duration {
	return time.Millisecond * time.Duration(getIntOrPanic(key))
}

func mustHaveKey(key string) {
	if !viper.IsSet(key) {
		panic(fmt.Sprintf("key %s is not set", key))
	}
}

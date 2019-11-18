package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	EntityProcessDuration time.Duration
	EntityDeleteDuration  time.Duration
	WithMigrations        bool
	Port                  string
}

func New() *Config {
	return &Config{
		WithMigrations:        boolEnv("WITH_MIGRATIONS", false),
		EntityDeleteDuration:  minutesEnv("ENTITY_DELETE_DURATION", 1),
		EntityProcessDuration: minutesEnv("ENTITY_PROCESS_DURATION", 1),
		Port:                  stringEnv("PORT", ":80"),
	}
}

func boolEnv(key string, def bool) bool {
	str := os.Getenv(key)
	res, err := strconv.ParseBool(str)
	if err != nil {
		return def
	}
	return res
}

func stringEnv(key string, def string) string {
	str := os.Getenv(key)
	if str == "" {
		return def
	}
	return str
}

func minutesEnv(key string, def int) time.Duration {
	str := os.Getenv(key)
	res, err := strconv.Atoi(str)
	if err != nil {
		return time.Duration(def) * time.Minute
	}
	return time.Duration(res) * time.Minute
}

var Cfg = New()

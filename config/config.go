package config

import (
	"github.com/labstack/gommon/log"
	"os"
)

type AppConfig struct {
	Port      string
	JWTSecret string
}

var Config AppConfig

func SetConfig() {
	Config = AppConfig{
		Port:      getEnv("port", "5000"),
		JWTSecret: getEnv("jwtSecret", ""),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	} else if defaultVal == "" {
		log.Fatalf("environment variable %s cannot have a nil value", key)
	}
	return defaultVal
}

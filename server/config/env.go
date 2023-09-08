package config

import (
	"os"
)

var (
	HttpHost,
	DbHost,
	DbPort,
	HttpPort,
	DbUsername,
	DbPassword,
	DbName,
	RedisURL string
)

func LoadEnv() {
	HttpPort = getEnv("HTTP_PORT", "8000")
	HttpHost = getEnv("HTTP_Host", "0.0.0.0")
	DbHost = getEnv("DB_HOST", "localhost")
	DbPort = getEnv("DB_PORT", "5433")
	DbUsername = getEnv("DB_USER", "test")
	DbPassword = getEnv("DB_PASSWORD", "test")
	DbName = getEnv("DB_NAME", "postgres")
	RedisURL = getEnv("REDIS_URL", "redis://localhost:6379")
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		return defaultValue
	}

	return value
}

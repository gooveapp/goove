package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env              string
	Port             string
	LogLevel         string
	DiscogsToken     string
	DiscogsUserAgent string
	DatabasePath     string
}

func FromEnv() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		Env:          getEnv("ENV", "production"),
		Port:         getEnv("PORT", "3000"),
		LogLevel:     getEnv("LOG_LEVEL", "info"),
		DiscogsToken: mustGetEnv("DISCOGS_TOKEN"),
		DatabasePath: getEnv("DATABASE_PATH", "/data/goove.db"),
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

func mustGetEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("Missing required env var: %s", key)
	}
	return val
}

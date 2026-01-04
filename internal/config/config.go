package config

import (
	"fmt"
	"os"
)

type Config struct {
	AppEnv           string
	ServerPort       string
	BaseURL          string
	DBHost           string
	DBPort           string
	DBUser           string
	DBPassword       string
	DBName           string
	RedisAddr        string
	JWTPublicKeyPath string
}

func Load() *Config {
	return &Config{
		AppEnv:     getEnv("APP_ENV", "local"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
		BaseURL:    getEnv("BASE_URL", "http://localhost:8080"),

		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
		DBName:     getEnv("DB_NAME", "shortener"),

		RedisAddr:        getEnv("REDIS_ADDR", "localhost:6379"),
		JWTPublicKeyPath: getEnv("JWT_PUBLIC_KEY_PATH", "./keys/public.pem"),
	}
}

func getEnv(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return defaultValue
}

func (c *Config) DBURL() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.DBUser,
		c.DBPassword,
		c.DBHost,
		c.DBPort,
		c.DBName,
	)
}

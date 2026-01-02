package config

import "os"

type Config struct {
	AppEnv     string
	ServerPort string
}

func Load() *Config {
	return &Config{
		AppEnv:     getEnv("APP_ENV", "local"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return defaultValue
}

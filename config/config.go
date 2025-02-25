package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	SecretKey  string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load() // загрузка вирт окружения
	if err != nil {
		return nil, err
	}

	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5555"),
		DBUser:     getEnv("DB_USER", "ps"),
		DBPassword: getEnv("DB_PASSWORD", "ps"),
		DBName:     getEnv("DB_NAME", "ps"),
		DBSSLMode:  getEnv("DB_SSL_MODE", "enable"),
		SecretKey:  getEnv("SECRET_KEY", "default_secret_key"),
	}, nil
}

// НЕОБЯЗАТЕЛЬНО, НО УДОБНО
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

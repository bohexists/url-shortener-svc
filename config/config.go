package config

import (
	"os"
)

// Config структура для хранения конфигурации
type Config struct {
	MongoURI string
	Port     string
}

// LoadConfig загружает конфигурацию из переменных среды
func LoadConfig() *Config {
	return &Config{
		MongoURI: getEnv("MONGO_URI", "mongodb://localhost:27017"),
		Port:     getEnv("PORT", "8080"),
	}
}

// Вспомогательная функция для получения переменных среды
func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

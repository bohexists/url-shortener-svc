package config

import (
	"os"
)

// Config structure for the config file
type Config struct {
	MongoURI string
	Port     string
}

// LoadConfig loads environment variables from Docker environment
func LoadConfig() *Config {
	return &Config{
		MongoURI: getEnv("MONGO_URI", "mongodb://localhost:27017"),
		Port:     getEnv("PORT", "8080"),
	}
}

// Helper function to retrieve environment variables or set a default value
func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

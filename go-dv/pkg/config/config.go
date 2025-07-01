package config

import (
	"os"
	"strconv"
)

func New() *AppConfig {
	return &AppConfig{
		Server: ServerConfig{
			Port: getEnvAsInt("PORT", 9000),
			Host: getEnv("HOST", "0.0.0.0"),
		},
		Database: DatabaseConfig{
			DSN: getEnv("DATABASE_DSN", "database.sqlite"),
		},
		Auth: AuthConfig{
			JWTSecret:       getEnv("JWT_SECRET", "a-ver5-02ve0uze-s9fau89-0ec3ee"),
			TokenCookieName: "access_token",
		},
	}
}

/*
* getEnv
* reads an environment variable or returns a default value.
*/
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

/* 
* getEnvAsInt
* reads an environment variable as an integer or returns a default value.
*/
func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

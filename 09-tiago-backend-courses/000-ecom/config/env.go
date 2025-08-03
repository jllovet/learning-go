package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var InitializedConfig = InitConfig()

type Config struct {
	PublicHost             string
	Port                   string
	DBUser                 string
	DBPassword             string
	DBAddress              string
	DBName                 string
	JWTExpirationInSeconds int64
	JWTSecret              string
}

func InitConfig() Config {
	godotenv.Load()
	return Config{
		PublicHost:             getEnv("ECOM_PUBLIC_HOST", "http://localhost"),
		Port:                   getEnv("ECOM_PORT", "8080"),
		DBUser:                 getEnv("ECOM_DB_USER", "mysql"),
		DBPassword:             getEnv("ECOM_DB_PASSWORD", "password"),
		DBAddress:              fmt.Sprintf("%s:%s", getEnv("ECOM_DB_HOST", "127.0.0.1"), getEnv("ECOM_DB_PORT", "3306")),
		DBName:                 getEnv("ECOM_DB_NAME", "ecom"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXP", 6000),
		JWTSecret:              getEnv("JWT_SECRET", "this-only-looks-secret-this-only-looks-secret-this-only-looks-secret-this-only-looks-secret-this-only-looks-secret-this-only-looks-secret-this-only-looks-secret-this-only-looks-secret"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return defaultValue
		}
		return i
	}
	return defaultValue
}

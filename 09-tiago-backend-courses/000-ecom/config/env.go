package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var Envs = initConfig()

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
}

func initConfig() Config {
	godotenv.Load()
	return Config{
		PublicHost: getEnv("ECOM_PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("ECOM_PORT", "8080"),
		DBUser:     getEnv("ECOM_DB_USER", "mysql"),
		DBPassword: getEnv("ECOM_DB_PASSWORD", "password"),
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("ECOM_DB_HOST", "127.0.0.1"), getEnv("ECOM_DB_PORT", "3306")),
		DBName:     getEnv("ECOM_DB_NAME", "ecom"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

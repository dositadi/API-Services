package config

import (
	"os"
)

type Configuration struct {
	ServerPort string
	DBUser     string
	DBHost     string
	DBPassword string
	DBName     string
}

func GetConfiguration() Configuration {
	return Configuration{
		ServerPort: os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}
}

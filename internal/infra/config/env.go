package config

import (
	"os"
)

type API struct {
	Port string
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Db       string
}

type Config struct {
	API
	Database
}

func GetConfig() *Config {
	return &Config{
		API: API{
			Port: os.Getenv("API_PORT"),
		},
		Database: Database{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Db:       os.Getenv("DB_NAME"),
		},
	}
}

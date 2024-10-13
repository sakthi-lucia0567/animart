package config

import (
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	Port               string
	DBConnectionString string
}

func Load() (*Config, error) {
	err := godotenv.Load("./configs/.env")
	if err != nil {
		return nil, err
	}

	return &Config{
		Port:               os.Getenv("PORT"),
		DBConnectionString: os.Getenv("DB_CONNECTION_STRING"),
	}, nil
}
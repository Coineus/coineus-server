package config

import (
	"github.com/joho/godotenv"
)

func GetConfig() error {
	err := godotenv.Load("./internal/config/.env")
	return err
}

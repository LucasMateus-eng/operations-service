package config

import (
	"github.com/joho/godotenv"
)

func GetConfig(pathToConfigFile string) error {
	if err := godotenv.Load(pathToConfigFile); err != nil {
		return err
	}

	return nil
}

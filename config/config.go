package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppName        string `mapstructure:"APP_NAME"`
	AppEnv         string `mapstructure:"APP_ENV"`
	AppLogLevel    string `mapstructure:"APP_LOG_LEVEL"`
	AppDefaultPort string `mapstructure:"APP_DEFAULT_PORT"`
	DBHost         string `mapstructure:"DB_HOST"`
	DBPort         string `mapstructure:"DB_PORT"`
	DBUser         string `mapstructure:"DB_USER"`
	DBPass         string `mapstructure:"DB_PASS"`
	DBName         string `mapstructure:"DB_NAME"`
}

func NewConfig(configType, configName, configPath string) *Config {
	config := Config{}
	viper.SetConfigType(configType)
	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if config.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &config
}

package configuration

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type AppConfig struct {
	Port string
	DatabaseUrl string
}

func CreateConfig() AppConfig {
	viper.SetConfigName(getConfigName())
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config/")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	config := AppConfig{}
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl != "" {
		config.DatabaseUrl = databaseUrl + "?sslmode=disable"
	}

	port := os.Getenv("PORT")
	if port != "" {
		config.Port = port
	}

	return config
}

func getConfigName() string {
	environment := os.Getenv("profile")
	configName := "application"

	if environment != "" {
		return fmt.Sprintf("%s-%s", configName, environment)
	}

	return configName
}
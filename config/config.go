package config

import (
	"log"

	"github.com/spf13/viper"
)

// config ...
var config *viper.Viper

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init() {
	var err error
	config = viper.New()

	config.SetConfigName("app")
	config.SetConfigType("env")
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")

	config.AutomaticEnv()

	err = config.ReadInConfig()
	if err != nil {
		log.Printf(err.Error())
		log.Fatal("Error on parsing configuration file")
	}
}

// GetConfig return value of variable load from config file and environments
func GetConfig(s string) string {
	return config.GetString(s)
}

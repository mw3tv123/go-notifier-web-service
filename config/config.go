package config

import (
	"log"

	"github.com/spf13/viper"
)

// config ...
var config *viper.Viper

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init(env string) {
	var err error
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")
	err = config.ReadInConfig()
	if err != nil {
		log.Printf(err.Error())
		log.Fatal("Error on parsing configuration file")
	}
}

// GetConfig return an instance of config which has been loaded
func GetConfig() *viper.Viper {
	return config
}

package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

var (
	once sync.Once
)

// config ...
var config *viper.Viper

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init() {
	once.Do(func() {
		config = viper.New()

		config.SetConfigName("app")
		config.SetConfigType("env")
		config.AddConfigPath("../config/")
		config.AddConfigPath("config/")

		config.AutomaticEnv()

		err := config.ReadInConfig()
		if err != nil {
			log.Printf(err.Error())
			log.Fatal("Error on parsing configuration file")
		}
	})
}

// GetConfig return value of variable load from config file and environments
func GetConfig(s string) string {
	return config.GetString(s)
}

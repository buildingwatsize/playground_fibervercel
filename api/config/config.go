package config

import (
	"log"

	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath("./_env")
	viper.SetConfigName("env")
	viper.SetConfigType("json")

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicf("Fatal error config file: %s \n", err)
	}
}

type Configuration struct {
	FirstTimeMessage string
}

func NewConfiguration() Configuration {
	return Configuration{
		FirstTimeMessage: viper.GetString("FIRST_TIME_MESSAGE"),
	}
}

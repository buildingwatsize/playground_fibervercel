package handler

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

type configuration struct {
	firstTimeMessage string
}

func newConfiguration() configuration {
	return configuration{
		firstTimeMessage: viper.GetString("FIRST_TIME_MESSAGE"),
	}
}

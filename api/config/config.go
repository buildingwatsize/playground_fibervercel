package config

import (
	"log"
	"net/http"

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

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(nil))
}

type Configuration struct {
	FirstTimeMessage string
}

func NewConfiguration() Configuration {
	return Configuration{
		FirstTimeMessage: viper.GetString("FIRST_TIME_MESSAGE"),
	}
}

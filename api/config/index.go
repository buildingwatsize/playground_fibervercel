package config

import (
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func init() {
	log.Println("Init: setting up the configuration")
	viper.AddConfigPath("./_env")
	viper.SetConfigName("env")
	viper.SetConfigType("json")

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Error: on initializing config file: %s \n", err)
	} else {
		log.Println("Finished: the configuration is already use")
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/api", http.StatusSeeOther)
}

type Configuration struct {
	FirstTimeMessage string
	AnotherVariable  string
}

func NewConfiguration() Configuration {
	return Configuration{
		FirstTimeMessage: viper.GetString("FIRST_TIME_MESSAGE"),
		AnotherVariable:  viper.GetString("ANOTHER_VARIABLE"),
	}
}

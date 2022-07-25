package config

import (
	"log"
	"os/exec"

	"github.com/spf13/viper"
)

func init() {
	out, _ := exec.Command("ls", "-la").Output()
	log.Println(string(out))
	out2, lsErr := exec.Command("ls", "-la", "../").Output()
	if lsErr != nil {
		log.Fatal(lsErr)
	}
	log.Println(string(out2))

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

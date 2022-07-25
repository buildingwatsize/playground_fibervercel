package handler

import (
	"log"
	"net/http"
	"os/exec"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
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

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	adaptor.FiberHandlerFunc(EchoAPI)(w, r)
}

func EchoAPI(c *fiber.Ctx) error {
	c.Response().Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	log.Println("Reached: Echo API")

	config := newConfiguration()
	log.Printf("%+v\n", config)
	return c.JSON(config)
}

type Configuration struct {
	FirstTimeMessage string
}

func newConfiguration() Configuration {
	return Configuration{
		FirstTimeMessage: viper.GetString("FIRST_TIME_MESSAGE"),
	}
}

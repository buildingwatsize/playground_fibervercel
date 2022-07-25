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
	out, lsErr := exec.Command("ls", "-l").Output()
	if lsErr != nil {
		log.Fatal(lsErr)
	}
	log.Println(string(out))

	viper.SetConfigFile(".env")
	viper.AddConfigPath("/api")
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

package handler

import (
	"log"
	"net/http"

	"playground/config"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
)

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	adaptor.FiberHandlerFunc(EchoAPI)(w, r)
}

func EchoAPI(c *fiber.Ctx) error {
	c.Response().Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	log.Println("Reached: Echo API")

	conf := config.NewConfiguration()
	log.Printf("%+v\n", conf)
	return c.JSON(conf)
}

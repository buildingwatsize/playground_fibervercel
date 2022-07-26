package handler

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"playground/config"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
)

const (
	TIMEOUT string = "1m"
)

func SpotifyHandler(w http.ResponseWriter, r *http.Request) {
	adaptor.FiberHandlerFunc(SpotifyClientAPI)(w, r)
}

func SpotifyClientAPI(c *fiber.Ctx) error {

	log.Println("Reached: Spotify Client API")

	conf := config.NewConfiguration()
	log.Printf("%+v\n", conf)

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	method := "GET"
	url := "https://api.spotify.com/v1/me/tracks?market=TH"
	headers := []struct {
		Key   string
		Value string
	}{
		{Key: fiber.HeaderAccept, Value: fiber.MIMEApplicationJSON},
		{Key: fiber.HeaderContentType, Value: fiber.MIMEApplicationJSON},
		{Key: fiber.HeaderAuthorization, Value: fmt.Sprintf("Bearer %s", conf.AnotherVariable)},
	}

	timeoutDuration, _ := time.ParseDuration(TIMEOUT)
	client := &http.Client{
		Timeout: timeoutDuration,
	}
	// bodyReader := strings.NewReader(body)
	request, _ := http.NewRequest(method, url, nil)

	for _, h := range headers {
		request.Header.Set(h.Key, h.Value)
	}

	response, err := client.Do(request)
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	log.Printf("%+v\n", response)

	responseBody, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	responseData := fiber.Map{}
	errMarshal := json.Unmarshal(responseBody, &responseData)
	if errMarshal != nil {
		return c.JSON(fiber.Map{
			"error": errMarshal.Error(),
		})
	}
	return c.Status(response.StatusCode).JSON(fiber.Map{
		"response": responseData,
	})
}

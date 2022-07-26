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

var (
	isRunning          bool  = false
	currentTotalTracks int64 = 0
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
	url := "https://api.spotify.com/v1/me/tracks?offset=0&limit=10&market=TH"
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

	responseBody, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	responseData := fiber.Map{}
	errMarshal := json.Unmarshal(responseBody, &responseData)
	if errMarshal != nil {
		return c.JSON(fiber.Map{
			"error": errMarshal.Error(),
		})
	}

	itemsList := ConvertResponseToStruct(responseData["items"].([]interface{}))
	totalTracks := responseData["total"].(int64)
	if !isRunning {
		currentTotalTracks = totalTracks
		isRunning = true
	} else if currentTotalTracks < totalTracks {
		addingList := itemsList[0 : totalTracks-currentTotalTracks]
		log.Printf("Adding List: %+v\n", addingList)
	} else {
		log.Println("just skipping")
	}

	return c.Status(response.StatusCode).JSON(fiber.Map{
		"response": responseData,
	})
}

func ConvertResponseToStruct(items []interface{}) []ItemsModel {
	var itemsList []ItemsModel
	dataByte, _ := json.Marshal(items)
	_ = json.Unmarshal(dataByte, &itemsList)
	return itemsList
}

type ItemsModel struct {
	Track struct {
		URI string `json:"uri"`
	} `json:"track"`
}

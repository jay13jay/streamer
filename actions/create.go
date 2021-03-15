package actions

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"fmt"
	"net/http"
	"os"

	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func CreateHandler(c buffalo.Context) error {
	createStream()
	return c.Render(http.StatusOK, r.HTML("index.html"))
}

type Payload struct {
	Name     string    `json:"name"`
	Profiles []Profile `json:"profiles"`
}

// Profile is linked to Payload struct with the []Profile object
type Profile struct {
	Name    string `json:"name"`
	Bitrate int    `json:"bitrate"`
	Fps     int    `json:"fps"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
}

func createStream() {
	bearer := string("Bearer " + os.Getenv("APIKEY"))
	apiEndpoint := string(os.Getenv("APIENDPOINT"))

	// Set up the payload to create the streams
	data := Payload{
		Name: "test_stream",
		Profiles: []Profile{
			{
				Name:    "720p",
				Bitrate: 2000000,
				Fps:     30,
				Width:   1280,
				Height:  720,
			},
			{
				Name:    "480p",
				Bitrate: 1000000,
				Fps:     30,
				Width:   854,
				Height:  480,
			},
			{
				Name:    "360p",
				Bitrate: 500000,
				Fps:     30,
				Width:   640,
				Height:  360,
			},
		},
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		// handle err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", apiEndpoint, body)
	if err != nil {
		// handle err
	}

	// Set request headers with content type and API key
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", bearer)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bytes))
	// return (string(bytes))
	return
}

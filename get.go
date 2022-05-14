package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func get(date string) []byte {

	const url string = "https://api.nasa.gov/mars-photos/api/v1/rovers/%s/photos?earth_date=%s&camera=%s&api_key=%s"

	resp, err := http.Get(fmt.Sprintf(url, selRover, date, selCam, apiKey))

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	json, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	return json
}

package main

import (
	"encoding/json"
	"fmt"
	"time"
)

/**
 * Returns JSON of most recent 10 days worth of photos from Mars rovers
 *
 * Rover name and camera args can be set when executing the binary, e.g './marsimg curiosity navcam'.
 *
 * The camera argument can be omitted, or both arguments can be omitted.
 *
 * The default arg for rover is 'curiosity'. Currently this is the only supported rover.
 *
 * The default arg for camera is 'fhaz' (Front Hazard Cam). Supported cams:
 * fhaz - Front Hazard Avoidance Camera
 * rhaz - Rear Hazard Avoidance Camera
 * mast - Mast Camera
 * navcam - Navigation Camera
 *
 * For frequent use, set the apiKey var below with your https://api.nasa.gov API key before compiling.
 */

//args
var apiKey string = "DEMO_KEY"
var selRover string
var selCam string

func main() {

	// validate passed arguments
	validateSetArgs(&selRover, &selCam)

	// get the 10 most recent dates
	dates := dateRange()

	// init cache
	var store cache
	store.init()

	// get the photos
	var resp resp

	for _, date := range dates {

		cached := false

		// try to get photos for the specified rover + cam + date from cache
		if _, ok := store.Rovers[selRover]; ok {

			if _, ok := store.Rovers[selRover][selCam]; ok {

				if _, ok := store.Rovers[selRover][selCam][date]; ok {

					cached = true
				}
			}
		}

		// otherwise hit https://api.nasa.gov for non-cached photos
		if !cached {
			respJson := get(date)
			json.Unmarshal(respJson, &resp)

			// Cache up to 3 photos if there are any for this date
			for i := 0; i < 3; i++ {

				if i < len(resp.Photos) {

					store.Rovers[selRover][selCam][date] = append(store.Rovers[selRover][selCam][date], resp.Photos[i])

				} else if len(resp.Photos) == 0 {

					store.Rovers[selRover][selCam][date] = []photo{}
				}
			}
		}

		store.serialize()
	}

	ret, _ := json.Marshal(store.format())
	fmt.Println(string(ret))
}

// Makes strings of the most recent 10 dates
func dateRange() []string {

	dates := make([]string, 10)

	for i := 1; i <= 10; i++ {

		t := time.Now().AddDate(0, 0, -i)
		dates[i-1] = t.Format("2006-01-02")
	}

	return dates
}

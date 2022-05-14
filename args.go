package main

import (
	"log"
	"os"
)

type options struct {
	args []string
}

// Supported rover options
var validRovers options = options{

	args: []string{"curiosity"},
}

/**
 * Supported camera options
 *
 * fhaz - Front Hazard Avoidance Camera
 * rhaz - Rear Hazard Avoidance Camera
 * mast - Mast Camera
 * navcam - Navigation Camera
 */
var validCams options = options{

	args: []string{"fhaz", "rhaz", "mast", "navcam"},
}

func (o options) contains(needle string) bool {

	for _, item := range o.args {

		if item == needle {
			return true
		}
	}
	return false
}

func validateSetArgs(selRover *string, selCam *string) {

	*selRover = validRovers.args[0]
	*selCam = validCams.args[0]

	if len(os.Args) == 2 {

		*selRover = validateRover(&os.Args[1])

	} else if len(os.Args) == 3 {

		*selRover = validateRover(&os.Args[1])
		*selCam = validateCam(&os.Args[2])
	}
}

func validateRover(passedRover *string) string {

	if !validRovers.contains(*passedRover) {

		log.Fatalln("An invalid rover was passed.")
	}

	return *passedRover
}

func validateCam(passedCam *string) string {

	if !validCams.contains(*passedCam) {

		log.Fatalln("An invalid camera was passed.")
	}

	return *passedCam
}

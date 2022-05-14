package main

import (
	"encoding/gob"
	"errors"
	"log"
	"os"
)

// Cache of images

// Structure:
//
// "rover": [
//		"cam": [
//			"date": [
//				photos...
//			]
//		]
// ]
type cache struct {
	Rovers map[string]map[string]map[string][]photo
}

// Path of serialized cache struct
var tmp string = "/tmp/marsimg"

func (cache *cache) init() {

	if _, err := os.Stat(tmp); err == nil {

		cache.unserialize()
		cache.allocate(&selRover, &selCam)

	} else if errors.Is(err, os.ErrNotExist) {

		cache.createFile()
		cache.allocate(&selRover, &selCam)
	}
}

func (cache *cache) allocate(rover *string, cam *string) {

	if cache.Rovers == nil {
		cache.Rovers = make(map[string]map[string]map[string][]photo)
	}

	// check if rover key exists in cache; if not, allocate it
	if _, ok := cache.Rovers[*rover]; !ok {

		cache.Rovers[*rover] = make(map[string]map[string][]photo)
	}

	// check if cam key for rover exists in cache; if not, allocate it
	if _, ok := cache.Rovers[*rover][*cam]; !ok {

		cache.Rovers[*rover][*cam] = make(map[string][]photo)
	}
}

func (*cache) createFile() *os.File {

	var file *os.File
	var err error

	if file, err = os.Create(tmp); err != nil {

		log.Fatalln("Error serializing the cache.")
	}

	return file
}

func (cache *cache) serialize() {

	file := cache.createFile()

	enc := gob.NewEncoder(file)

	if err := enc.Encode(cache); err != nil {

		log.Fatalln("Error serializing the cache.")
	}

	file.Close()
}

func (cache *cache) unserialize() {

	var file *os.File
	var err error

	if file, err = os.Open(tmp); err != nil {

		log.Fatalln("Error opening cache file.")
	}

	dec := gob.NewDecoder(file)

	// Decode
	if err := dec.Decode(cache); err != nil {

		log.Fatalln("Error unserializing the cache.")
	}
}

func (cache *cache) format() map[string][3]string {

	result := make(map[string][3]string)

	for date := range cache.Rovers[selRover][selCam] {

		photos := [3]string{}

		for i := 0; i < len(cache.Rovers[selRover][selCam][date]); i++ {

			photos[i] = cache.Rovers[selRover][selCam][date][i].Img_src
		}

		result[date] = photos
	}

	return result
}

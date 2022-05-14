package main

type resp struct {
	Photos []photo
}

type photo struct {
	Id         string
	Sol        uint32
	Camera     camera
	Img_src    string
	Earth_date string
	Rover      rover
}

type camera struct {
	Id        string
	Name      string
	Rover_id  uint8
	Full_name string
}

type rover struct {
	Id           string
	Name         string
	Landing_date uint8
	Launch_date  string
	Status       string
}

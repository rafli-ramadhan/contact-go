package model

type ContactRequest struct {
	Name   string `json:"name"`
	NoTelp string `json:"no_telp"`
}

type Contact struct {
	Id     int
	Name   string
	NoTelp string
}

var ContactSlice []Contact = []Contact{}
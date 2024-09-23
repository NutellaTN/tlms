package main

// This service will expose an API that allows other services to control the lights (GPIO pins).


import (
	"fmt"
	"log"
	"net/http"
	"github.com/stianeikeland/go-rpio"
)

var (
	redPin    = rpio.Pin(9)
	yellowPin = rpio.Pin(10)
	greenPin  = rpio.Pin(11)
)

func setupPins() {
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		log.Fatal("Failed to initialize GPIO")
	}
	redPin.Output()
	yellowPin.Output()
	greenPin.Output()
}

func controlLight(w http.ResponseWriter, r *http.Request) {
	light := r.URL.Query().Get("light")
	state := r.URL.Query().Get("state")

	switch light {
	case "red":
		if state == "ON" {
			redPin.High()
		} else {
			redPin.Low()
		}
	case "yellow":
		if state == "ON" {
			yellowPin.High()
		} else {
			yellowPin.Low()
		}
	case "green":
		if state == "ON" {
			greenPin.High()
		} else {
			greenPin.Low()
		}
	}
	fmt.Fprintf(w, "Controlled %s light: %s", light, state)
}

func main() {
	setupPins()
	http.HandleFunc("/control-light", controlLight)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

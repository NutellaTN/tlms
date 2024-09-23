package main

// This service manages the traffic light's logic (timings and transitions). Instead of directly controlling the GPIO pins, it will send HTTP requests to the Light Controller Service.

import (
	"fmt"
	"net/http"
	"time"
)

func sendRequest(light string, state string) {
	url := fmt.Sprintf("http://light-controller-service/control-light?light=%s&state=%s", light, state)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to control light:", err)
		return
	}
	defer resp.Body.Close()
}

func controlTrafficLights() {
	for {
		// Red
		sendRequest("red", "ON")
		time.Sleep(time.Second * 4)

		// Red and yellow
		sendRequest("yellow", "ON")
		time.Sleep(time.Second * 2)

		// Green
		sendRequest("red", "OFF")
		sendRequest("yellow", "OFF")
		sendRequest("green", "ON")
		time.Sleep(time.Second * 6)

		// Yellow
		sendRequest("green", "OFF")
		sendRequest("yellow", "ON")
		time.Sleep(time.Second * 3)

		// Yellow off
		sendRequest("yellow", "OFF")
	}
}

func main() {
	controlTrafficLights()
}

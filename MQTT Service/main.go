package main

//This service will handle MQTT communication, subscribing to a topic to receive commands or publish updates when the light changes.


import (
	"encoding/json"
	"fmt"
	"github.com/yosssi/gmq/mqtt"
	"github.com/yosssi/gmq/mqtt/client"
)

type DeviceTwinUpdate struct {
	State string `json:"state,omitempty"`
}

func connectToMqtt() *client.Client {
	cli := client.New(&client.Options{
		ErrorHandler: func(err error) {
			fmt.Println(err)
		},
	})
	err := cli.Connect(&client.ConnectOptions{
		Network:  "tcp",
		Address:  "127.0.0.1:1883",
		ClientID: []byte("mqtt-client"),
	})
	if err != nil {
		panic(err)
	}
	return cli
}

func publishUpdate(cli *client.Client, state string) {
	updateMessage := DeviceTwinUpdate{State: state}
	twinUpdateBody, _ := json.Marshal(updateMessage)
	cli.Publish(&client.PublishOptions{
		TopicName: []byte("device/twin/update"),
		QoS:       mqtt.QoS0,
		Message:   twinUpdateBody,
	})
}

func main() {
	cli := connectToMqtt()
	defer cli.Terminate()

	// Example of publishing state updates
	publishUpdate(cli, "Red Light ON")
}

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

func publishUpdate(cli *client.Client, property string, propertyValue string) {
	updateMessage := DeviceTwinUpdate{State: state}
	twinUpdateBody, _ := json.Marshal(updateMessage)
	cli.Publish(&client.PublishOptions{
		TopicName: []byte("device/twin/update"),
		QoS:       mqtt.QoS0,
		Message:   twinUpdateBody,
	})
}

//createActualUpdateMessage function is used to create the device twin update message
func createActualUpdateMessage(actualValue string, property string) DeviceTwinUpdate {
	var deviceTwinUpdateMessage DeviceTwinUpdate
	actualMap := map[string]*MsgTwin{property: {Actual: &TwinValue{Value: &actualValue}, Metadata: &TypeMetadata{Type: "Updated"}}}
	deviceTwinUpdateMessage.Twin = actualMap
	return deviceTwinUpdateMessage
}


func main() {
	cli := connectToMqtt()
	defer cli.Terminate()

	// Example of publishing state updates
	//publishUpdate(cli, "Red","ON")
}

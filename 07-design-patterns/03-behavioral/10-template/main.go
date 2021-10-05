package main

import (
	"fmt"
	"time"
)

type IConnector interface {
	Connect()
	Disconnect()
}

type Connector struct {
	Self IConnector
}

func (c Connector) Connect() {
	c.Self.Connect()
}

func (c Connector) Disconnect() {
	c.Self.Disconnect()
}

func (c Connector) Reconnect() {
	c.Self.Disconnect()
	c.Self.Connect()
}

type PostgreConnector struct {

}

func (r PostgreConnector) Connect() {
	fmt.Println("Connecting...")
}

func (r PostgreConnector) Disconnect() {
	fmt.Println("Disconnecting...")
}

func main() {
	var connector Connector

	connector = Connector{PostgreConnector{}}

	connector.Connect()
	time.Sleep(100 * time.Millisecond)
	connector.Reconnect()
	time.Sleep(100 * time.Millisecond)
	connector.Disconnect()
}
package main

import (
	"log"

	"github.com/zeromq/goczmq"
)

func main() {
	// Create a router socket and bind it to port 5555.
	router, err := goczmq.NewRouter("tcp://*:5555")
	if err != nil {
		log.Fatal(err)
	}
	defer router.Destroy()

	log.Println("router created and bound")

	// Create a dealer socket and connect it to the router.
	dealer, err := goczmq.NewDealer("tcp://127.0.0.1:5555")
	if err != nil {
		log.Fatal(err)
	}
	defer dealer.Destroy()

	log.Println("dealer created and connected")

	// Send a 'Hello' message from the dealer to the router.
	// Here we send it as a frame ([]byte), with a FlagNone
	// flag to indicate there are no more frames following.
	err = dealer.SendFrame([]byte("Hello"), goczmq.FlagNone)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("dealer sent 'Hello'")

	// Receive the message. Here we call RecvMessage, which
	// will return the message as a slice of frames ([][]byte).
	// Since this is a router socket that support async
	// request / reply, the first frame of the message will
	// be the routing frame.
	request, err := router.RecvMessage()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("router received '%s' from '%v'", request[1], request[0])

	// Send a reply. First we send the routing frame, which
	// lets the dealer know which client to send the message.
	// The FlagMore flag tells the router there will be more
	// frames in this message.
	err = router.SendFrame(request[0], goczmq.FlagMore)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("router sent 'World'")

	// Next send the reply. The FlagNone flag tells the router
	// that this is the last frame of the message.
	err = router.SendFrame([]byte("World"), goczmq.FlagNone)
	if err != nil {
		log.Fatal(err)
	}

	// Receive the reply.
	reply, err := dealer.RecvMessage()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("dealer received '%s'", string(reply[0]))
}
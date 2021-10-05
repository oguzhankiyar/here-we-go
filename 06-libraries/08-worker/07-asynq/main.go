package main

import (
	"os"
	"os/signal"
)

func main() {
	go RunServer()
	go RunClient()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}
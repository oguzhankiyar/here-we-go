package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	Sample("Notify", Notify)
	Sample("Ignore", Ignore)
	Sample("Reset", Reset)
	Sample("Stop", Stop)
}

func Notify() {
	fmt.Println("Press Ctrl + C")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	s := <-c

	fmt.Println()
	fmt.Println("Got signal:", s)
}

func Ignore() {
	signal.Ignore(os.Interrupt)
	fmt.Println("Ignored interrupt for 2 secs")
	time.Sleep(2 * time.Second)
	signal.Reset(os.Interrupt)

	if signal.Ignored(os.Interrupt) {
		fmt.Println("Ignore clear")
	}
}

func Reset() {
	signal.Reset()

	fmt.Println("Reset completed")
}

func Stop() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	fmt.Println("Waiting interrupt for 2 secs, then stops")

	select {
	case x := <-c:
		fmt.Println("Received", x)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout, stopped notify")
		signal.Stop(c)
	}
}

func Sample(name string, fn func()) {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}

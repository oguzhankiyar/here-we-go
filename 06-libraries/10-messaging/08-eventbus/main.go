package main

import (
	"fmt"

	"github.com/asaskevich/EventBus"
)

func main() {
	bus := EventBus.New()

	bus.Subscribe("main:calculator", calculator)
	bus.Publish("main:calculator", 20, 40)
	bus.Unsubscribe("main:calculator", calculator)
}

func calculator(a int, b int) {
	fmt.Printf("%d\n", a + b)
}
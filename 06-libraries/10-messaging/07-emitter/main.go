package main

import "github.com/olebedev/emitter"

func main() {
	e := &emitter.Emitter{}
	go func(){
		<-e.Emit("change", 42)
		<-e.Emit("change", 37)
		e.Off("*")
	}()

	for event := range e.On("change") {
		println(event.Int(0))
	}
}
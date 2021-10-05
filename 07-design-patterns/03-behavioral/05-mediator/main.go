package main

import (
	"fmt"
	"reflect"
)

type Command interface {

}

type CommandHandler interface {
	Handle(command Command) interface{}
}

type Mediator struct {
	Map map[string]reflect.Type
}

func NewMediator() *Mediator {
	return &Mediator{make(map[string]reflect.Type)}
}

func (m Mediator) Register(command Command, commandHandler CommandHandler) {
	commandType := reflect.TypeOf(command)
	commandHandlerType := reflect.TypeOf(commandHandler)

	m.Map[commandType.Name()] = commandHandlerType
}

func (m Mediator) Send(command Command) interface{} {
	commandType := reflect.TypeOf(command)

	handlerType, ok := m.Map[commandType.Name()]
	if ok {
		handler := reflect.New(handlerType).Interface().(CommandHandler)
		return handler.Handle(command)
	}

	return nil
}

type HelloRequest struct {
	Name string
}

type HelloResponse struct {
	Greeting string
}

type HelloRequestHandler struct {

}

func (h HelloRequestHandler) Handle(command Command) interface{} {
	helloRequest, ok := command.(HelloRequest)
	helloResponse := HelloResponse{}

	if ok {
		helloResponse.Greeting = fmt.Sprintf("Hello, %s!\n", helloRequest.Name)
	}

	return helloResponse
}

func main() {
	mediator := NewMediator()
	mediator.Register(HelloRequest{}, HelloRequestHandler{})

	command := HelloRequest{"Gopher"}
	result := mediator.Send(command).(HelloResponse)

	fmt.Println(result.Greeting)
}
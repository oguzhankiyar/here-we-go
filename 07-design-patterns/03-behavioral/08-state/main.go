package main

import (
	"fmt"
	"time"
)

type State struct {
	Id	int
}

func NewState() *State {
	return &State{time.Now().Nanosecond()}
}

type Context struct {
	State State
}

func (c *Context) SetState(state *State) {
	c.State = *state
}

func (c Context) GetId() int {
	return c.State.Id
}

func main() {
	context := Context{}

	state1 := NewState()
	context.SetState(state1)
	fmt.Println(context.GetId())

	time.Sleep(111 * time.Millisecond)

	state2 := NewState()
	context.SetState(state2)
	fmt.Println(context.GetId())
}
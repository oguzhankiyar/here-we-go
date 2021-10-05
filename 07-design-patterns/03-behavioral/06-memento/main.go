package main

import (
	"fmt"
)

type Originator struct {
	State string
}

func (o *Originator) Set(state string) {
	fmt.Println("Setting state to " + state)
	o.State = state
}

func (o *Originator) SaveToMemento() Memento {
	fmt.Println("Saving to Memento.")
	return Memento{o.State}
}

func (o *Originator) RestoreFromMemento(memento Memento) {
	o.State = memento.GetSavedState()
	fmt.Println("State after restoring from Memento: " + o.State)
}

type Memento struct {
	State string
}

func (m *Memento) GetSavedState() string {
	return m.State
}

func main() {
	history := make([]Memento, 0)
	originator := new(Originator)

	originator.Set("State1")
	originator.Set("State2")
	history = append(history, originator.SaveToMemento())

	originator.Set("State3")
	history = append(history, originator.SaveToMemento())

	originator.Set("State4")
	originator.RestoreFromMemento(history[1])
}
package main

import (
	"context"
	"fmt"

	"github.com/qmuntal/stateless"
)

const (
	stateA = "A"
	stateB = "B"
	stateC = "C"

	trigger1 = "trigger1"
	trigger2 = "trigger2"
	trigger3 = "trigger3"
)

func main() {
	machine := stateless.NewStateMachine(stateA)

	machine.
		Configure(stateA).
		OnEntry(func(ctx context.Context, args ...interface{}) error {
			fmt.Println("entering: state A")
			return nil
		}).
		Permit(trigger1, stateB)

	machine.
		Configure(stateB).
		OnEntry(func(ctx context.Context, args ...interface{}) error {
			fmt.Println("entering: state B")
			return nil
		}).
		Permit(trigger2, stateC)

	machine.
		Configure(stateC).
		OnEntry(func(ctx context.Context, args ...interface{}) error {
			fmt.Println("entering: state C")
			return nil
		}).
		Permit(trigger3, stateA)

	triggers := []string{
		trigger1,
		trigger2,
		trigger3,
		trigger1,
	}

	for _, v := range triggers {
		fmt.Println(machine)
		fmt.Println("firing:", v)

		err := machine.Fire(v)
		if err != nil {
			fmt.Println("error: ", err)
			break
		}

		fmt.Println(machine)
		fmt.Println()
	}
}
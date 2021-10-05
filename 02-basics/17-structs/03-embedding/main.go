package main

import (
	"fmt"
)

func main() {
	type Employee struct {
		FirstName	string
		LastName	string
	}

	type Manager struct {
		Employee

		Unit	string
	}

	manager := Manager{
		Employee: Employee{
			FirstName: "Gopher",
			LastName: "Go",
		},
		Unit: "Development",
	}

	fmt.Printf("%+v\n", manager)
}
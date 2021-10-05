package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/r3labs/diff/v2"
)

type Order struct {
	ID    		string 		`diff:"id"`
	Customer 	string 		`diff:"customer"`
	Items 		[]int  		`diff:"items"`
}

func main() {
	before := Order{
		ID: "1234",
		Customer: "Gopher I",
		Items: []int{1, 2, 3, 4},
	}

	after := Order{
		ID: "1234",
		Customer: "Gopher II",
		Items: []int{1, 2, 4},
	}

	changelog, err := diff.Diff(before, after)
	if err != nil {
		log.Fatal(err)
		return
	}

	PrintJson(changelog)
}

func PrintJson(data interface{}) {
	j, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("json error")
		return
	}
	fmt.Printf("%s\n", j)
}
package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

func main() {
	Sample("Simple", Simple)
	Sample("Embedded", Embedded)
	Sample("Errors", Errors)
	Sample("Metadata", Metadata)
	Sample("Omitempty", Omitempty)
	Sample("RemainingData", RemainingData)
	Sample("Tags", Tags)
	Sample("WeaklyTypedInput", WeaklyTypedInput)
}

func Simple() {
	type Order struct {
		Id		string
		Amount  int
		Items 	[]string
		UTM  	map[string]string
	}

	input := map[string]interface{}{
		"id": "Gopher",
		"amount": 10,
		"items": []string{"computer", "monitor", "keyboard", "mouse"},
		"utm": map[string]string{
			"source": "twitter",
		},
	}

	var order Order
	err := mapstructure.Decode(input, &order)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("id", order.Id)
	fmt.Println("amount:", order.Amount)
	fmt.Println("items:", order.Items)
	fmt.Println("utm:", order.UTM)
}

func Embedded() {
	type Family struct {
		LastName string
	}

	type Location struct {
		City string
	}

	type Person struct {
		Family    	`mapstructure:",squash"`
		Location  	`mapstructure:",squash"`
		FirstName 	string
	}

	input := map[string]interface{}{
		"FirstName": "Gopher",
		"LastName":  "Go",
		"City":      "Metropolis",
	}

	var result Person
	err := mapstructure.Decode(input, &result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s %s, %s\n", result.FirstName, result.LastName, result.City)
}

func Errors() {
	type Person struct {
		Name   string
		Age    int
		Emails []string
		Extra  map[string]string
	}

	input := map[string]interface{}{
		"name":   123,
		"age":    "bad value",
		"emails": []int{1, 2, 3},
	}

	var result Person
	err := mapstructure.Decode(input, &result)
	if err == nil {
		panic("should have an error")
	}

	fmt.Println(err.Error())
}

func Metadata() {
	type User struct {
		Id  		int
		Username 	string
	}

	input := map[string]interface{} {
		"id": 100,
		"username": "gopher",
		"password": "123@456",
	}

	var md mapstructure.Metadata
	var user User

	config := &mapstructure.DecoderConfig{
		Metadata: &md,
		Result:   &user,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	if err := decoder.Decode(input); err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("Keys: %#v\n", md.Keys)
	fmt.Printf("Unused keys: %#v\n", md.Unused)
}

func Omitempty() {
	type Family struct {
		LastName string
	}

	type Location struct {
		City string
	}

	type Person struct {
		*Family   	`mapstructure:",omitempty"`
		*Location 	`mapstructure:",omitempty"`
		Age       	int
		FirstName 	string
	}

	result := &map[string]interface{}{}
	input := Person{FirstName: "Somebody"}

	err := mapstructure.Decode(input, &result)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%+v\n", result)
}

func RemainingData() {
	type Person struct {
		Name  string
		Age   int
		Other map[string]interface{} `mapstructure:",remain"`
	}

	input := map[string]interface{}{
		"name":  "Gopher",
		"age":   1,
		"email": "gopher@golang.org",
	}

	var person Person
	err := mapstructure.Decode(input, &person)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("name:", person.Name)
	fmt.Println("age:", person.Age)
	fmt.Println("other:", person.Other)
}

func Tags() {
	type Person struct {
		Name string `mapstructure:"person_name"`
		Age  int    `mapstructure:"person_age"`
	}

	input := map[string]interface{}{
		"person_name": "Gopher",
		"person_age":  10,
	}

	var person Person
	err := mapstructure.Decode(input, &person)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("name:", person.Name)
	fmt.Println("age:", person.Age)
}

func WeaklyTypedInput() {
	type Person struct {
		Name   string
		Age    int
		Emails []string
	}

	input := map[string]interface{}{
		"name":   123,
		"age":    "42",
		"emails": map[string]interface{}{},
	}

	var person Person
	config := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result: &person,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	err = decoder.Decode(input)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("name:", person.Name)
	fmt.Println("age:", person.Age)
	fmt.Println("emails:", person.Emails)
}

func Sample(name string, fn func())  {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}
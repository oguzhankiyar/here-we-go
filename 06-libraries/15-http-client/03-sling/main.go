package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dghubble/sling"
)

func main() {
	Sample("Get", Get)
	Sample("Post", Post)
	Sample("Put", Put)
	Sample("Delete", Delete)
	Sample("Add", Add)
	Sample("Base", Base)
}

func Get() {
	var client http.Client

	type Params struct {
		Count int `url:"count,omitempty"`
	}
	params := &Params{Count: 5}

	req, err := sling.New().
		Get("https://run.mocky.io/v3/41ac0f88-ff00-4fb5-a7b7-a749781c3260").
		QueryStruct(params).
		Request()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

func Post() {
	var client http.Client

	type Payload struct {
		Title string `url:"title"`
	}
	payload := &Payload{Title: "The Gopher"}

	req, err := sling.New().
		Post("https://run.mocky.io/v3/41ac0f88-ff00-4fb5-a7b7-a749781c3260").
		BodyJSON(payload).
		Request()

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

func Put() {
	var client http.Client

	type Payload struct {
		Title string `url:"title"`
	}
	payload := &Payload{Title: "The Gopher"}

	req, err := sling.New().
		Put("https://run.mocky.io/v3/41ac0f88-ff00-4fb5-a7b7-a749781c3260").
		BodyJSON(payload).
		Request()

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

func Delete() {
	var client http.Client

	req, err := sling.New().
		Delete("https://run.mocky.io/v3/41ac0f88-ff00-4fb5-a7b7-a749781c3260").
		Request()

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

func Add() {
	var client http.Client

	req, err := sling.New().
		Add("Content-Type", "application/json").
		Add("Authorization", "Bearer TOKEN_HERE").
		Get("https://run.mocky.io/v3/41ac0f88-ff00-4fb5-a7b7-a749781c3260").
		Request()

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

func Base() {
	var client http.Client

	base := sling.New().
		Base("https://run.mocky.io/v3/").
		Client(&client)

	req, err := base.New().
		Get("41ac0f88-ff00-4fb5-a7b7-a749781c3260").
		Request()

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

func Sample(name string, fn func())  {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}
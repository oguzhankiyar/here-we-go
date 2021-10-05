package main

import (
	"fmt"

	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
)

func main() {
	Sample("Get", Get)
	Sample("Post", Post)
}

func Get() {
	cli := gentleman.New()
	cli.URL("https://run.mocky.io")

	req := cli.Request()
	req.Path("/v3/41ac0f88-ff00-4fb5-a7b7-a749781c3260")
	req.SetHeader("Client", "gentleman")
	req.Method("GET")

	res, err := req.Send()
	if err != nil {
		fmt.Printf("Request error: %s\n", err)
		return
	}

	if !res.Ok {
		fmt.Printf("Invalid server response: %d\n", res.StatusCode)
		return
	}

	fmt.Printf("Status: %d\n", res.StatusCode)
	fmt.Printf("Body:\n%s\n", res.String())
}

func Post() {
	cli := gentleman.New()
	cli.URL("https://run.mocky.io/v3/41ac0f88-ff00-4fb5-a7b7-a749781c3260")

	req := cli.Request()
	req.SetHeader("Content-Type", "application/json")
	req.SetHeader("Authorization", "Bearer TOKEN_HERE")
	req.Method("POST")
	req.Use(body.JSON(map[string]string{"foo": "bar"}))

	res, err := req.Send()
	if err != nil {
		fmt.Printf("Request error: %s\n", err)
		return
	}

	if !res.Ok {
		fmt.Printf("Invalid server response: %d\n", res.StatusCode)
		return
	}

	fmt.Printf("Status: %d\n", res.StatusCode)
	fmt.Printf("Body:\n%s\n", res.String())
}

func Sample(name string, fn func())  {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}
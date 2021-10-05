package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

func main() {
	data := map[string]string{
		"name": "gopher",
	}

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	result, err := json.MarshalToString(&data)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println(result)

	err = json.UnmarshalFromString(result, &data)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println(data)
}
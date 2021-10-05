package main

import (
	"fmt"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

func main() {
	Sample("Simple", Simple)
	Sample("Watch", Watch)
}

func Simple() {
	k := koanf.New(".")

	provider := file.Provider("./06-libraries/02-config/02-koanf/config/config.json")

	if err := k.Load(provider, json.Parser()); err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("appName:", k.String("app.name"))
}

func Watch() {
	k := koanf.New(".")

	f := file.Provider("./06-libraries/02-config/02-koanf/config/config.yaml")
	if err := k.Load(f, yaml.Parser()); err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("appName:", k.String("app.name"))

	done := make(chan bool)

	f.Watch(func(event interface{}, err error) {
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		fmt.Println("config changed.")

		k = koanf.New(".")
		k.Load(f, yaml.Parser())

		fmt.Println("appName:", k.String("app.name"))

		done <- true
	})

	fmt.Println("waiting for change config.yaml")

	<-done
}

func Sample(name string, fn func())  {
	fmt.Println(">", name)
	fn()
	fmt.Println()
}
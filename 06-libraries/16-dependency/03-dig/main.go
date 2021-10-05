package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"go.uber.org/dig"
)

func main() {
	type Config struct {
		Prefix string
	}

	c := dig.New()
	err := c.Provide(func() (*Config, error) {
		// In real program there should reading from the file, for example
		var cfg Config
		err := json.Unmarshal([]byte(`{"prefix": "[foo] "}`), &cfg)
		return &cfg, err
	})
	if err != nil {
		panic(err)
	}

	// Function to create logger by using config
	err = c.Provide(func(cfg *Config) *log.Logger {
		return log.New(os.Stdout, cfg.Prefix, 0)
	})
	if err != nil {
		panic(err)
	}

	err = c.Invoke(func(l *log.Logger) {
		l.Print("You've been invoked")
	})
	if err != nil {
		panic(err)
	}

	// создаем еще один логгер
	err = c.Provide(
		func(cfg *Config) *log.Logger {
			return log.New(os.Stdout, cfg.Prefix, 0)
		},
		dig.Name("logger2"), // передаем опцию имени
	)
	if err != nil {
		panic(err)
	}

	c = dig.New()
	c.Provide("gopher", dig.Name("username"))
	c.Provide("123456", dig.Name("password"))

	err = c.Invoke(func(p struct {
		dig.In

		U string `name:"username"`
		P string `name:"password"`
	}) {
		fmt.Println("user >>>", p.U)
		fmt.Println("pwd  >>>", p.P)
	})
}
package main

import (
	"log"

	"cli-sample/internal/application/commands"
	"cli-sample/internal/infrastructure/config/parser"
	"cli-sample/internal/infrastructure/logger"
)

func main() {
	configParser := parser.NewConfigParser("./configs", "dev", "json")
	config, err := configParser.Parse()
	if err != nil {
		log.Fatal(err)
	}

	logger := logger.NewAppLogger(config.Logger)
	err = logger.Init()
	if err != nil {
		log.Fatal(err)
	}

	commands.Config = config.App
	commands.Logger = logger

	commands.Execute()
}

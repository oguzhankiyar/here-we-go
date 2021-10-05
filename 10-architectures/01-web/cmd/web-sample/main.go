package main

import (
	"log"

	_ "web-sample/docs"
	"web-sample/internal/api"
	"web-sample/internal/infrastructure/config/parser"
	"web-sample/internal/infrastructure/logger"
	"web-sample/internal/infrastructure/persistence"
)

// @title Web Sample API
// @version 1.0
// @description This is a sample web.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:2805
// @BasePath /
// @schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
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

	postgresPersistence, err := persistence.NewPostgresPersistence(config.Persistence)
	if err != nil {
		log.Fatal(err)
	}

	api := api.NewAPI(config.App, logger, postgresPersistence)
	err = api.Init()
	if err != nil {
		log.Fatal(err)
	}

	api.Start()
}

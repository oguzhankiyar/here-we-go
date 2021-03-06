package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"

	"swag-sample/api"
	_ "swag-sample/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/testapi/get-string-by-int/", api.GetStringByInt)
	r.GET("/testapi/get-struct-array-by-string/", api.GetStructArrayByString)
	r.POST("/testapi/upload", api.Upload)

	log.Println("open http://localhost:2805/swagger/index.html")

	r.Run(":2805")
}

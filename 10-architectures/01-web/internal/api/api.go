package api

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	"web-sample/internal/api/controllers"
	"web-sample/internal/api/middlewares"
	"web-sample/internal/application/handlers/product"
	productMappers "web-sample/internal/application/mappers/product"
	"web-sample/internal/infrastructure/auth"
	"web-sample/internal/infrastructure/config/models"
	"web-sample/internal/infrastructure/logger/interfaces"
	"web-sample/internal/infrastructure/persistence"
	"web-sample/internal/persistence/repositories"
)

type API struct {
	appConfig           models.AppConfig
	logger              interfaces.Logger
	postgresPersistence *persistence.PostgresPersistence
	echo                *echo.Echo
}

func NewAPI(appConfig models.AppConfig, logger interfaces.Logger, postgresPersistence *persistence.PostgresPersistence) *API {
	return &API{
		appConfig:           appConfig,
		logger:              logger,
		postgresPersistence: postgresPersistence,
		echo:                echo.New(),
	}
}

func (a *API) Init() error {
	// Middlewares
	a.echo.Use(middleware.Logger())
	a.echo.Use(middleware.Recover())
	a.echo.Use(middleware.CORS())
	a.echo.Use(middlewares.Error())

	jwtMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &auth.AuthClaims{},
		SigningKey: []byte(a.appConfig.JwtSecret),
	})

	v := validator.New()

	homeController := controllers.NewHomeController(a.appConfig)
	healthController := controllers.NewHealthController()
	authController := controllers.NewAuthController(a.appConfig)

	productRepository := repositories.NewProductRepository(a.postgresPersistence)
	productMapper := productMappers.NewProductMapper()
	productHandler := product.NewProductHandler(productRepository, productMapper)
	productsController := controllers.NewProductsController(productHandler, v)

	// Swagger
	a.echo.GET("/swagger/*", echoSwagger.WrapHandler)

	// Home
	a.echo.GET("", homeController.Handle)

	// Health
	a.echo.GET("/health", healthController.Handle)

	// Auth
	authGroup := a.echo.Group("auth")
	authGroup.POST("/token", authController.HandleToken)
	authGroup.GET("/me", authController.HandleMe, jwtMiddleware)

	// Products
	productsGroup := a.echo.Group("products", jwtMiddleware)
	productsGroup.GET("", productsController.HandleGetAll)
	productsGroup.GET("/:id", productsController.HandleGetById)
	productsGroup.POST("", productsController.HandleCreate)
	productsGroup.PUT("/:id", productsController.HandleUpdate)
	productsGroup.DELETE("/:id", productsController.HandleDelete)

	return nil
}

func (a *API) Start() {
	a.logger.Info(a.appConfig.Name + " is starting")

	a.echo.Logger.Fatal(a.echo.Start(":" + strconv.Itoa(a.appConfig.Port)))
}

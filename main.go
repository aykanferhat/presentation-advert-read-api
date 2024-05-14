package main

import (
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"os"
	"os/signal"
	_ "presentation-advert-read-api/docs"
	"presentation-advert-read-api/infrastructure/configuration/configreader"
	"presentation-advert-read-api/infrastructure/configuration/custom_error"
	"presentation-advert-read-api/infrastructure/configuration/elastic/elasticv7"
	"presentation-advert-read-api/infrastructure/configuration/log"
	"presentation-advert-read-api/infrastructure/configuration/server"
	"presentation-advert-read-api/infrastructure/controller"
	"presentation-advert-read-api/infrastructure/handlers"
	"presentation-advert-read-api/infrastructure/repository"
	"strings"
	"syscall"
)

func main() {
	e := echo.New()

	logConfig := configreader.ReadLogConfig("log-config")
	serverConfig := configreader.ReadServerConf("server-config")
	elasticConfigMap := configreader.ReadElasticConfig("elastic-config")

	logger := log.NewLogger(logConfig.Level)
	e.Logger = logger

	// Elastic
	elasticClientMap, err := elasticv7.Initialize(elasticConfigMap)
	if err != nil {
		e.Logger.Fatal(err)
	}

	categoryElasticRepository, err := repository.NewCategoryElasticRepository(elasticClientMap, "local", "categories")
	if err != nil {
		e.Logger.Fatal(err)
	}
	advertElasticRepository, err := repository.NewAdvertElasticRepository(elasticClientMap, "local", "adverts")
	if err != nil {
		e.Logger.Fatal(err)
	}

	queryHandler, err := handlers.InitializeQueryHandler(categoryElasticRepository, advertElasticRepository)
	if err != nil {
		e.Logger.Fatal(err)
	}

	controller.NewAdvertController(e, queryHandler)
	controller.NewCategoryController(e, queryHandler)

	//Middleware
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	//HealthCheck
	server.RegisterHealthCheck(e)

	//Swagger
	server.RegisterSwaggerRedirect(e)

	e.HTTPErrorHandler = custom_error.CustomEchoHTTPErrorHandler

	//Start Server
	go func() {
		if err := e.Start(serverConfig.Port); err != nil {
			if !strings.Contains(err.Error(), "client: Server closed") {
				e.Logger.Fatal(err)
			}
		}
	}()
	serverChannel := make(chan struct{})

	// Stop Server
	go func() {
		sigint := make(chan os.Signal, 1)
		// interrupt signal sent from terminal
		signal.Notify(sigint, os.Interrupt, os.Kill)
		// sigterm signal sent from kubernetes
		signal.Notify(sigint, syscall.SIGTERM)
		<-sigint
		close(serverChannel)
	}()
	<-serverChannel
}

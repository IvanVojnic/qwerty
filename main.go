package main

import (
	"EFpractic2/pkg/handler"
	"EFpractic2/pkg/repository"
	"EFpractic2/pkg/service" //nolint:goimports
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	e := echo.New()
	logger := log.New()
	logger.Out = os.Stdout
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			log.WithFields(log.Fields{
				"URI":    values.URI,
				"status": values.Status,
			}).Info("request")
			return nil
		},
	}))

	db, err := repository.NewPostgresDB()
	if err != nil {
		log.WithFields(log.Fields{
			"Error connection to database rep.NewPostgresDB()": err,
		}).Fatal("DB ERROR CONNECTION")
	}
	defer repository.ClosePool(db)

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	handlers.InitRoutes(e)
}

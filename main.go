// Package main file to run app
package main

import (
	"EFpractic2/pkg/config"
	"EFpractic2/pkg/handler"
	"EFpractic2/pkg/repository"
	"EFpractic2/pkg/service"
	"context"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
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
	cfg, err := config.NewConfig()
	if err != nil {
		log.WithFields(log.Fields{
			"Error":  err,
			"config": cfg,
		}).Fatal("failed to get config")
	}
	dbType := 0
	initProjectStruct(cfg, e, dbType)
}

func initProjectStruct(cfg *config.Config, e *echo.Echo, dbType int) {
	var profileServ *service.AuthService
	var bookServ *service.BookActSrv
	var imgServ *service.ImgUpSrv
	switch dbType {
	case 0:
		db, err := repository.NewPostgresDB(cfg)
		if err != nil {
			log.WithFields(log.Fields{
				"Error connection to database rep.NewPostgresDB()": err,
			}).Fatal("DB ERROR CONNECTION")
		}
		defer repository.ClosePool(db)
		profileRepo := repository.NewUserAuthPostgres(db)
		bookRepo := repository.NewBookActPostgres(db)
		imgRepo := repository.NewImgPostgres(db)
		profileServ = service.NewAuthService(profileRepo)
		bookServ = service.NewBookActSrv(bookRepo)
		imgServ = service.NewImgUpSrv(imgRepo)
	case 1:
		mDB, err := repository.NewMongoDB(cfg)
		if err != nil {
			log.WithFields(log.Fields{
				"Error connection to database rep.MongoDB": err,
			}).Fatal("DB ERROR CONNECTION")
		}
		defer mDB.Client().Disconnect(context.Background())

		bookRepo := repository.NewBookActMongo(mDB)

		bookServ = service.NewBookActSrv(bookRepo)
	}

	profileHandlers := handler.NewHandler(profileServ, bookServ, imgServ)
	profileHandlers.InitRoutes(e)
}

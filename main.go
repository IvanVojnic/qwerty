// Package main file to run app
package main

import (
	_ "EFpractic2/docs"
	"EFpractic2/pkg/config"
	"EFpractic2/pkg/handler"
	"EFpractic2/pkg/repository"
	"EFpractic2/pkg/service"
	"github.com/segmentio/kafka-go"
	echoSwagger "github.com/swaggo/echo-swagger"

	"context"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
)

// @title Swagger EF_CRUD API
// @version 1.0
// @description This is a CRUD ENTITY server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:40000
// @BasePath /
func main() {
	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)

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
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
	})
	defer rdb.Close()
	rds := &repository.Redis{Client: *rdb}

	network := "tcp"
	address := "127.0.0.1:9092"
	topic := "myTopic"
	partition := 0
	conn, errKafka := kafka.DialLeader(context.Background(), network, address, topic, partition)
	if errKafka != nil {
		log.WithFields(log.Fields{
			"Error connection to database rep.NewPostgresDB()": errKafka,
		}).Fatal("DB ERROR CONNECTION")
	}
	kafkaWriter := repository.NewKafkaConn(conn)
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
		bookServ = service.NewBookActSrv(bookRepo, rds, kafkaWriter)
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

		bookServ = service.NewBookActSrv(bookRepo, rds, kafkaWriter)
	}

	profileHandlers := handler.NewHandler(profileServ, bookServ, imgServ)
	profileHandlers.InitRoutes(e)
}

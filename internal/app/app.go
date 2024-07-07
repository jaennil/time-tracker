package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jaennil/time-tracker/config"
	"github.com/jaennil/time-tracker/internal/controller/http/v1"
	"github.com/jaennil/time-tracker/internal/repository"
	"github.com/jaennil/time-tracker/internal/service"
	"github.com/jaennil/time-tracker/pkg/database"
	"github.com/jaennil/time-tracker/pkg/database/postgres"
	"github.com/jaennil/time-tracker/pkg/logger"
	"go.uber.org/zap"
)

func Run(config *config.Config) {
	log := logger.NewZapLogger()
	zapLogger, ok := log.(*logger.ZapLogger)
	if ok {
		defer zapLogger.Sync()
	}
	log.Debug("zapLogger: ", zap.Bool("ok", ok))
	log.Info("hello world")
	log.Debug("config", zap.Any("config", config))

	db, err := postgres.NewPostgres(database.Config{
		Host:     config.DBHost,
		Port:     config.DBPort,
		User:     config.DBUser,
		Name:     config.DBName,
		Password: config.DBPassword,
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatal("faled to connect to database: ", err)
	}

	handler := gin.New()
	repository := repository.NewRepository(db)
	service := service.New(repository)
	http.NewRouter(handler, service, log)
	handler.Run()
}

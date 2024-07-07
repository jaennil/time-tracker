package app

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jaennil/time-tracker/config"
	"github.com/jaennil/time-tracker/internal/controller/http/v1"
	"github.com/jaennil/time-tracker/internal/repository"
	"github.com/jaennil/time-tracker/internal/service"
	"github.com/jaennil/time-tracker/pkg/database/postgres"
	"github.com/jaennil/time-tracker/pkg/logger"
)

func Run(config *config.Config) {
	log := logger.NewZapLogger()
	log.Info("initialized logger")
	zapLogger, ok := log.(*logger.ZapLogger)
	if ok {
		defer zapLogger.Sync()
	}

	db, err := postgres.NewPostgres(config.PG_DSN)
	if err != nil {
		log.Fatal("faled to connect to database: ", err)
	}
	defer db.Close(context.Background())
	log.Info("connected to database")

	m, err := migrate.New(
		"file://migrations",
		// TODO: replace database config with fields instead of DSN and
		// build dsn from config here
		"pgx://jaennil:naen@localhost:5432/time_tracker?sslmode=disable",
	)
	// TODO: logs
	if err != nil {
		log.Fatal("err", err)
	}
	err = m.Up()
	// TODO: logs
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal("up err", err)
	}
	// TODO: logs
	log.Info("mig suc")

	handler := gin.New()
	repository := repository.NewRepository(db)
	service := service.New(repository)
	http.NewRouter(handler, service, log)
	handler.Run()
}

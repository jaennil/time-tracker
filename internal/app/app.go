package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jaennil/time-tracker/config"
	"github.com/jaennil/time-tracker/internal/controller/http/v1"
	"github.com/jaennil/time-tracker/internal/repository"
	"github.com/jaennil/time-tracker/internal/service"
	"github.com/jaennil/time-tracker/pkg/database/postgres"
	"github.com/jaennil/time-tracker/pkg/httpserver"
	"github.com/jaennil/time-tracker/pkg/logger"
	"go.uber.org/zap"
)

func Run(config *config.Config) {
	log := logger.NewZapLogger()
	log.Info("initialized logger")
	log.Debug("", zap.Any("config", config))
	zapLogger, ok := log.(*logger.ZapLogger)
	if ok {
		defer zapLogger.Sync()
	}

	db, err := postgres.NewPostgres(config.PG_DSN)
	if err != nil {
		log.Fatal("faled to connect to database: ", err)
	}
	defer func() {
		err := db.Close(context.Background())
		if err != nil {
			log.Fatal("faled to close database connection: ", err)
		}
	}()
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
	httpServer := httpserver.New(handler, httpserver.Port(config.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		log.Error("app - Run - httpServer.Notify: ", err)
	}

	log.Info("Shutdown server")

	err = httpServer.Shutdown()
	if err != nil {
		log.Error("app - Run - httpServer.Shutdown: ", err)
	}

}

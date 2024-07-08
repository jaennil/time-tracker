package app

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
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
)

func Run(config *config.Config) {
	//TODO: wrap server errors
	log := logger.NewZapLogger()
	log.Info("initialized logger")
	log.Debug("global app config: ", zap.Any("config", config))

	if zapLogger, ok := log.(*logger.ZapLogger); ok {
		defer func() {
			err := zapLogger.Sync()
			if err != nil {
				log.Error("failed to sync zap logger", err)
			}
		}()
	}

	log.Info("connecting to database")
	db, err := postgres.NewPostgres(config)
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}
	defer func() {
		log.Info("closing database connection")
		err := db.Close(context.Background())
		if err != nil {
			log.Fatal("failed to close database connection: ", err)
		}
		log.Info("database connection closed")
	}()
	log.Info("connected to database")

	log.Info("creating migration instance")
	m, err := migrate.New(
		"file://migrations",
		fmt.Sprintf("pgx://%s:%s@%s:%d/%s?sslmode=disable", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName),
	)
	if err != nil {
		log.Fatal("failed to create Migrate instance", err)
	}

	log.Info("starting migrations")
	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal("failed to migrate database", err)
	}
	log.Info("successfully migrated database")

	userApi := service.NewUserAPI(config)
	handler := gin.New()
	repositories := repository.NewRepository(db)
	services := service.New(repositories, userApi)
	http.NewRouter(handler, services, log)
	httpServer := httpserver.New(handler, httpserver.Port(config.Port))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case s := <-quit:
		log.Info("got signal" + s.String())
	case err = <-httpServer.Notify():
		log.Error("got server error", err)
	}
	log.Info("shutting down server")

	err = httpServer.Shutdown()
	if err != nil {
		log.Error("failed to shutdown server", err)
	}

}

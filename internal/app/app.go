package app

import (
	"github.com/jaennil/time-tracker/config"
	"github.com/jaennil/time-tracker/pkg/logger"
	"go.uber.org/zap"
)

func Run(config *config.Config) {
	logger := logger.New()
	logger.Info("hello world")
	logger.Debug("config", zap.Any("config", config))
}

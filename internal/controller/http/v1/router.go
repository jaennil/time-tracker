package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jaennil/time-tracker/internal/service"
	"github.com/jaennil/time-tracker/pkg/logger"
)

func NewRouter(handler *gin.Engine, services *service.Service, log logger.Loggable, validate *validator.Validate) {
	v1 := handler.Group("/v1")
	{
		NewUserRoutes(v1, services.User, log, validate)
	}
}

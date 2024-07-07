package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jaennil/time-tracker/internal/service"
	"github.com/jaennil/time-tracker/pkg/logger"
)

type userRoutes struct {
	service service.User
	logger  logger.Loggable
}

func NewUserRoutes(handler *gin.RouterGroup, userService service.User, log logger.Loggable) {
	routes := &userRoutes{userService, log}

	user := handler.Group("/user")
	{
		user.POST("", routes.create)
	}
}

func (r *userRoutes) create(c *gin.Context) {
	c.String(http.StatusOK, "/user create")
}

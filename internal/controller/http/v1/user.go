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
	var input struct {
		Passport string `json:"passportNumber"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "must provide passport number")
		return
	}

	user, err := r.service.Create(input.Passport)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, user)
}

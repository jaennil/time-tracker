package http

import (
	"errors"
	"github.com/jaennil/time-tracker/internal/repository/postgres"
	"go.uber.org/zap"
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
		user.DELETE(":id", routes.delete)
	}
}

func (r *userRoutes) create(c *gin.Context) {
	r.logger.Debug("hit endpoint", zap.String("url", c.FullPath()), zap.String("method", c.Request.Method))
	var input struct {
		Passport string `json:"passportNumber"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, "must provide passport number")
		return
	}

	user, err := r.service.Create(input.Passport)
	if err != nil {
		switch {
		case errors.Is(err, postgres.InvalidPassportFormat):
			errorResponse(c, http.StatusBadRequest, err.Error())
		default:
			r.logger.Error("failed to create user", err)
			errorResponse(c, http.StatusInternalServerError, postgres.InternalServerError.Error())
		}
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (r *userRoutes) delete(c *gin.Context) {
	r.logger.Debug("hit endpoint", zap.String("url", c.FullPath()), zap.String("method", c.Request.Method))
	id, err := readIDParam(c)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}
	err = r.service.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, postgres.RecordNotFound):
			errorResponse(c, http.StatusBadRequest, "user not founded")
		default:
			r.logger.Error("failed to delete user", err)
			errorResponse(c, http.StatusInternalServerError, postgres.InternalServerError.Error())
		}
		return
	}

	c.JSON(http.StatusOK, "user deleted")
}

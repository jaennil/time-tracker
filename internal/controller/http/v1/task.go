package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
	"github.com/jaennil/time-tracker/internal/repository/postgres"
	"github.com/jaennil/time-tracker/internal/service"
	"github.com/jaennil/time-tracker/pkg/logger"
	"go.uber.org/zap"
	"net/http"
)

type taskRoutes struct {
	service  service.Task
	logger   logger.Loggable
	validate *validator.Validate
}

func NewTaskRoutes(handler *gin.RouterGroup, taskService service.Task, log logger.Loggable, validate *validator.Validate) {
	r := &taskRoutes{taskService, log, validate}

	tasks := handler.Group("/tasks")
	{
		tasks.POST("/start", r.start)
		tasks.POST("/end", r.end)
	}
}

func (r *taskRoutes) start(c *gin.Context) {
	var input struct {
		UserId int64  `json:"user_id" binding:"required" validate:"gt=0"`
		Name   string `json:"name" binding:"required" validate:"min=1,max=255"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid task data")
		return
	}
	if err := r.validate.Struct(input); err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid task data")
		return
	}

	task, err := r.service.Start(input.UserId, input.Name)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			errorResponse(c, http.StatusBadRequest, "specified user not found")
		default:
			r.logger.Error("failed to start task", zap.Error(err))
			errorResponse(c, http.StatusInternalServerError, postgres.InternalServerError.Error())
		}
		return
	}

	c.JSON(200, gin.H{"message": "task started", "task": task})
}

func (r *taskRoutes) end(c *gin.Context) {
	var input struct {
		TaskId int64 `json:"task_id" binding:"required" validate:"gt=0"`
		UserId int64 `json:"user_id" binding:"required" validate:"gt=0"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid task data")
		return
	}
	if err := r.validate.Struct(input); err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid task data")
		return
	}

	task, err := r.service.End(input.TaskId, input.UserId)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			errorResponse(c, http.StatusBadRequest, "task not found")
		case errors.Is(err, postgres.RecordNotFound):
			errorResponse(c, http.StatusBadRequest, "task associated with provided user not found")
		default:
			r.logger.Error("failed to stop task", zap.Error(err))
			errorResponse(c, http.StatusInternalServerError, postgres.InternalServerError.Error())
		}
		return
	}

	c.JSON(200, gin.H{"message": "task ended", "task": task})
}

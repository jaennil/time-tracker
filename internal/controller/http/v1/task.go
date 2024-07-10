package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
	"github.com/jaennil/time-tracker/internal/model"
	"github.com/jaennil/time-tracker/internal/repository/postgres"
	"github.com/jaennil/time-tracker/internal/service"
	"github.com/jaennil/time-tracker/pkg/logger"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
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
		tasks.GET("/activity/:user_id", r.activity)
	}
}

// start task
//
//	@Summary		Start task
//	@Description	Start task with name for specified user by id
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Param			input	body		model.StartTask	true	"task data"
//	@Success		200				{object}	model.Task "task started"
//	@Failure		400				{object}	http.Response
//	@Failure		500				{object}	http.InternalServerErrorResponse
//	@Router			/tasks/start [post]
func (r *taskRoutes) start(c *gin.Context) {
	var input model.StartTask
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
			internalServerErrorResponse(c)
		}
		return
	}

	c.JSON(http.StatusOK, task)
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
			internalServerErrorResponse(c)
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task ended", "task": task})
}

func (r *taskRoutes) activity(c *gin.Context) {
	userIdStr := c.Param("user_id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid user id")
		return
	}
	err = r.validate.Var(userId, "gt=0")
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid user id")
		return
	}

	var input struct {
		StartTime time.Time `form:"start_time" binding:"required" time_format:"2006-01-02T15:04:05Z"`
		EndTime   time.Time `form:"end_time" binding:"required" time_format:"2006-01-02T15:04:05Z"`
	}
	if err := c.ShouldBindQuery(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, "start or end time not provided or have invalid format(2006-01-02T15:04:05Z)")
		return
	}

	activities, err := r.service.Activity(userId, input.StartTime, input.EndTime)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			errorResponse(c, http.StatusBadRequest, "activities not found")
		default:
			r.logger.Error("failed to get activities", zap.Error(err))
			internalServerErrorResponse(c)
		}
		return
	}

	c.JSON(http.StatusOK, activities)
}

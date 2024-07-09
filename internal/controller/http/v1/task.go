package http

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jaennil/time-tracker/internal/model"
	"github.com/jaennil/time-tracker/internal/repository/postgres"
	"github.com/jaennil/time-tracker/internal/service"
	"github.com/jaennil/time-tracker/pkg/logger"
	"go.uber.org/zap"
	"net/http"
	"strconv"
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
		tasks.POST(":user_id/start", r.start)
		tasks.POST("/end", r.end)
	}
}

func (r *taskRoutes) start(c *gin.Context) {
	userIdStr := c.Param("user_id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid user_id")
		return
	}

	task := model.Task{UserId: userId}
	if err := c.ShouldBind(&task); err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid task name")
		return
	}

	err = r.service.Start(&task)
	if err != nil {
		r.logger.Error("failed to start task", zap.Error(err))
		errorResponse(c, http.StatusInternalServerError, postgres.InternalServerError.Error())
		return
	}

	c.JSON(200, gin.H{"message": "task created", "task": task})
}

func (r *taskRoutes) end(c *gin.Context) {
	//userIdStr := c.Param("user_id")
	//userId, err := strconv.ParseInt(userIdStr, 10, 64)
	//if err != nil {
	//	errorResponse(c, http.StatusBadRequest, "invalid user_id")
	//	return
	//}
	//r.logger.Debug("user_id", zap.Int64("user_id", userId))
	//
	//taskIdStr := c.Param("task_id")
	//taskId, err := strconv.ParseInt(taskIdStr, 10, 64)
	//if err != nil {
	//	errorResponse(c, http.StatusBadRequest, "invalid user_id")
	//	return
	//}
	//r.logger.Debug("task_id", zap.Int64("task_id", taskId))

	var input struct {
		TaskId int64 `json:"task_id"`
	}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid task id")
		return
	}

	task, err := r.service.End(input.TaskId)
	if err != nil {
		// TODO: handle err no record
		r.logger.Error("failed to stop task", zap.Error(err))
		errorResponse(c, http.StatusInternalServerError, postgres.InternalServerError.Error())
		return
	}

	c.JSON(200, gin.H{"message": "task stopped", "task": task})
}

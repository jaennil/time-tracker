package http

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/jaennil/time-tracker/internal/model"
	"strconv"

	"github.com/jaennil/time-tracker/internal/repository/postgres"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jaennil/time-tracker/internal/service"
	"github.com/jaennil/time-tracker/pkg/logger"
)

type userRoutes struct {
	service  service.User
	logger   logger.Loggable
	validate *validator.Validate
}

func NewUserRoutes(handler *gin.RouterGroup, userService service.User, log logger.Loggable, validate *validator.Validate) {
	routes := &userRoutes{userService, log, validate}

	users := handler.Group("/users")
	{
		users.POST("", routes.create)
		users.DELETE(":id", routes.delete)
		users.PATCH(":id", routes.update)
		users.GET("", routes.get)
	}
}

func (r *userRoutes) create(c *gin.Context) {
	var input struct {
		Passport string `json:"passportNumber" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid passport data")
		return
	}
	err := r.validate.Var(input.Passport, "passport")
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid passport format")
		return
	}

	user, err := r.service.Create(input.Passport)
	if err != nil {
		r.logger.Error("failed to create user", err)
		errorResponse(c, http.StatusInternalServerError, postgres.InternalServerError.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created", "id": user.Id})
}

func (r *userRoutes) delete(c *gin.Context) {
	id, err := readIDParam(c)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}
	err = r.validate.Var(id, "gt=0")
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	err = r.service.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, postgres.RecordNotFound):
			errorResponse(c, http.StatusBadRequest, "user not found")
		default:
			r.logger.Error("failed to delete user", err)
			errorResponse(c, http.StatusInternalServerError, postgres.InternalServerError.Error())
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}

func (r *userRoutes) update(c *gin.Context) {
	id, err := readIDParam(c)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}
	err = r.validate.Var(id, "gt=0")
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	var user model.User
	if err = c.ShouldBindJSON(&user); err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid json user data")
		return
	}
	err = r.validate.Struct(user)
	if err != nil {
		r.logger.Error("failed to validate user data", err)
		errorResponse(c, http.StatusBadRequest, "invalid user data")
		return
	}

	err = r.service.Update(id, &user)
	if err != nil {
		switch {
		case errors.Is(err, postgres.RecordNotFound):
			errorResponse(c, http.StatusBadRequest, "user not found")
		default:
			r.logger.Error("failed to update user", err)
			errorResponse(c, http.StatusInternalServerError, postgres.InternalServerError.Error())
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user updated successfully"})
}

func (r *userRoutes) get(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid page number")
		return
	}
	err = r.validate.Var(page, "min=0")
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "page must be positive number")
		return
	}
	pageSizeStr := c.DefaultQuery("page_size", "100")
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid page size")
		return
	}
	err = r.validate.Var(pageSize, "min=0")
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "page size must be positive number")
		return
	}
	userIdStr := c.Query("user_id")
	var userId int
	if userIdStr == "" {
		userId = 0
	} else {
		userId, err = strconv.Atoi(userIdStr)
		if err != nil {
			errorResponse(c, http.StatusBadRequest, "invalid user id")
			return
		}
	}
	filter := &model.User{
		Id:             int64(userId),
		PassportSeries: c.Query("passport_series"),
		PassportNumber: c.Query("passport_number"),
		Name:           c.Query("name"),
		Surname:        c.Query("surname"),
		Patronymic:     c.Query("patronymic"),
		Address:        c.Query("address"),
	}

	pagination := &model.Pagination{Page: page, PageSize: pageSize}
	users, err := r.service.Get(pagination, filter)
	if err != nil {
		r.logger.Error("failed to get users", err)
		errorResponse(c, http.StatusInternalServerError, postgres.InternalServerError.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}

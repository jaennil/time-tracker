package http

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
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
		users.GET(":id", routes.getById)
	}
}

func (r *userRoutes) create(c *gin.Context) {
	var input struct {
		Passport string `json:"passportNumber" binding:"required" validate:"passport"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid or no passport data")
		return
	}
	err := r.validate.Struct(input)
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

	c.JSON(http.StatusCreated, gin.H{"message": "user created", "user": user})
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

	user, err := r.service.GetById(id)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			errorResponse(c, http.StatusBadRequest, "user not found")
		default:
			r.logger.Error("failed to update user", err)
			errorResponse(c, http.StatusInternalServerError, postgres.InternalServerError.Error())
		}
		return
	}

	if err = c.ShouldBindJSON(user); err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid json user data")
		return
	}
	err = r.validate.Struct(user)
	if err != nil {
		r.logger.Error("failed to validate user data", err)
		errorResponse(c, http.StatusBadRequest, "invalid user data")
		return
	}

	err = r.service.Update(id, user)
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

	c.JSON(http.StatusOK, gin.H{"message": "user updated successfully", "user": user})
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

	userIdStr := c.DefaultQuery("user_id", "0")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid user id")
		return
	}

	// TODO: validate filters
	patronymic := c.Query("patronymic")
	filter := model.User{
		Id:             userId,
		PassportSeries: c.Query("passport_series"),
		PassportNumber: c.Query("passport_number"),
		Name:           c.Query("name"),
		Surname:        c.Query("surname"),
		Patronymic:     &patronymic,
		Address:        c.Query("address"),
	}

	pagination := model.Pagination{Page: page, PageSize: pageSize}
	users, err := r.service.Get(&pagination, &filter)
	if err != nil {
		r.logger.Error("failed to get users", err)
		errorResponse(c, http.StatusInternalServerError, postgres.InternalServerError.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}

func (r *userRoutes) getById(c *gin.Context) {
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

	user, err := r.service.GetById(id)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			noContentResponse(c)
		default:
			r.logger.Error("failed to get user", err)
			errorResponse(c, http.StatusInternalServerError, postgres.InternalServerError.Error())
		}
		return
	}

	c.JSON(http.StatusOK, user)
}

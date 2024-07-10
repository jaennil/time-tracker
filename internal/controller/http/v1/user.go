package http

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
	"github.com/jaennil/time-tracker/internal/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jaennil/time-tracker/internal/repository/postgres"
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

// create user
//
//	@Summary		Create a user
//	@Description	Create user by passport number
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			passportNumber	body		model.CreateUser	true	"Full Passport Number"
//	@Success		200				{object}	model.User			"user created"
//	@Failure		400				{object}	http.Response
//	@Failure		500				{object}	http.InternalServerErrorResponse
//	@Router			/users [post]
func (r *userRoutes) create(c *gin.Context) {
	var input model.CreateUser
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
		switch {
		case errors.Is(err, service.UserAPIBadRequest):
			errorResponse(c, http.StatusBadRequest, "user not found")
		default:
			r.logger.Error("failed to create user", err)
			internalServerErrorResponse(c)
		}
		return
	}

	c.JSON(http.StatusCreated, user)
}

// delete user
//
//	@Summary		Delete a user
//	@Description	Delete user by id
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"User ID"
//	@Success		204
//	@Failure		400	{object}	http.Response
//	@Failure		500	{object}	http.InternalServerErrorResponse
//	@Router			/users/{id} [delete]
//
// TODO: maybe delete produce json because return code is 204
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
			internalServerErrorResponse(c)
		}
		return
	}

	noContentResponse(c)
}

// update user
//
//	@Summary		Update a user
//	@Description	Update all or several user fields
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int			true	"User ID"	example(1)	minimum(1)
//	@Param			user	body		model.User	true	"User"
//	@Success		200		{object}	model.User	"user updated"
//	@Failure		400		{object}	http.Response
//	@Failure		500		{object}	http.InternalServerErrorResponse
//	@Router			/users/{id} [patch]
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
			internalServerErrorResponse(c)
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
			internalServerErrorResponse(c)
		}
		return
	}

	c.JSON(http.StatusOK, user)
}

// get users
//
//	@Summary		Get users
//	@Description	Retrieve users info with filtering and pagination support
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			pagination	query		model.Pagination	false	"Pagination"
//	@Param			filter		query		model.User			false	"Filter"
//	@Success		200			{array}		model.User
//	@Failure		400			{object}	http.Response
//	@Failure		500			{object}	http.InternalServerErrorResponse
//	@Router			/users [get]
func (r *userRoutes) get(c *gin.Context) {
	var pagination model.Pagination
	if err := c.ShouldBindQuery(&pagination); err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid pagination data")
		return
	}
	if err := r.validate.Struct(pagination); err != nil {
		errorResponse(c, http.StatusBadRequest, "invalid pagination data")
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

	users, err := r.service.Get(&pagination, &filter)
	if err != nil {
		r.logger.Error("failed to get users", err)
		internalServerErrorResponse(c)
		return
	}

	c.JSON(http.StatusOK, users)
}

// getById
//
//	@Summary		Get user
//	@Description	Get user by id
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"User ID"	example(1)	minimum(1)
//	@Success		200	{object}	model.User
//	@Success		204
//	@Failure		400	{object}	http.Response
//	@Failure		500	{object}	http.InternalServerErrorResponse
//	@Router			/users/{id} [get]
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
			internalServerErrorResponse(c)
		}
		return
	}

	c.JSON(http.StatusOK, user)
}

package repository

import (
	"github.com/jackc/pgx/v5"
	"github.com/jaennil/time-tracker/internal/model"
	"github.com/jaennil/time-tracker/internal/repository/postgres"
)

type Repository struct {
	User
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		User: postgres.NewUserRepository(db),
	}
}

type User interface {
	Store(*model.User) error
}

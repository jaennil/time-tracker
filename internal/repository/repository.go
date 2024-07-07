package repository

import (
	"github.com/jackc/pgx/v5"
	"github.com/jaennil/time-tracker/internal/repository/postgres"
)

type Repository struct {
	UserRepository
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository {
		UserRepository: postgres.NewUserRepository(db),
	}
}

type UserRepository interface {
	Create()
}

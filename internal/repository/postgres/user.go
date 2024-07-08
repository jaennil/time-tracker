package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jaennil/time-tracker/internal/model"
	"time"
)

type UserRepository struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Store(user *model.User) error {
	query := `INSERT INTO users (name, surname, patronymic, address, passport_series, passport_number)
			VALUES ($1, $2, $3, $4, $5, $6)
			RETURNING user_id`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return r.db.QueryRow(ctx, query, user.Name, user.Surname, user.Patronymic, user.Address, user.PassportSeries, user.PassportNumber).
		Scan(&user.Id)
}

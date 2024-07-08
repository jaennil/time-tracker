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

func (r *UserRepository) Delete(id int64) error {
	query := `DELETE FROM users WHERE user_id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return RecordNotFound
	}

	return nil
}

func (r *UserRepository) Update(id int64, user *model.User) error {
	query := `UPDATE users
				SET name=$1, surname=$2, patronymic=$3, address=$4, passport_series=$5, passport_number=$6
				WHERE user_id = $7`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := r.db.Exec(ctx, query, user.Name, user.Surname, user.Patronymic, user.Address, user.PassportSeries, user.PassportNumber, id)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return RecordNotFound
	}

	return nil
}

func (r *UserRepository) Get() ([]model.User, error) {
	query := `SELECT user_id, name, surname, patronymic, address, passport_series, passport_number FROM users`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.User])
	if err != nil {
		return nil, err
	}

	return users, nil
}

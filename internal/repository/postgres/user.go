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
				SET name = $1, surname = $2, patronymic = $3, address = $4, passport_series = $5, passport_number = $6
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

func (r *UserRepository) Get(pagination *model.Pagination, filter *model.User) ([]model.User, error) {
	query := `SELECT user_id, name, surname, patronymic, address, passport_series, passport_number
				FROM users
				WHERE (user_id = $1 OR $1 = 0) AND
					(name = $2 OR $2 = '') AND
					(surname = $3 OR $3 = '') AND
					(patronymic = $4 OR $4 = '') AND
					(address = $5 OR $5 = '') AND
					(passport_series = $6 OR $6 = '') AND
					(passport_number = $7 OR $7 = '')
				ORDER BY user_id
				LIMIT $8 OFFSET $9`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	offset := (pagination.Page - 1) * pagination.PageSize
	rows, err := r.db.Query(ctx,
		query,
		filter.Id,
		filter.Name,
		filter.Surname,
		filter.Patronymic,
		filter.Address,
		filter.PassportSeries,
		filter.PassportNumber,
		pagination.PageSize,
		offset,
	)
	if err != nil {
		return nil, err
	}

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.User])
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) GetById(id int64) (*model.User, error) {
	query := `SELECT user_id, name, surname, patronymic, address, passport_series, passport_number
				FROM users
				WHERE user_id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := r.db.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}

	user, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[model.User])
	if err != nil {
		return nil, err
	}

	return &user, nil
}

package repository

import (
	"github.com/jackc/pgx/v5"
	"github.com/jaennil/time-tracker/internal/model"
	"github.com/jaennil/time-tracker/internal/repository/postgres"
	"time"
)

type Repository struct {
	User
	Task
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		User: postgres.NewUserRepository(db),
		Task: postgres.NewTaskRepository(db),
	}
}

type User interface {
	Store(user *model.User) error
	Delete(id int64) error
	Update(id int64, user *model.User) error
	Get(pagination *model.Pagination, filter *model.User) ([]model.User, error)
	GetById(id int64) (*model.User, error)
}

type Task interface {
	Store(task *model.Task) error
	End(task *model.Task) error
	GetById(id int64) (*model.Task, error)
	Activity(userId int64, startTime, endTime time.Time) ([]model.Activity, error)
}

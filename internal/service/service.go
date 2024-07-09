package service

import (
	"github.com/jaennil/time-tracker/internal/model"
	"github.com/jaennil/time-tracker/internal/repository"
)

type Service struct {
	User
	Task
}

func New(repositories *repository.Repository, userApi *UserAPI) *Service {
	return &Service{
		User: NewUserService(repositories.User, userApi),
		Task: NewTaskService(repositories.Task, repositories.User),
	}
}

type User interface {
	Create(passport string) (*model.User, error)
	Delete(id int64) error
	Update(id int64, user *model.User) error
	Get(pagination *model.Pagination, filter *model.User) ([]model.User, error)
	GetById(id int64) (*model.User, error)
}

type Task interface {
	Start(userId int64, name string) (*model.Task, error)
	End(taskId int64, userId int64) (*model.Task, error)
}

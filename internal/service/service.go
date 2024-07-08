package service

import (
	"github.com/jaennil/time-tracker/internal/model"
	"github.com/jaennil/time-tracker/internal/repository"
)

type Service struct {
	User
}

func New(repositories *repository.Repository, userApi *UserAPI) *Service {
	return &Service{
		User: NewUserService(repositories, userApi),
	}
}

type User interface {
	Create(passport string) (*model.User, error)
	Delete(id int64) error
	Update(id int64, user *model.User) error
}

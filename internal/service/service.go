package service

import "github.com/jaennil/time-tracker/internal/repository"

type Service struct {
	User
}

func New(repositories *repository.Repository) *Service {
	return &Service {
		User: NewUserService(repositories),
	}
}

type User interface {
	Create()
}

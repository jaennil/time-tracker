package service

import "github.com/jaennil/time-tracker/internal/repository"

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *UserService {
	return &UserService{repository}
}

func (s *UserService) Create() {
}

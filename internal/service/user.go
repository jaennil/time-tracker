package service

import (
	"errors"
	"github.com/jaennil/time-tracker/internal/model"
	"github.com/jaennil/time-tracker/internal/repository"
	"strings"
)

type UserService struct {
	userRepository repository.User
	userApi        *UserAPI
}

func NewUserService(repository repository.User, userApi *UserAPI) *UserService {
	return &UserService{repository, userApi}
}

func (s *UserService) Create(passport string) (*model.User, error) {
	seriesAndNumber := strings.Split(passport, " ")
	if len(seriesAndNumber) != 2 {
		return nil, errors.New("invalid passport format")
	}

	series := seriesAndNumber[0]
	number := seriesAndNumber[1]
	user, err := s.userApi.UserInfo(series, number)
	if err != nil {
		return nil, err
	}

	user.PassportNumber = number
	user.PassportSeries = series

	err = s.userRepository.Store(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Delete(id int64) error {
	return s.userRepository.Delete(id)
}

package service

import (
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

	series := seriesAndNumber[0]
	number := seriesAndNumber[1]
	user, err := s.userApi.UserInfo(series, number)
	if err != nil {
		return nil, err
	}
	// TODO: maybe validate user api response values

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

func (s *UserService) Update(id int64, user *model.User) error {
	return s.userRepository.Update(id, user)
}

func (s *UserService) Get(pagination *model.Pagination, filter *model.User) ([]model.User, error) {
	return s.userRepository.Get(pagination, filter)
}

func (s *UserService) GetById(id int64) (*model.User, error) {
	return s.userRepository.GetById(id)
}

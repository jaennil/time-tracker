package service

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jaennil/time-tracker/internal/model"
	"github.com/jaennil/time-tracker/internal/repository"
	"time"
)

type TaskService struct {
	taskRepository repository.Task
	userRepository repository.User
}

func NewTaskService(taskRepository repository.Task, userRepository repository.User) *TaskService {
	return &TaskService{taskRepository, userRepository}
}

func (s *TaskService) Start(userId int64, name string) (*model.Task, error) {
	// verify that provided user exists by userId
	_, err := s.userRepository.GetById(userId)
	if err != nil {
		return nil, err
	}

	task := model.Task{
		StartTime: pgtype.Timestamptz{Time: time.Now(), Valid: true},
		UserId:    userId,
		Name:      name,
	}

	err = s.taskRepository.Store(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (s *TaskService) End(taskId, userId int64) (*model.Task, error) {
	task, err := s.taskRepository.GetById(taskId)
	if err != nil {
		return nil, err
	}

	task.UserId = userId
	task.EndTime = pgtype.Timestamptz{Time: time.Now(), Valid: true}

	err = s.taskRepository.End(task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

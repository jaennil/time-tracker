package service

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jaennil/time-tracker/internal/model"
	"github.com/jaennil/time-tracker/internal/repository"
	"time"
)

type TaskService struct {
	taskRepository repository.Task
}

func NewTaskService(repository repository.Task) *TaskService {
	return &TaskService{repository}
}

func (s *TaskService) Start(task *model.Task) error {
	task.StartTime = pgtype.Timestamptz{Time: time.Now(), Valid: true}

	err := s.taskRepository.Store(task)
	if err != nil {
		return err
	}

	return nil
}

func (s *TaskService) End(id int64) (*model.Task, error) {
	task, err := s.taskRepository.GetById(id)
	if err != nil {
		// TODO: handle err no records
		return nil, err
	}

	task.EndTime = pgtype.Timestamptz{Time: time.Now(), Valid: true}

	err = s.taskRepository.Update(task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

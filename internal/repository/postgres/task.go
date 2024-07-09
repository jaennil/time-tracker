package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jaennil/time-tracker/internal/model"
	"time"
)

type TaskRepository struct {
	db *pgx.Conn
}

func NewTaskRepository(db *pgx.Conn) *TaskRepository {
	return &TaskRepository{db}
}

func (r *TaskRepository) Store(task *model.Task) error {
	query := `INSERT INTO tasks (user_id, name, start_time)
				VALUES ($1, $2, $3) RETURNING task_id`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return r.db.QueryRow(ctx, query, task.UserId, task.Name, task.StartTime.Time).
		Scan(&task.TaskId)
}

func (r *TaskRepository) End(task *model.Task) error {
	query := `UPDATE tasks
				SET end_time = $1
				WHERE task_id = $2 AND user_id = $3`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := r.db.Exec(ctx, query, task.EndTime.Time, task.TaskId, task.UserId)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return RecordNotFound
	}

	return nil
}

func (r *TaskRepository) GetById(id int64) (*model.Task, error) {
	query := `SELECT task_id, user_id, name, start_time, end_time FROM tasks WHERE task_id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := r.db.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}

	task, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[model.Task])
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *TaskRepository) Activity(userId int64, startTime, endTime time.Time) ([]model.Activity, error) {
	query := `SELECT name, end_time - start_time AS duration
				FROM tasks
				WHERE user_id = $1 AND end_time IS NOT NULL AND start_time >= $2 AND end_time <= $3
				ORDER BY duration DESC;`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := r.db.Query(ctx, query, userId, startTime, endTime)
	if err != nil {
		return nil, err
	}

	activities, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.Activity])
	if err != nil {
		return nil, err
	}

	return activities, nil
}

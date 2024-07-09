package model

import "github.com/jackc/pgx/v5/pgtype"

type Task struct {
	TaskId    int64              `db:"task_id" json:"task_id"`
	UserId    int64              `db:"user_id"`
	Name      string             `db:"name" json:"name"`
	StartTime pgtype.Timestamptz `db:"start_time"`
	EndTime   pgtype.Timestamptz `db:"end_time"`
}

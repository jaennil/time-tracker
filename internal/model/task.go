package model

import (
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type Task struct {
	TaskId    int64              `db:"task_id" json:"task_id"`
	UserId    int64              `db:"user_id" json:"user_id"`
	Name      string             `db:"name" json:"name"`
	StartTime pgtype.Timestamptz `db:"start_time" json:"start_time"`
	EndTime   pgtype.Timestamptz `db:"end_time" json:"end_time"`
}

type Activity struct {
	Name     string        `db:"name"`
	Duration time.Duration `db:"duration"`
}

type PrettyActivity struct {
	Name    string  `json:"name"`
	Hours   int     `json:"hours"`
	Minutes float64 `json:"minutes"`
}

package model

import (
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type Task struct {
	TaskId    int64              `db:"task_id" json:"task_id" example:"1" minimum:"1"`
	UserId    int64              `db:"user_id" json:"user_id" example:"1" minimum:"1"`
	Name      string             `db:"name" json:"name" example:"do stuff"`
	StartTime pgtype.Timestamptz `db:"start_time" json:"start_time" swaggertype:"string" example:"2024-07-10T07:00:43.047939731+03:00"`
	EndTime   pgtype.Timestamptz `db:"end_time" json:"end_time" swaggertype:"string" example:"2025-07-10T07:00:43.047939731+03:00"`
}

type StartTask struct {
	UserId int64  `json:"user_id" binding:"required" validate:"gt=0" example:"1" minimum:"1"`
	Name   string `json:"name" binding:"required" validate:"min=1,max=255" example:"do stuff"`
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

package model

import "time"

type Task struct {
	TaskId    int
	UserId    int
	Name      string
	StartTime time.Time
	EndTime   time.Time
}

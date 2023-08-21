package model

import "time"

type Task struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	Body      string    `gorm:"column:body"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

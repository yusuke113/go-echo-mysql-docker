package repository

import (
	"go/model"

	"gorm.io/gorm"
)

type ITaskRepository interface {
	GetList(tasks *[]model.Task) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &taskRepository{db}
}

func (ur *taskRepository) GetList(tasks *[]model.Task) error {

	if err := ur.db.Find(&tasks).Error; err != nil {
		return err
	}

	return nil
}

package usecase

import (
	"go/model"
	"go/repository"
	"time"
)

type GetListResponse struct {
	ID        int    `json:"id"`
	Body      string `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

type ITaskUseCase interface {
	GetList() ([]GetListResponse, error)
}

type taskUseCase struct {
	ur repository.ITaskRepository
}

func NewTaskUseCase(ur repository.ITaskRepository) ITaskUseCase {
	return &taskUseCase{ur}
}

func (uc *taskUseCase) GetList() ([]GetListResponse, error) {
	tasks := []model.Task{}

	if err := uc.ur.GetList(&tasks); err != nil {
		return nil, err
	}

	res := []GetListResponse{}
	for _, v := range tasks {
		res = append(res, GetListResponse{
			ID:        v.ID,
			Body:      v.Body,
			CreatedAt: v.CreatedAt,
		})
	}

	return res, nil
}

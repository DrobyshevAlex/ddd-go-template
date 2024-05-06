package local

import (
	"context"

	"github.com/DrobyshevAlex/ddd-go-template/internal/domain"
	domainErrors "github.com/DrobyshevAlex/ddd-go-template/internal/domain/errors"
)

func (s Service) SaveTask(ctx context.Context, task *domain.Task) (domain.TaskID, error) {
	return domain.TaskID(1), nil
}

type Task struct {
	ID        uint64
	Name      string
	CreatorID uint64
}

func (s Service) GetTask(ctx context.Context, taskID domain.TaskID) (Task, error) {
	switch taskID {
	case 1:
		return Task{ID: uint64(taskID), Name: "Task 1", CreatorID: 1}, nil
	case 2:
		return Task{ID: uint64(taskID), Name: "Task 2", CreatorID: 2}, nil
	case 3:
		return Task{ID: uint64(taskID), Name: "Task 3", CreatorID: 3}, nil
	}

	return Task{}, domainErrors.NotFound("задача не найдена")
}

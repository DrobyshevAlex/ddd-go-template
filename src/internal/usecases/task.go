package usecases

import (
	"context"

	"github.com/DrobyshevAlex/ddd-go-template/internal/domain"
	"github.com/DrobyshevAlex/ddd-go-template/internal/domain/dto"
	domainErrors "github.com/DrobyshevAlex/ddd-go-template/internal/domain/errors"
)

type TaskRepo interface {
	GetUser(ctx context.Context, id domain.UserID) (domain.User, error)
	SaveTask(ctx context.Context, task *domain.Task) (domain.TaskID, error)
	GetTask(ctx context.Context, id domain.TaskID) (*domain.Task, error)
}

func TaskCreate(ctx context.Context, r TaskRepo, params dto.TaskCreate) (*domain.Task, error) {
	user, err := r.GetUser(ctx, domain.UserID(params.Session.UserID))
	if err != nil && !domainErrors.IsNotFound(err) {
		// log err
		return nil, domainErrors.Validation("произошла ошибка, повторите позже")
	}

	task, err := domain.NewTask(params, user)
	if err != nil {
		return nil, err
	}
	taskID, err := r.SaveTask(ctx, task)
	if err != nil {
		return nil, err
	}

	task.SetID(taskID)
	return task, nil
}

func GetTask(ctx context.Context, r TaskRepo, params dto.TaskFilter) (*domain.Task, error) {
	return r.GetTask(ctx, domain.TaskID(params.TaskID))
}

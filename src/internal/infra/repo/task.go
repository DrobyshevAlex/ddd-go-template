package repo

import (
	"context"

	"github.com/DrobyshevAlex/ddd-go-template/internal/domain"
	"github.com/DrobyshevAlex/ddd-go-template/internal/infra/repo/adapters"
)

func (r Repository) SaveTask(ctx context.Context, task *domain.Task) (domain.TaskID, error) {
	return r.db.SaveTask(ctx, task)
}

func (r Repository) GetTask(ctx context.Context, taskID domain.TaskID) (*domain.Task, error) {
	task, err := r.db.GetTask(ctx, taskID)
	if err != nil {
		return nil, err
	}

	user, err := r.user.GetUserByID(ctx, task.CreatorID)
	if err != nil {
		// log
	}

	return adapters.ToTask(task, user), nil
}

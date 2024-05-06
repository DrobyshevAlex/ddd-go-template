package adapters

import (
	"github.com/DrobyshevAlex/ddd-go-template/internal/domain"
	"github.com/DrobyshevAlex/ddd-go-template/internal/infra/repo/local"
	"github.com/DrobyshevAlex/ddd-go-template/internal/infra/repo/user"
)

func ToTask(task local.Task, user user.User) *domain.Task {
	return &domain.Task{
		ID:   domain.TaskID(task.ID),
		Name: domain.TaskName(task.Name),
		Creator: domain.User{
			ID:   domain.UserID(task.CreatorID),
			Name: user.Name,
		},
	}
}

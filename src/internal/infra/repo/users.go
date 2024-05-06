package repo

import (
	"context"

	"github.com/DrobyshevAlex/ddd-go-template/internal/domain"
	"github.com/DrobyshevAlex/ddd-go-template/internal/infra/repo/adapters"
)

func (r Repository) GetUser(ctx context.Context, userID domain.UserID) (domain.User, error) {
	user, err := r.user.GetUserByID(ctx, uint64(userID))
	if err != nil {
		return domain.User{}, err
	}
	return adapters.ToUser(user), nil
}

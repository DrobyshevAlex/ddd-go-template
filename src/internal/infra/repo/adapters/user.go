package adapters

import (
	"github.com/DrobyshevAlex/ddd-go-template/internal/domain"
	"github.com/DrobyshevAlex/ddd-go-template/internal/infra/repo/user"
)

func ToUser(u user.User) domain.User {
	return domain.User{ID: domain.UserID(u.ID), Name: u.Name}
}

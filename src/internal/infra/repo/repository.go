package repo

import (
	"github.com/DrobyshevAlex/ddd-go-template/internal/infra/repo/local"
	"github.com/DrobyshevAlex/ddd-go-template/internal/infra/repo/user"
)

type Repository struct {
	db   *local.Service
	user *user.Service
}

func New(db *local.Service, user *user.Service) (*Repository, error) {
	return &Repository{db: db, user: user}, nil
}

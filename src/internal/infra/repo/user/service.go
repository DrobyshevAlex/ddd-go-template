package user

import (
	"context"
	"errors"

	domainErrors "github.com/DrobyshevAlex/ddd-go-template/internal/domain/errors"
)

type Service struct {
	// apiClient
}

func New() (*Service, error) {
	return &Service{}, nil
}

type User struct {
	ID   uint64
	Name string
}

func (s Service) GetUserByID(ctx context.Context, userID uint64) (User, error) {
	// res, err := s.apiClient.GetUserByID(ctx, userID)
	if userID == 2 {
		return User{}, domainErrors.NotFound("не найден")
	} else if userID == 3 {
		return User{}, errors.New("error")
	}
	return User{
		ID:   userID,
		Name: "Username",
	}, nil
}

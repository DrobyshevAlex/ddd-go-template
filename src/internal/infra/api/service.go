package api

import (
	"github.com/DrobyshevAlex/ddd-go-template/internal/gen/pb"
	"github.com/DrobyshevAlex/ddd-go-template/internal/infra/repo"
)

type Service struct {
	pb.UnimplementedTaskServiceServer

	r *repo.Repository
}

func New(r *repo.Repository) *Service {
	return &Service{r: r}
}

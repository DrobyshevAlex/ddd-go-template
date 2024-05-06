package api

import (
	"context"

	"github.com/DrobyshevAlex/ddd-go-template/internal/gen/pb"
	"github.com/DrobyshevAlex/ddd-go-template/internal/infra/api/adapters"
	"github.com/DrobyshevAlex/ddd-go-template/internal/usecases"
)

func (s *Service) TaskCreate(ctx context.Context, req *pb.TaskCreate_Request) (*pb.TaskCreate_Response, error) {
	task, err := usecases.TaskCreate(ctx, s.r, adapters.FromTaskCreateRequest(req))
	if err != nil {
		return nil, err
	}
	return adapters.ToTaskCreateResponse(task), nil
}

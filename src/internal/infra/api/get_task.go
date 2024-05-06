package api

import (
	"context"

	"github.com/DrobyshevAlex/ddd-go-template/internal/gen/pb"
	"github.com/DrobyshevAlex/ddd-go-template/internal/infra/api/adapters"
	"github.com/DrobyshevAlex/ddd-go-template/internal/usecases"
)

func (s *Service) GetTask(ctx context.Context, req *pb.GetTask_Request) (*pb.GetTask_Response, error) {
	task, err := usecases.GetTask(ctx, s.r, adapters.FromGetTaskRequest(req))
	if err != nil {
		return nil, err
	}
	return adapters.ToGetTaskResponse(task), nil
}

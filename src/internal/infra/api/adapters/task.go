package adapters

import (
	"github.com/DrobyshevAlex/ddd-go-template/internal/domain"
	"github.com/DrobyshevAlex/ddd-go-template/internal/domain/dto"
	"github.com/DrobyshevAlex/ddd-go-template/internal/gen/pb"
)

func FromTaskCreateRequest(req *pb.TaskCreate_Request) dto.TaskCreate {
	return dto.TaskCreate{
		Name: req.GetData().GetName(),
		Session: dto.Session{
			UserID: req.GetSession().GetUserId(),
		},
	}
}

func ToTaskCreateResponse(task *domain.Task) *pb.TaskCreate_Response {
	return &pb.TaskCreate_Response{
		Id:       uint64(task.ID),
		Name:     string(task.Name),
		UserId:   uint64(task.Creator.ID),
		Username: &task.Creator.Name,
	}
}

func FromGetTaskRequest(req *pb.GetTask_Request) dto.TaskFilter {
	return dto.TaskFilter{
		TaskID: req.GetData().GetId(),
		Session: dto.Session{
			UserID: req.GetSession().GetUserId(),
		},
	}
}

func ToGetTaskResponse(task *domain.Task) *pb.GetTask_Response {
	return &pb.GetTask_Response{
		Id:       uint64(task.ID),
		Name:     string(task.Name),
		UserId:   uint64(task.Creator.ID),
		Username: &task.Creator.Name,
	}
}

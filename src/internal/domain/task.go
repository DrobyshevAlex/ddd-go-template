package domain

import (
	"fmt"
	"strings"

	"github.com/DrobyshevAlex/ddd-go-template/internal/domain/dto"
	domainErrors "github.com/DrobyshevAlex/ddd-go-template/internal/domain/errors"
)

type TaskID uint64

type TaskName string

const taskNameMinLen = 5

func NewTaskName(v string) (TaskName, error) {
	v = strings.TrimSpace(v)
	if len(v) < taskNameMinLen {
		return TaskName(""), domainErrors.Validation(fmt.Sprintf("имя не может быть меньше %d символов", taskNameMinLen))
	}
	return TaskName(v), nil
}

type Task struct {
	ID      TaskID
	Name    TaskName
	Creator User
}

func NewTask(params dto.TaskCreate, user User) (*Task, error) {
	if user.ID == 0 {
		return nil, domainErrors.Validation("пользователь не найден")
	}
	name, err := NewTaskName(params.Name)
	if err != nil {
		return nil, err
	}
	return &Task{
		Name:    name,
		Creator: user,
	}, nil
}

func (t *Task) SetID(taskID TaskID) {
	t.ID = taskID
}

func (t *Task) SetCreator(user User) {
	t.Creator = user
}

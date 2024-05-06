package dto

type TaskCreate struct {
	Session Session
	Name    string
}

type TaskFilter struct {
	Session Session
	TaskID  uint64
}

package local

type Service struct {
}

func New() (*Service, error) {
	return &Service{}, nil
}

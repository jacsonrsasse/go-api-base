package something

import "context"

type Service interface {
	ListSomething(ctx context.Context) ([]string, error)
}

type svc struct {
}

func NewService() Service {
	return &svc{}
}

func (s *svc) ListSomething(ctx context.Context) ([]string, error) {
	return []string{"something"}, nil
}

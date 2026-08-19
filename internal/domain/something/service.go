package something

import (
	"context"

	repository "github.com/jacsonrsasse/go-api-base/internal/adapters/postgresql/sqlc"
)

type Service interface {
	ListSomething(ctx context.Context) ([]repository.Something, error)
}

type svc struct {
	repo repository.Querier
}

func NewService(repo repository.Querier) Service {
	return &svc{
		repo,
	}
}

func (s *svc) ListSomething(ctx context.Context) ([]repository.Something, error) {
	return s.repo.ListSomething(ctx)
}

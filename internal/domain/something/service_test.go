package something

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	repository "github.com/jacsonrsasse/go-api-base/internal/adapters/postgresql/sqlc"
)

// 1. Defina a struct do mock
type mockQuerier struct {
	// Você pode embutir a interface para não precisar implementar 
	// TODOS os métodos do sqlc caso a interface Querier seja gigante.
	repository.Querier 
	
	// Campos para controlar o comportamento do mock
	listSomethingFn func(ctx context.Context) ([]repository.Something, error)
}

// 2. Implemente o método que o seu serviço chama
func (m *mockQuerier) ListSomething(ctx context.Context) ([]repository.Something, error) {
	return m.listSomethingFn(ctx)
}

func TestSvc_ListSomething(t *testing.T) {
	now := time.Now()
	tz := pgtype.Timestamptz{
		Time: now,
		Valid: true,
	}
	something := repository.Something{ID: 1, Name: "Something", CreatedAt: tz, UpdatedAt: tz}

	// 3. Instancie o mock definindo o comportamento esperado
	mockRepo := &mockQuerier{
		listSomethingFn: func(ctx context.Context) ([]repository.Something, error) {
			return []repository.Something{
				something, 
			}, nil
		},
	}

	s := NewService(mockRepo)

	got, err := s.ListSomething(context.Background())
	if err != nil {
		t.Fatalf("ListSomething returned unexpected error: %v", err)
	}

	want := []repository.Something{something}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ListSomething = %v, want %v", got, want)
	}
}

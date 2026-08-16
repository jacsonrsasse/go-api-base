package something

import (
	"context"
	"reflect"
	"testing"
)

func TestSvc_ListSomething(t *testing.T) {
	s := NewService()

	got, err := s.ListSomething(context.Background())
	if err != nil {
		t.Fatalf("ListSomething returned unexpected error: %v", err)
	}

	want := []string{"something"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ListSomething = %v, want %v", got, want)
	}
}

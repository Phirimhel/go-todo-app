package core_http_types

import (
	"encoding/json"
	"fmt"

	"github.com/Phirimhel/go-todo-app/internal/core/domain"
)

type Nullable[T any] struct {
	domain.Nullable[T]
}

func (n *Nullable[T]) UnmarshalJSON(b []byte) error {
	n.Set = true

	if string(b) == "null" {
		n.Value = nil
		return nil
	}

	var value T
	if err := json.Unmarshal(b, &value); err != nil {
		return fmt.Errorf("[nullable unmarshal]: failed to unmarshal json %w", err)
	}

	n.Value = &value

	return nil
}

func (p *Nullable[T]) Domain() domain.Nullable[T] {
	return domain.Nullable[T]{}
}

package usecase

import (
	"context"
	"github.com/gictorbit/ice/internal/todo/domain"
)

// CreateTodo creates a new TodoItem and publishes it to the stream
func (uc *TodoUseCase) CreateTodo(ctx context.Context, todo *domain.TodoItem) error {
	if err := uc.repo.CreateTodo(ctx, todo); err != nil {
		return err
	}
	return uc.stream.Publish(ctx, todo)
}

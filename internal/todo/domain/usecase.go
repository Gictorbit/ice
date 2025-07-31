package domain

import "context"

type TodoUseCase interface {
	CreateTodo(ctx context.Context, todo *TodoItem) error
}

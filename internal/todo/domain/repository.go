package domain

import "context"

type TodoRepository interface {
	CreateTodo(ctx context.Context, todo *TodoItem) error
}

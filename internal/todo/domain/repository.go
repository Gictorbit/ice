package domain

import "context"

//go:generate mockgen -source=$GOFILE -destination=../../../tests/mocks/todorepo_mock.go -package=$GOPACKAG
type TodoRepository interface {
	CreateTodo(ctx context.Context, todo *TodoItem) error
}

package mysql

import (
	"context"
	"github.com/gictorbit/ice/internal/todo/domain"
	"github.com/google/uuid"
)

const CreateTodoSQL = `INSERT INTO todos (id, description, due_date) VALUES (?,?, ?, ?)`

// CreateTodo creates a new TodoItem in mysql
func (m *TodoMysql) CreateTodo(ctx context.Context, todo *domain.TodoItem) error {
	if todo.ID == "" {
		todo.ID = uuid.NewString()
	}
	_, err := m.db.ExecContext(ctx, CreateTodoSQL, todo.ID, todo.Description, todo.DueDate)
	return err
}

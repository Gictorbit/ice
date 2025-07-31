package mysql

import (
	"context"
	"github.com/gictorbit/ice/internal/todo/domain"
)

const CreateTodoSQL = `INSERT INTO todos (id, description, due_date) VALUES (?,?,?)`

// CreateTodo creates a new TodoItem in mysql
func (m *TodoMysql) CreateTodo(ctx context.Context, todo *domain.TodoItem) error {
	_, err := m.db.ExecContext(ctx, CreateTodoSQL, todo.ID.String(), todo.Description, todo.DueDate)
	return err
}

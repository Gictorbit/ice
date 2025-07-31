package domain

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type TodoItem struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"dueDate"`
}

func (t *TodoItem) GetTableName() string {
	return "todos"
}

func (t *TodoItem) Validate() error {
	if t.Description == "" {
		return errors.New("description cannot be empty")
	}
	if t.DueDate.IsZero() {
		return errors.New("dueDate must be set")
	}
	return nil
}

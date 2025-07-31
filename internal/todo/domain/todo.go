package domain

type TodoItem struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
}

func (t *TodoItem) GetTableName() string {
	return "todos"
}

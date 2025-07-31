package delivery

import (
	"context"
	"encoding/json"
	"github.com/gictorbit/ice/internal/todo/domain"
	"github.com/google/uuid"
	"net/http"
)

type TodoHandler struct {
	uc domain.TodoUseCase
}

func NewTodoHandler(uc domain.TodoUseCase) *TodoHandler {
	return &TodoHandler{uc: uc}
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	todo := &domain.TodoItem{}
	if err := json.NewDecoder(r.Body).Decode(todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	todo.ID = uuid.NewString()
	if err := h.uc.CreateTodo(context.Background(), todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

package delivery

import (
	"context"
	"encoding/json"
	"github.com/gictorbit/ice/internal/todo/domain"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/http"
)

type TodoHandler struct {
	uc     domain.TodoUseCase
	logger *zap.Logger
}

func NewTodoHandler(uc domain.TodoUseCase, logger *zap.Logger) *TodoHandler {
	return &TodoHandler{
		uc:     uc,
		logger: logger,
	}
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	todo := &domain.TodoItem{}
	if err := json.NewDecoder(r.Body).Decode(todo); err != nil {
		http.Error(w, "invalid request payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := todo.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todo.ID = uuid.New()

	if err := h.uc.CreateTodo(context.Background(), todo); err != nil {
		h.logger.Error("failed to create todo",
			zap.Error(err),
			zap.String("todoId", todo.ID.String()),
		)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

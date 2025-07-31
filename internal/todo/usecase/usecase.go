package usecase

import (
	"github.com/gictorbit/ice/internal/todo/domain"
	"go.uber.org/zap"
)

type TodoUseCase struct {
	repo   domain.TodoRepository
	stream domain.StreamPublisher
	logger *zap.Logger
}

// NewTodoUseCase creates a new instance of TodoUseCase
func NewTodoUseCase(repo domain.TodoRepository, streamer domain.StreamPublisher, logger *zap.Logger) domain.TodoUseCase {
	return &TodoUseCase{
		repo:   repo,
		stream: streamer,
		logger: logger,
	}
}

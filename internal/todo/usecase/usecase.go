package usecase

import "github.com/gictorbit/ice/internal/todo/domain"

type TodoUseCase struct {
	repo   domain.TodoRepository
	stream domain.StreamPublisher
}

// NewTodoUseCase creates a new instance of TodoUseCase
func NewTodoUseCase(repo domain.TodoRepository, streamer domain.StreamPublisher) domain.TodoUseCase {
	return &TodoUseCase{
		repo:   repo,
		stream: streamer,
	}
}

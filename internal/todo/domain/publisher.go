package domain

import "context"

type StreamPublisher interface {
	Publish(ctx context.Context, todo *TodoItem) error
}

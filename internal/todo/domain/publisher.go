package domain

import "context"

//go:generate mockgen -source=$GOFILE -destination=../mocks/stream_publisher_mock.go -package=$GOPACKAG
type StreamPublisher interface {
	Publish(ctx context.Context, todo *TodoItem) error
}

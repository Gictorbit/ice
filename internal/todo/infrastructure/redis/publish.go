package redis

import (
	"context"
	"encoding/json"
	"github.com/gictorbit/ice/internal/todo/domain"
	"github.com/redis/go-redis/v9"
)

// Publish publishes a TodoItem to the Redis stream
func (s *StreamPublisher) Publish(ctx context.Context, todo *domain.TodoItem) error {
	data, _ := json.Marshal(todo)
	return s.rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: s.streamName,
		Values: map[string]any{"data": string(data)},
	}).Err()
}

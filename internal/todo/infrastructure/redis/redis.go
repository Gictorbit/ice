package redis

import (
	"github.com/redis/go-redis/v9"
)

type StreamPublisher struct {
	rdb        *redis.Client
	streamName string
}

func NewStreamPublisher(rdb *redis.Client, streamName string) *StreamPublisher {
	return &StreamPublisher{rdb: rdb, streamName: streamName}
}

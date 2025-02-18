package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

// CreateStream initializes a Redis stream if it doesn't exist
func CreateStream(streamName string) error {
	client := GetRedisClient()
	if client == nil {
		return fmt.Errorf("Redis client is not initialized")
	}

	// Add an empty message to create the stream if it doesn't exist
	_, err := client.XAdd(context.Background(), &redis.XAddArgs{
		Stream: streamName,
		Values: map[string]interface{}{"init": "stream_created"},
	}).Result()

	if err != nil {
		return fmt.Errorf("failed to create stream: %v", err)
	}
	return nil
}

// AddToStream adds a message to a Redis stream
func AddToStream(streamName string, message map[string]string) error {
	client := GetRedisClient()
	if client == nil {
		return fmt.Errorf("Redis client is not initialized")
	}

	_, err := client.XAdd(context.Background(), &redis.XAddArgs{
		Stream: streamName,
		Values: message,
	}).Result()

	if err != nil {
		return fmt.Errorf("failed to add message to stream: %v", err)
	}
	return nil
}

// ReadStream retrieves messages from a Redis stream
func ReadStream(streamName string) ([]map[string]interface{}, error) {
	client := GetRedisClient()
	if client == nil {
		return nil, fmt.Errorf("Redis client is not initialized")
	}

	// Read latest messages from stream
	streams, err := client.XRead(context.Background(), &redis.XReadArgs{
		Streams: []string{streamName, "0"},
		Count:   10,
		Block:   0,
	}).Result()

	if err != nil {
		return nil, fmt.Errorf("failed to read stream: %v", err)
	}

	var messages []map[string]interface{}
	for _, stream := range streams {
		for _, message := range stream.Messages {
			messages = append(messages, message.Values)
		}
	}

	return messages, nil
}

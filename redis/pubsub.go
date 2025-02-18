package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

// Publish sends a message to a Redis channel
func Publish(channel, message string) error {
	client := GetRedisClient()
	if client == nil {
		return fmt.Errorf("Redis client is not initialized")
	}

	err := client.Publish(context.Background(), channel, message).Err()
	if err != nil {
		return fmt.Errorf("❌ Failed to publish message: %v", err)
	}

	fmt.Printf("📡 Published to channel [%s]: %s\n", channel, message)
	return nil
}

// Subscribe listens for messages on a Redis channel
func Subscribe(channel string) (*redis.PubSub, error) {
	client := GetRedisClient()
	if client == nil {
		log.Println("❌ Redis client is not initialized")
		return nil, fmt.Errorf("redis client is not initialized")
	}

	subscription := client.Subscribe(context.Background(), channel)
	fmt.Printf("🔔 Subscribed to channel: %s\n", channel)

	return subscription, nil // ✅ Correctly return *redis.PubSub
}

// ListenMessages listens for messages and processes them
func ListenMessages(pubsub *redis.PubSub, channel string) {
	ch := pubsub.Channel() // ✅ This listens for messages

	go func() {
		for msg := range ch {
			fmt.Printf("📩 Received message on %s: %s\n", channel, msg.Payload)
		}
	}()
}

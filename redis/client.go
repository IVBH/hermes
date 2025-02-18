package redis

// redis/client.go - Redis connection setup

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

var rdb *redis.Client

// InitRedis initializes the Redis client
func InitRedis() {
	redisHost := "redis:6379" // Default for Docker
	if os.Getenv("TEST_ENV") == "true" || os.Getenv("LOCAL_ENV") == "true" {
		redisHost = "localhost:6379" // ✅ Use localhost when running locally
	}

	rdb = redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: "",
		DB:       0,
	})

	if err := Ping(); err != nil {
		log.Fatalf("❌ Failed to connect to Redis: %v", err)
	}
	log.Println("✅ Connected to Redis!")
}

// Ping checks Redis connectivity
func Ping() error {
	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	return err
}

// GetRedisClient returns the Redis client instance
func GetRedisClient() *redis.Client {
	if rdb == nil {
		log.Fatal("❌ Redis client is not initialized")
	}
	return rdb
}

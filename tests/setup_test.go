package tests

import (
	"context"
	"hermes/redis"
	"testing"
)

func TestMain(m *testing.M) {
	// Initialize Redis before running tests
	redis.InitRedis()
	ctx := context.Background()

	// Preload test admin key
	redis.GetRedisClient().Set(ctx, "hermes_admin_key", "test-admin-key", 0)

	// Preload test API key
	redis.GetRedisClient().HSet(ctx, "hermes_api_keys:test-app", "key", "test-api-key")

	// Run tests
	m.Run()
}

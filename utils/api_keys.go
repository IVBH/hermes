package utils

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"hermes/redis"
	"log"
)

// ValidateAPIKey checks if an API key is valid
func ValidateAPIKey(apiKey string) bool {
	ctx := context.Background()
	keys, err := redis.GetRedisClient().Keys(ctx, "hermes_api_keys:*").Result()
	if err != nil {
		return false
	}

	for _, key := range keys {
		storedHash, err := redis.GetRedisClient().HGet(ctx, key, "key").Result()

		// 🔍 Debugging logs
		log.Printf("🔍 Checking API Key - Given: %s | Stored: %s\n", apiKey, storedHash)

		if err == nil && storedHash == apiKey {
			return true
		}
	}
	return false
}

// ValidateAdminKey checks if the given admin key is valid
func ValidateAdminKey(key string) bool {
	ctx := context.Background()
	storedKey, err := redis.GetRedisClient().Get(ctx, "hermes_admin_key").Result()

	// 🔍 Debugging logs
	log.Printf("🔍 Checking Admin Key - Given: %s | Stored: %s\n", key, storedKey)

	return err == nil && storedKey == key
}

// hashAPIKey hashes an API key using SHA-256
func hashAPIKey(apiKey string) string {
	hash := sha256.Sum256([]byte(apiKey))
	return hex.EncodeToString(hash[:])
}

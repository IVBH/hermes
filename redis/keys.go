package redis

// redis/keys.go - Handles API key storage and validation in Redis

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"time"
)

const apiKeyPrefix = "hermes_api_keys:"
const adminKey = "hermes_admin_key"
const whitelistKey = "hermes_app_whitelist"

// StoreAPIKey saves an API key for an app in Redis with a TTL
func StoreAPIKey(appName, apiKey string, ttl int) error {
	hashedKey := hashAPIKey(apiKey)
	ctx := context.Background()
	if err := rdb.HSet(ctx, apiKeyPrefix+appName, "key", hashedKey).Err(); err != nil {
		return err
	}
	if err := rdb.Expire(ctx, apiKeyPrefix+appName, time.Duration(ttl)*time.Second).Err(); err != nil {
		return err
	}
	return nil
}

// ValidateAPIKey checks if an API key is valid
func ValidateAPIKey(apiKey string) bool {
	ctx := context.Background()
	keys, err := rdb.Keys(ctx, apiKeyPrefix+"*").Result()
	if err != nil {
		return false
	}
	for _, key := range keys {
		storedHash, err := rdb.HGet(ctx, key, "key").Result()
		if err == nil && storedHash == hashAPIKey(apiKey) {
			return true
		}
	}
	return false
}

// ValidateAdminKey checks if the given admin key is valid
func ValidateAdminKey(key string) bool {
	ctx := context.Background()
	storedKey, err := rdb.Get(ctx, adminKey).Result()
	return err == nil && storedKey == key
}

// AddToWhitelist adds an app to the whitelist
func AddToWhitelist(appName string) error {
	ctx := context.Background()
	return rdb.HSet(ctx, whitelistKey, appName, "true").Err()
}

// RemoveFromWhitelist removes an app from the whitelist
func RemoveFromWhitelist(appName string) error {
	ctx := context.Background()
	return rdb.HDel(ctx, whitelistKey, appName).Err()
}

// IsWhitelisted checks if an app is in the whitelist
func IsWhitelisted(appName string) bool {
	ctx := context.Background()
	exists, err := rdb.HExists(ctx, whitelistKey, appName).Result()
	return err == nil && exists
}

// hashAPIKey hashes an API key using SHA-256
func hashAPIKey(apiKey string) string {
	hash := sha256.Sum256([]byte(apiKey))
	return hex.EncodeToString(hash[:])
}

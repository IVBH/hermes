package tests

// tests/publish_test.go - Unit test for publishing messages

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"hermes/handlers"
)

func TestPublish(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/publish", handlers.Publish)

	// Mock request payload
	payload := map[string]string{
		"channel": "hermes_channel",
		"message": "Test message",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/publish", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "test-api-key")
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "message published")
}

package tests

// tests/subscribe_test.go - Unit test for subscribing to messages

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

func TestSubscribe(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/subscribe", handlers.Subscribe)

	// Mock request payload
	payload := map[string]string{
		"channel": "hermes_channel",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/subscribe", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "test-api-key")
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "messages")
}

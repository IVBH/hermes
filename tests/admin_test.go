package tests

// tests/admin_test.go - Unit test for admin whitelist management

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

func TestManageWhitelist(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/admin/whitelist", handlers.ManageWhitelist)

	// Mock request payload for adding to whitelist
	payload := map[string]string{
		"app_name": "test-app",
		"action":   "add",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/admin/whitelist", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "test-admin-key")
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "whitelist updated")
}

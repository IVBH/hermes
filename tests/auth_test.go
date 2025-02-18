package tests

// tests/auth_test.go - Unit test for API key validation

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"hermes/handlers"
)

func TestValidateAPIKeyMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.Use(handlers.ValidateAPIKeyMiddleware())
	r.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	// Test without API key
	req, _ := http.NewRequest("GET", "/protected", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// Test with valid API key
	req, _ = http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "test-api-key")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

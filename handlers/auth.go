package handlers

import (
	"github.com/gin-gonic/gin"
	"hermes/utils"
	"net/http"
)

// ValidateAPIKeyMiddleware ensures API key is valid before processing requests

// ValidateAPIKeyMiddleware godoc
// @Summary Middleware for API Key Validation
// @Description This middleware validates the API key for incoming requests
// @Tags Authentication
// @Accept  json
// @Produce  json
// @Param api_key header string true "API Key for authentication"
// @Success 200 {object} map[string]string "Valid API Key"
// @Failure 403 {object} map[string]string "Unauthorized request"
// @Router /protected [get]
func ValidateAPIKeyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("Authorization")
		if !utils.ValidateAPIKey(apiKey) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
			c.Abort()
			return
		}
		c.Next()
	}
}

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"hermes/redis"
	"hermes/utils"
)

// PublishRequest represents the request structure for publishing a message
type PublishRequest struct {
	Channel string `json:"channel" binding:"required"` // Channel to publish the message
	Message string `json:"message" binding:"required"` // The message content
}

// PublishResponse represents the response after publishing a message
type PublishResponse struct {
	Message string `json:"message"` // Confirmation message
}

// ErrorResponse represents a standard error response
type ErrorResponse struct {
	Error string `json:"error"` // Error description
}

// Publish handles publishing messages to a Redis channel
// @Summary Publish a message
// @Description Publish a message to a specific Redis channel
// @Tags Messaging
// @Accept json
// @Produce json
// @Param request body PublishRequest true "Publish request payload"
// @Success 200 {object} PublishResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /publish [post]
func Publish(c *gin.Context) {
	var req PublishRequest

	// Validate request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid request"})
		return
	}

	// Publish message to Redis
	err := redis.Publish(req.Channel, req.Message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to publish message"})
		return
	}

	// Increment Prometheus metric
	utils.MessagesPublished.Inc()

	// Success response
	c.JSON(http.StatusOK, PublishResponse{Message: "Message published"})
}

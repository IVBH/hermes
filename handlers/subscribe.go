package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"hermes/redis"
	"hermes/utils"
)

// SubscribeRequest represents the request structure for subscribing to a channel
type SubscribeRequest struct {
	Channel string `json:"channel" binding:"required"` // Channel to subscribe to
}

// SubscribeResponse represents the response when messages are received
type SubscribeResponse struct {
	Messages []string `json:"messages"` // List of received messages
}

// Subscribe handles subscription to a Redis channel
// @Summary Subscribe to a channel
// @Description Subscribe to a Redis channel and receive messages
// @Tags Messaging
// @Accept json
// @Produce json
// @Param request body SubscribeRequest true "Subscribe request payload"
// @Success 200 {object} SubscribeResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /subscribe [post]
func Subscribe(c *gin.Context) {
	var req SubscribeRequest

	// Validate request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid request"})
		return
	}

	// Subscribe to Redis channel
	subscription, err := redis.Subscribe(req.Channel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to subscribe"})
		return
	}

	// Increment Prometheus metric
	utils.ActiveSubscribers.Inc()

	// Start listening for messages
	go redis.ListenMessages(subscription, req.Channel)

	c.JSON(http.StatusOK, gin.H{"message": "Subscribed to " + req.Channel})
}

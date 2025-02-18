package handlers

import (
	"github.com/gin-gonic/gin"
	"hermes/redis"
	"net/http"
)

// CreateStreamRequest represents the JSON request structure for creating a stream
type CreateStreamRequest struct {
	StreamName string `json:"stream_name" binding:"required"`
}

// AddToStreamRequest represents the JSON request structure for adding a message to a stream
type AddToStreamRequest struct {
	StreamName string            `json:"stream_name" binding:"required"`
	Message    map[string]string `json:"message" binding:"required"`
}

// CreateStream handles the creation of a Redis Stream
// @Summary Create a Redis Stream
// @Description Creates a new stream in Redis
// @Tags Streams
// @Accept json
// @Produce json
// @Param request body CreateStreamRequest true "Stream Name"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /stream/create [post]
func CreateStream(c *gin.Context) {
	var req CreateStreamRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := redis.CreateStream(req.StreamName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create stream"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Stream created", "stream_name": req.StreamName})
}

// AddToStream handles adding a message to a Redis Stream
// @Summary Add message to Redis Stream
// @Description Adds a message to an existing Redis Stream
// @Tags Streams
// @Accept json
// @Produce json
// @Param request body AddToStreamRequest true "Stream Data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /stream/add [post]
func AddToStream(c *gin.Context) {
	var req AddToStreamRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := redis.AddToStream(req.StreamName, req.Message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add message to stream"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message added to stream"})
}

// ReadStream handles reading messages from a Redis Stream
// @Summary Read messages from Redis Stream
// @Description Reads messages from a Redis Stream
// @Tags Streams
// @Accept json
// @Produce json
// @Param stream_name query string true "Stream Name"
// @Success 200 {array} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /stream/read [get]
func ReadStream(c *gin.Context) {
	streamName := c.Query("stream_name")
	if streamName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Stream name is required"})
		return
	}

	messages, err := redis.ReadStream(streamName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read stream"})
		return
	}

	c.JSON(http.StatusOK, messages)
}

package handlers

import (
	"github.com/gin-gonic/gin"
	"hermes/redis"
	"net/http"
)

// HealthCheck godoc
// @Summary Checks the health status of Hermes API
// @Description This endpoint checks if Hermes and Redis are running properly
// @Tags System
// @Produce  json
// @Success 200 {object} map[string]string "System is healthy"
// @Failure 500 {object} map[string]string "Redis is unreachable"
// @Router /health [get]
func HealthCheck(c *gin.Context) {
	status := "ok"
	redisStatus := "connected"

	if err := redis.Ping(); err != nil {
		redisStatus = "disconnected"
		status = "error"
	}

	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"redis":  redisStatus,
	})
}

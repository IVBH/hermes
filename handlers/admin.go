package handlers

// handlers/admin.go - Handles whitelisted app management

import (
	"github.com/gin-gonic/gin"
	"hermes/redis"
	"hermes/utils"
	"net/http"
)

// ManageWhitelistRequest represents the JSON request body for whitelist management
type ManageWhitelistRequest struct {
	AppName string `json:"app_name" binding:"required"`
	Action  string `json:"action" binding:"required"` // "add" or "remove"
}

// ManageWhitelist allows admins to add or remove apps from the whitelist

// ManageWhitelist godoc
// @Summary Manages the whitelist of approved applications
// @Description This endpoint allows an admin to add or remove applications from the whitelist
// @Tags Admin
// @Accept  json
// @Produce  json
// @Param admin_key header string true "Admin Key for authentication"
// @Param request body object true "Whitelist Update Request"
// @Success 200 {object} map[string]string "Whitelist updated successfully"
// @Failure 400 {object} map[string]string "Invalid request format"
// @Failure 403 {object} map[string]string "Unauthorized request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /admin/whitelist [post]
func ManageWhitelist(c *gin.Context) {
	var req ManageWhitelistRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Validate Admin API Key
	adminKey := c.GetHeader("Authorization")
	if !utils.ValidateAdminKey(adminKey) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid admin key"})
		return
	}

	if req.Action == "add" {
		if err := redis.AddToWhitelist(req.AppName); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add app to whitelist"})
			return
		}
	} else if req.Action == "remove" {
		if err := redis.RemoveFromWhitelist(req.AppName); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove app from whitelist"})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid action"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "whitelist updated"})
}

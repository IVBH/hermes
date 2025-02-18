package handlers

import "github.com/gin-gonic/gin"

// Register handles app registration

// Register godoc
// @Summary Registers a new app
// @Description This endpoint registers an application for pub/sub
// @Tags auth
// @Accept  json
// @Produce  json
// @Param app body object true "App Registration Data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /register [post]
func Register(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Register endpoint working!"})
}

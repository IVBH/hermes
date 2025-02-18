// main.go - Entry point for Hermes API
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"       // Swagger files
	"github.com/swaggo/gin-swagger" // Gin Swagger middleware
	_ "hermes/docs"
	"hermes/handlers"
	"hermes/redis"
	"hermes/utils"
	"log"
	"os"
)

func main() {
	// Initialize Redis connection
	redis.InitRedis()

	// Initialize Prometheus metrics
	utils.InitMetrics()

	// ✅ Use GIN_MODE environment variable (default: release)
	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		mode = gin.ReleaseMode // Default to release mode if not set
	}
	gin.SetMode(mode)

	// Set up Gin router
	r := gin.Default()
	err := r.SetTrustedProxies([]string{"192.168.0.0/16", "10.0.0.0/8"})
	if err != nil {
		return
	}

	// API Endpoints
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/register", handlers.Register)
	r.POST("/renew", handlers.Renew)
	r.POST("/publish", handlers.Publish)
	r.POST("/subscribe", handlers.Subscribe)
	r.POST("/admin/whitelist", handlers.ManageWhitelist)
	r.GET("/health", handlers.HealthCheck)
	r.GET("/metrics", handlers.MetricsHandler)
	r.POST("/stream/create", handlers.CreateStream)
	r.POST("/stream/add", handlers.AddToStream)
	r.GET("/stream/read", handlers.ReadStream)

	// Start Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("✅ Hermes API Server Running on :%s", port)
	r.Run(":8080") // Start server
}

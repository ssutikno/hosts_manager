package main

import (
	"hosts-manager/config"
	"hosts_manager/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	config.LoadConfig("config.json")

	// Initialize Gin router
	router := gin.Default()

	// Serve static files
	router.Static("/static", "./static")

	// Load HTML templates
	router.LoadHTMLGlob("templates/**/*.html")

	// Initialize routes
	routes.InitializeRoutes(router)

	// Start the server
	router.Run(":" + config.GetConfig().Server.Port)
}

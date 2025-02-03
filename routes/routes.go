package routes

import (
	"hosts_manager/controllers"

	"github.com/gin-gonic/gin"
)

// InitializeRoutes initializes the routes for the application
func InitializeRoutes(router *gin.Engine) {
	// Home route
	router.GET("/", controllers.Home)

	// Host management routes
	router.GET("/hosts", controllers.ListHosts)
	router.GET("/hosts/:id", controllers.GetHost)
	router.POST("/hosts", controllers.CreateHost)
	router.PUT("/hosts/:id", controllers.UpdateHost)
	router.DELETE("/hosts/:id", controllers.DeleteHost)
}

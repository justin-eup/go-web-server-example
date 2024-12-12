package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/justin-eup/web-server-example/internal/handlers"
)

// SetupRouter initializes the application routes
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Route groups
	api := router.Group("/api")
	{
		api.GET("/hello/:name", handlers.HelloHandler)
		api.GET("/health", handlers.HealthCheckHandler)
	}

	return router
}

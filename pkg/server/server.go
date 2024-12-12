package server

import (
	"fmt"

	"github.com/justin-eup/web-server-example/config"
	"github.com/justin-eup/web-server-example/internal/routes"
)

// Start initializes and starts the server
func Start() error {
	// Load configuration
	cfg := config.Load()

	// Setup routes
	router := routes.SetupRouter()

	// Start the server
	address := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	return router.Run(address)
}

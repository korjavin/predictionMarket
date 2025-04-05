package main

import (
	"fmt"
	"log"

	"github.com/korjavin/predictionMarket/backend/pkg/api"
	"github.com/korjavin/predictionMarket/backend/pkg/config"
	"github.com/korjavin/predictionMarket/backend/pkg/ethereum"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize Ethereum service
	ethService, err := ethereum.NewService(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize Ethereum service: %v", err)
	}
	defer ethService.Close()

	// Create API handler
	handler := api.NewHandler(ethService)

	// Setup router
	router := api.SetupRouter(handler)

	// Start server
	serverAddr := fmt.Sprintf(":%s", cfg.Port)
	fmt.Printf("Server starting on %s\n", serverAddr)
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

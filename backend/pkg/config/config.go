package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the application
type Config struct {
	// Ethereum RPC URL (e.g. Infura, Alchemy, local node)
	EthereumRPC string
	// Port the API server will listen on
	Port string
	// Network ID (e.g. 1 for mainnet, 11155111 for Sepolia, etc.)
	NetworkID string
	// Private key for signing transactions (without 0x prefix)
	PrivateKey string
}

// Load reads configuration from environment variables
func Load() (*Config, error) {
	// Load .env file if it exists
	_ = godotenv.Load()

	config := &Config{
		EthereumRPC: getEnv("ETHEREUM_RPC", "http://localhost:8545"),
		Port:        getEnv("PORT", "8080"),
		NetworkID:   getEnv("NETWORK_ID", "11155111"), // Default to Sepolia testnet
		PrivateKey:  getEnv("PRIVATE_KEY", ""),
	}

	return config, nil
}

// Helper function to read environment variables with fallback
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return strings.TrimSpace(value)
	}
	return fallback
}

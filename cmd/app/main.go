package main

import (
	"api/config"
	"api/internal/app"
	"api/pkg/logging"
)

func main() {
	// Configuration
	logger := logging.NewLogger()
	cfg, err := config.NewConfig()

	if err != nil {
		logger.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg, &logger)
}

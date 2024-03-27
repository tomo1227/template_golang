package main

import (
	"log/slog"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/tomo1227/template_golang/internal/adapter/router"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("Start up the server...")

	router, err := router.NewRouter()
	startFiberServer(":8080", router.Fiber, logger)
	if err != nil {
		logger.Error("Failed to initialize router", "error: ", err)
		os.Exit(1)
	}
}

// startFiberServer starts a REST server.
func startFiberServer(port string, server *fiber.App, logger *slog.Logger) {
	if err := server.Listen(port); err != nil {
		logger.Error("Failed to start server: ", err)
	}
}

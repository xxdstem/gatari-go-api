package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xxdstem/gatari-go-api/handlers"
	"github.com/xxdstem/gatari-go-api/logger"
)

func init() {
	logger.Init()
	L := logger.New()
	L.Info("Logger loaded!")
	L.Error("This is error")
}

func main() {
	f := fiber.New()

	api := f.Group("/api/v1")

	api.Get("/hello", handlers.Hello)

	f.Listen(":8080")
}

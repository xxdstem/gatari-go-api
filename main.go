package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xxdstem/gatari-go-api/handlers"
)

func main() {
	f := fiber.New()

	api := f.Group("/api/v1")

	api.Get("/hello", handlers.Hello)

	f.Listen(":8080")
}

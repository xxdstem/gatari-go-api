package main

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/xxdstem/gatari-go-api/handlers"
	"github.com/xxdstem/gatari-go-api/logger"
	"github.com/xxdstem/gatari-go-api/utils"
)

var Config *utils.Config

var L *logger.Logger

func init() {
	utils.CursorHide()
	logger.Init()

	L = logger.New()

	L.Info("Logger loaded!")

	var err error
	Config, err = utils.NewConfig()
	if err != nil {
		time.Sleep(3 * time.Second)
		os.Exit(1)
	}
}

func main() {
	f := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	utils.SigHandler(f) //handling CTRL+C signals to shutdown gracefully

	api := f.Group("/api/v1")

	api.Get("/hello", handlers.Hello)

	L.Info("Starting web server...")
	f.Listen(":8008")
}

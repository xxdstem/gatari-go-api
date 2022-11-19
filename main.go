package main

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/xxdstem/gatari-go-api/internal/handlers"
	"github.com/xxdstem/gatari-go-api/internal/logger"
	"github.com/xxdstem/gatari-go-api/utils"
	cfg "github.com/xxdstem/gatari-go-api/utils/config"
)

var Config *cfg.Config

var L *logger.Logger

func init() {
	utils.CursorHide()
	logger.Init()

	L = logger.New()

	L.Info("Logger loaded!")

	var err error
	Config, err = cfg.NewConfig()
	if err != nil {
		time.Sleep(3 * time.Second)
		os.Exit(1)
	}
}

func main() {
	f := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	f.Use(cors.New()) //pls try default values, if it doesn't work, set config manually

	// HANDLERS
	//   api/v1

	api := f.Group("/api/v1")

	api.Get("/hello", handlers.Hello)

	// starting web

	L.Info("Starting web server...")
	L.Info("Press ctrl+c to exit")

	utils.SigHandler(f) //handling CTRL+C signals to shutdown gracefully

	if err := f.Listen(":8008"); err != nil {
		L.Error(err.Error())
		time.Sleep(time.Millisecond * 400)
	}
}

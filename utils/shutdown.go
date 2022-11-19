package utils

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/xxdstem/gatari-go-api/internal/logger"
)

var Run = make(chan bool)

func SigHandler(f *fiber.App) {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-exit
		L := logger.New()

		fmt.Print("\r")
		L.Warn("Shutting down...")
		time.Sleep(time.Millisecond * 400)

		CursorShow()
		f.Shutdown()
		Run <- false
	}()
}

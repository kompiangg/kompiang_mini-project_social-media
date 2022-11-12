package config

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
)

type Server struct {
	E    *echo.Echo
	Port string
}

func StartServer(param Server) error {
	errChan := make(chan error, 1)

	go func() {
		if err := param.E.Start(param.Port); err != nil {
			errChan <- err
		}
	}()
	defer func() {
		param.E.Shutdown(context.Background())
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errChan:
		log.Println("[HANDLER ERROR] Error while starting server: ", err.Error())
		return err
	case <-signalChan:
		log.Println("[INFO] Server closed gracefully")
		return nil
	}
}

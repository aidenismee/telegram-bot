package server

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/nekizz/telegram-bot/configs"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type server struct {
	Engine *echo.Echo
}

func NewServer(configuration *configs.Configuration) *server {
	cfg := loadConfig(withPort(configuration.Port),
		withReadTimeout(configuration.ReadTimeout),
		withWriteTimeout(configuration.WriteTimeout))

	engine := echo.New()

	for _, middleware := range cfg.middlewares {
		engine.Use(middleware)
	}

	engine.Debug = cfg.Debug
	engine.Binder = NewBinder()
	engine.Logger.SetLevel(log.ERROR)
	engine.HTTPErrorHandler = NewErrorHandler(engine).Handle
	if engine.Debug {
		engine.Logger.SetLevel(log.DEBUG)
	}

	engine.Server.Addr = fmt.Sprintf(":%d", cfg.Port)
	engine.Server.ReadTimeout = time.Duration(cfg.ReadTimeout) * time.Minute
	engine.Server.WriteTimeout = time.Duration(cfg.WriteTimeout) * time.Minute

	return &server{
		Engine: engine,
	}
}

func Start(e *echo.Echo) {
	idleConnectionClosed := make(chan struct{})
	go Shutdown(e, idleConnectionClosed)

	e.HideBanner = false
	if err := e.StartServer(e.Server); err != nil {
		if err == http.ErrServerClosed {
			e.Logger.Info("http server stopped")
		} else {
			e.Logger.Errorf("http server StartServer: %v", err)
		}
	}
	<-idleConnectionClosed
}

func Shutdown(e *echo.Echo, idleConnectionClosed chan struct{}) {
	sigint := make(chan os.Signal, 1)

	// interrupt signal sent from terminal
	signal.Notify(sigint, os.Interrupt)

	// sigterm signal sent from kubernetes
	signal.Notify(sigint, syscall.SIGTERM)

	sigrev := <-sigint
	e.Logger.Infof("signal received: %s", sigrev.String())

	// We received an interrupt signal, shut down.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		// Error from closing listeners, or context timeout:
		e.Logger.Errorf("http server Shutdown: %v", err)
	}

	close(idleConnectionClosed)
}

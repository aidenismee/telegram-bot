package server

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/nekizz/telegram-bot/configs"
	"github.com/nekizz/telegram-bot/pkg/server/middleware"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	defaultStage        = "local"
	defaultPort         = 8080
	defaultReadTimeout  = 10
	defaultWriteTimeout = 5
)

type config struct {
	Stage        string
	Port         int
	ReadTimeout  int
	WriteTimeout int
	Debug        bool
	middlewares  []echo.MiddlewareFunc
}

type Config func(*config)

type Server struct {
	Engine *echo.Echo
}

func fillDefaults() *config {
	return &config{
		Stage:        defaultStage,
		Port:         defaultPort,
		ReadTimeout:  defaultReadTimeout,
		WriteTimeout: defaultWriteTimeout,
		middlewares: []echo.MiddlewareFunc{
			middleware.WithCORS(),
			middleware.WithTimeout(),
			middleware.WithCorrelationID(),
			middleware.WithRecover(),
			middleware.WithRateLimiter(),
		},
	}
}

func loadConfig(cfg ...Config) *config {
	cf := fillDefaults()

	for _, fn := range cfg {
		fn(cf)
	}

	return cf
}

func NewServer(configuration *configs.Configuration) *Server {
	cfg := loadConfig()

	engine := echo.New()

	for _, middleware := range cfg.middlewares {
		engine.Use(middleware)
	}

	engine.HTTPErrorHandler = NewErrorHandler(engine).Handle
	engine.Binder = NewBinder()
	engine.Debug = cfg.Debug
	if engine.Debug {
		engine.Logger.SetLevel(log.DEBUG)
	} else {
		engine.Logger.SetLevel(log.ERROR)
	}

	engine.Server.Addr = fmt.Sprintf(":%d", cfg.Port)
	engine.Server.ReadTimeout = time.Duration(cfg.ReadTimeout) * time.Minute
	engine.Server.WriteTimeout = time.Duration(cfg.WriteTimeout) * time.Minute

	return &Server{
		Engine: engine,
	}
}

func Start(e *echo.Echo) {
	idleConnectionClosed := make(chan struct{})
	go func() {
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
	}()

	e.HideBanner = true
	if err := e.StartServer(e.Server); err != nil {
		if err == http.ErrServerClosed {
			e.Logger.Info("http server stopped")
		} else {
			e.Logger.Errorf("http server StartServer: %v", err)
		}
	}
	<-idleConnectionClosed
}

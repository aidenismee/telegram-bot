package server

import (
	"github.com/labstack/echo/v4"
	"github.com/nekizz/telegram-bot/pkg/server/middleware"
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

func withPort(port int) Config {
	return func(c *config) {
		if port != 0 {
			c.Port = port
		}
	}
}

func withReadTimeout(readTimeout int) Config {
	return func(c *config) {
		if readTimeout != 0 {
			c.ReadTimeout = readTimeout
		}
	}
}

func withWriteTimeout(writeTimeout int) Config {
	return func(c *config) {
		if writeTimeout != 0 {
			c.WriteTimeout = writeTimeout

		}
	}
}

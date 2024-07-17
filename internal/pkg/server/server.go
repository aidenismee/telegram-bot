package server

import (
	"github.com/labstack/echo/v4"
	"github.com/nekizz/telegram-bot/configs"
	"github.com/nekizz/telegram-bot/internal/pkg/router"
	"github.com/nekizz/telegram-bot/pkg/server"
)

func initServerConfig() *echo.Echo {
	cfg, err := configs.Load()
	if err != nil {
		panic(err)
	}

	server := server.NewServer(cfg)
	router.NewRouter(cfg).RegisterHandler(server.Engine)

	return server.Engine
}

func Start() {
	server.Start(initServerConfig())
}

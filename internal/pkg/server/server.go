package server

import (
	"github.com/labstack/echo/v4"
	"github.com/nekizz/telegram-bot/configs"
	"github.com/nekizz/telegram-bot/internal/migration"
	"github.com/nekizz/telegram-bot/internal/pkg/router"
	"github.com/nekizz/telegram-bot/pkg/server"
	"log"
)

func initServerConfig() *echo.Echo {
	cfg, err := configs.Load()
	if err != nil {
		log.Println(err)
	}
	cfg.DbPsn = "dbhuan:0866444202@tcp(103.173.254.82:3306)/telegram_bot?charset=utf8mb4&parseTime=True&loc=Local"

	migration.Run(cfg)
	server := server.NewServer(cfg)
	router.NewRouter(cfg).RegisterHandler(server.Engine)

	return server.Engine
}

func Start() {
	server.Start(initServerConfig())
}

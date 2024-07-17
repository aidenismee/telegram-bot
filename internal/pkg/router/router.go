package router

import (
	"github.com/labstack/echo/v4"
	"github.com/nekizz/telegram-bot/configs"
	"github.com/nekizz/telegram-bot/internal/api/health"
)

type Router interface {
	RegisterHandler(*echo.Echo) Router
}

type router struct {
	service *service
}

func NewRouter(cfg *configs.Configuration) Router {
	return &router{service: NewService(cfg)}
}

func (r *router) RegisterHandler(e *echo.Echo) Router {
	apiGroup := e.Group("/api")
	v1Group := apiGroup.Group("/v1")

	health.NewHandler(v1Group.Group("/health"))

	go r.service.telegramSvc.CommandHandler()

	return r
}

package router

import (
	"github.com/labstack/echo/v4"
	"github.com/nekizz/telegram-bot/configs"
	"github.com/nekizz/telegram-bot/internal/api/health"
	"github.com/nekizz/telegram-bot/internal/api/user"
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
	apiGroup := e.Group("/apis")
	{
		v1Group := apiGroup.Group("/v1")
		{
			health.NewHandler(v1Group.Group("/healths"))
			user.NewHandler(v1Group.Group("/users"))
		}
	}

	//go r.service.telegramSvc.CommandHandler()

	return r
}

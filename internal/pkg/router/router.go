package router

import (
	"github.com/labstack/echo/v4"
	"github.com/nekizz/telegram-bot/configs"
	"github.com/nekizz/telegram-bot/internal/api/health"
	"github.com/nekizz/telegram-bot/internal/pkg/manager"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Router interface {
	RegisterHandler(*echo.Echo) Router
}

type router struct {
	manager manager.Manager
}

func NewRouter(cfg *configs.Configuration) Router {
	return &router{manager: manager.NewManager(cfg)}
}

func (r *router) RegisterHandler(e *echo.Echo) Router {
	apiGroup := e.Group("/apis")
	{
		v1Group := apiGroup.Group("/v1")
		{
			userGroup := v1Group.Group("/users")
			{
				userGroup.GET("/hellos", r.manager.UserHandler().Hello)
			}
			telegramGroup := v1Group.Group("/telegrams")
			{
				telegramGroup.POST("/alerts", r.manager.TelegramHandler().AlertJob)
				telegramGroup.POST("/birthdays", r.manager.TelegramHandler().CheckBirthdays)
			}
			health.NewHandler(v1Group.Group("/healths"))
		}

	}

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	r.manager.TelegramHandler().HandleCommand()

	return r
}

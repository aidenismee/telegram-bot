package router

import (
	"github.com/labstack/echo/v4"
	"github.com/nekizz/telegram-bot/configs"
	"github.com/nekizz/telegram-bot/internal/api/health"
	"github.com/nekizz/telegram-bot/internal/api/telegram"
	"github.com/nekizz/telegram-bot/internal/api/user"
	"github.com/nekizz/telegram-bot/internal/pkg/manager"
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
			health.NewHandler(v1Group.Group("/healths"))
			user.NewHandler(r.manager.UserService(), v1Group.Group("/users"))
			telegram.NewHandler(r.manager.TeleService(), v1Group.Group("/telegrams"))
		}
	}

	return r
}

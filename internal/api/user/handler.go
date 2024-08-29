package user

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Service interface {
	Hello() error
}

type Handler struct {
	userService Service
}

func NewHandler(userService Service) *Handler {
	return &Handler{
		userService: userService,
	}
}

func (h *Handler) Hello(c echo.Context) error {
	if err := h.userService.Hello(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

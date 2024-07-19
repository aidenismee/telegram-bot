package user

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	s *Service
}

func NewHandler(eg *echo.Group) {
	handler := Handler{}

	eg.POST("/alerts", handler.alertJob)
	eg.POST("/birthdays", handler.checkBirthdays)
}

func (h *Handler) checkBirthdays(c echo.Context) error {
	if err := h.s.checkBirthdays(c); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

func (h *Handler) alertJob(c echo.Context) error {
	if err := h.s.alertJob(c); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

package telegram

import (
	"net/http"

	logger "github.com/nekizz/telegram-bot/pkg/log"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	s *Service
}

func NewHandler(s *Service, eg *echo.Group) {
	handler := Handler{s: s}

	eg.POST("/alerts", handler.alertJob)
	eg.POST("/birthdays", handler.checkBirthdays)

	go handler.handleCommand(nil)
}

func (h *Handler) alertJob(c echo.Context) error {
	logFields := map[string]any{"handler": "alertJob"}
	log := h.newLogger(c)

	if err := h.s.alertJob(log.WithLogger(c)); err != nil {
		log.WithFields(logFields).WithErr(err).Error("bind request payload failed")
		return err
	}

	return c.NoContent(http.StatusOK)
}

func (h *Handler) checkBirthdays(c echo.Context) error {
	logFields := map[string]any{"handler": "checkBirthdays"}
	log := h.newLogger(c)

	if err := h.s.checkBirthdays(log.WithLogger(c)); err != nil {
		log.WithFields(logFields).WithErr(err).Error("bind request payload failed")
		return err
	}

	return c.NoContent(http.StatusOK)
}

func (h *Handler) handleCommand(c echo.Context) error {
	return h.s.handleCommand()
}

func (h *Handler) newLogger(c echo.Context) *logger.Logger {
	return logger.NewLog("telegram").WithCid(c.Request().Header.Get(echo.HeaderXRequestID))
}

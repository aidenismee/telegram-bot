package telegram

import (
	"net/http"

	logger "github.com/nekizz/telegram-bot/pkg/log"

	"github.com/labstack/echo/v4"
)

type Service interface {
	handleCommand() error
	alertJob(c echo.Context) error
	checkBirthdays(c echo.Context) error
}

type Handler struct {
	s Service
}

func NewHandler(s Service) *Handler {
	return &Handler{s: s}
}

func (h *Handler) AlertJob(c echo.Context) error {
	logFields := map[string]any{"handler": "alertJob"}
	log := h.newLogger(c)

	if err := h.s.alertJob(log.WithLogger(c)); err != nil {
		log.WithFields(logFields).WithErr(err).Error("bind request payload failed")
		return c.NoContent(http.StatusBadRequest)
	}

	return c.NoContent(http.StatusOK)
}

func (h *Handler) CheckBirthdays(c echo.Context) error {
	logFields := map[string]any{"handler": "checkBirthdays"}
	log := h.newLogger(c)

	if err := h.s.checkBirthdays(log.WithLogger(c)); err != nil {
		log.WithFields(logFields).WithErr(err).Error("bind request payload failed")
		return err
	}

	return c.NoContent(http.StatusOK)
}

func (h *Handler) HandleCommand() error {
	return h.s.handleCommand()
}

func (h *Handler) newLogger(c echo.Context) *logger.Logger {
	return logger.NewLog("telegram").WithCid(c.Request().Header.Get(echo.HeaderXRequestID))
}

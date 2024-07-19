package errors

import (
	"github.com/nekizz/telegram-bot/pkg/server"
	"net/http"
)

var (
	ErrInvalidRequest = server.NewHTTPError(http.StatusBadRequest, "INVALID_REQUEST", "Request is invalid or malformed")
)

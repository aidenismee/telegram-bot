package health

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

var (
	version   = "local"                         // sha1 revision used to build the server
	buildTime = time.Now().Format(time.RFC3339) // when the server was built
)

type Handler struct{}

func NewHandler(eg *echo.Group) {
	h := Handler{}

	eg.GET("", h.check)
}

func (h *Handler) check(c echo.Context) error {

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":     "OK!",
		"version":    version,
		"build_time": buildTime,
	})
}

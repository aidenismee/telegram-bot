package user

import (
	"github.com/labstack/echo/v4"
)

type Handler struct {
	s *Service
}

func NewHandler(svc *Service, eg *echo.Group) {
	_ = Handler{s: svc}
}

package middleware

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	nanoId "github.com/matoous/go-nanoid/v2"
	"net/http"
	"strings"
	"time"
)

func WithTimeout() echo.MiddlewareFunc {
	return middleware.TimeoutWithConfig(middleware.TimeoutConfig{Timeout: 60 * time.Second})
}

func WithRecover() echo.MiddlewareFunc {
	return middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 2 << 5,
		LogErrorFunc: func(e echo.Context, err error, stack []byte) error {
			fmt.Println("panic error:", err)
			return nil
		},
	})
}

func WithRateLimiter() echo.MiddlewareFunc {
	return middleware.RateLimiterWithConfig(middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: 10, Burst: 50, ExpiresIn: 2 * time.Minute},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			return context.JSON(http.StatusForbidden, nil)
		},
		DenyHandler: func(context echo.Context, identifier string, err error) error {
			return context.JSON(http.StatusTooManyRequests, nil)
		},
	})
}

func WithCORS() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowCredentials: true,
	})
}

func WithCorrelationID() echo.MiddlewareFunc {
	config := middleware.RequestIDConfig{
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, "health")
		},
		Generator: func() string {
			cid, err := nanoId.New()
			if cid != "" && err == nil {
				return cid
			}

			return uuid.New().String()
		},
		TargetHeader: echo.HeaderXCorrelationID,
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			req := c.Request()
			res := c.Response()

			cid := req.Header.Get(config.TargetHeader)
			if cid == "" {
				cid = config.Generator()
				req.Header.Set(config.TargetHeader, cid)
			}

			res.Header().Set(config.TargetHeader, cid)

			if config.RequestIDHandler != nil {
				config.RequestIDHandler(c, cid)
			}

			return next(c)
		}
	}
}

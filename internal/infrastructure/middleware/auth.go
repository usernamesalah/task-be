package middleware

import (
	"crypto/subtle"

	"task-be/internal/infrastructure/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func BasicAuth(cfg *config.Config) echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if subtle.ConstantTimeCompare([]byte(username), []byte(cfg.Auth.Username)) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte(cfg.Auth.Password)) == 1 {
			return true, nil
		}
		return false, nil
	})
}

package middleware

import (
	"github.com/labstack/echo"
)

func Track(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		println("incoming request to ", c.Request().RequestURI)
		return next(c)
	}
}

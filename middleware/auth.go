package middleware

import (
	"log"

	"github.com/labstack/echo"
)

// Track : Log path
func Track(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Printf("Task Sended" + c.Request().RequestURI)
		return next(c)
	}
}

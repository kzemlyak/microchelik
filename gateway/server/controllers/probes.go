package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitProbesController(e *echo.Echo) {
	// for startup probe
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	// for readiness probe
	e.GET("/ready", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	e.GET("/hello", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello, gateway service! <3")
	})
}

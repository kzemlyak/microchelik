package server

import (
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	controllers "github.com/kzemlyak/microchelik/gateway/server/controllers"
)

func InitHttpServer() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	controllers.InitProbesController(e)
	controllers.InitUserController(e)
	controllers.InitSocialController(e)
	controllers.InitChatController(e)

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	time.Sleep(5 * time.Second)

	e.Logger.Fatal(e.Start(":" + httpPort))
}

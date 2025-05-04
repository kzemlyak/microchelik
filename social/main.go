package main

import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// for startup probe
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	// for readiness probe
	e.GET("/ready", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	e.GET("/hello", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello, social service! <3")
	})

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	time.Sleep(5 * time.Second)

	e.Logger.Fatal(e.Start(":" + httpPort))
}

post /social/profiles
	body: {
		"user_id": "1",
		"nickname": "test",
	}

get /social/profiles/search
	query: {
		"nickname": "test",
	}
	response: {
		"profiles": [
			{
				"user_id": "1",
			},
		],
	}

post /social/friends/requests
	body: {
		"user_id": "1",
		"friend_id": "2",
	}

get /social/friends/requests
	body: {
		"user_id": "1",
		"friend_id": "2",
	}

post /social/friends/requests/accept
	body: {
		"user_id": "1",
		"friend_id": "2",
	}

post /social/friends/requests/decline
	body: {
		"user_id": "1",
		"friend_id": "2",
	}

get /social/friends
	body: {
		"user_id": "1",
		"friend_id": "2",
	}

delete /social/friends
	body: {
		"user_id": "1",
		"friend_id": "2",
	}


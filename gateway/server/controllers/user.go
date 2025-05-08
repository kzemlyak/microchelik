package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-openapi/strfmt"
	models "github.com/kzemlyak/microchelik/gateway/server/models"
	"github.com/labstack/echo/v4"
)

func registerUser(c echo.Context) error {
	var request models.RegisterRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := request.Validate(strfmt.Default); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "User created")
}

func loginUser(c echo.Context) error {
	var request models.LoginRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := request.Validate(strfmt.Default); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "User logged in")
}

func logoutUser(c echo.Context) error {
	var request models.LogoutRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := request.Validate(strfmt.Default); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "User logged out")
}

func InitUserController(e *echo.Echo) {
	e.POST("/api/v1/user/register", registerUser)
	e.POST("/api/v1/user/login", loginUser)
	e.POST("/api/v1/user/logout", logoutUser)
}

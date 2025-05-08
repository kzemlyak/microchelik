package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-openapi/strfmt"
	models "github.com/kzemlyak/microchelik/gateway/server/models"
	"github.com/labstack/echo/v4"
)

func createChat(c echo.Context) error {
	var request models.CreateChatRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := request.Validate(strfmt.Default); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "Chat created")
}

func deleteChat(c echo.Context) error {
	chatId := c.Param("chatId")
	if chatId == "" {
		return c.JSON(http.StatusBadRequest, "chatId parameter is required")
	}

	return c.JSON(http.StatusOK, "Chat deleted")
}

func getChatMessages(c echo.Context) error {
	chatId := c.Param("chatId")
	if chatId == "" {
		return c.JSON(http.StatusBadRequest, "chatId parameter is required")
	}

	return c.JSON(http.StatusOK, "Chat messages retrieved")
}

func sendMessage(c echo.Context) error {
	chatId := c.Param("chatId")
	if chatId == "" {
		return c.JSON(http.StatusBadRequest, "chatId parameter is required")
	}

	var request models.SendMessageRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := request.Validate(strfmt.Default); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "Message sent")
}

func InitChatController(e *echo.Echo) {
	e.POST("/api/v1/chat", createChat)
	e.DELETE("/api/v1/chat/:chatId", deleteChat)
	e.GET("/api/v1/chat/:chatId/messages", getChatMessages)
	e.POST("/api/v1/chat/:chatId/messages", sendMessage)
}

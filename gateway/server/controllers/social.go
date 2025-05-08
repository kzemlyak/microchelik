package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-openapi/strfmt"
	models "github.com/kzemlyak/microchelik/gateway/server/models"
	"github.com/labstack/echo/v4"
)

func createProfile(c echo.Context) error {
	var request models.CreateProfileRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := request.Validate(strfmt.Default); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "Profile created")
}

func searchProfiles(c echo.Context) error {
	query := c.QueryParam("query")
	if query == "" {
		return c.JSON(http.StatusBadRequest, "query parameter is required")
	}

	return c.JSON(http.StatusOK, "Profiles found")
}

func createFriendRequest(c echo.Context) error {
	var request models.CreateFriendRequestRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := request.Validate(strfmt.Default); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "Friend request created")
}

func getFriendRequests(c echo.Context) error {
	return c.JSON(http.StatusOK, "Friend requests retrieved")
}

func acceptFriendRequest(c echo.Context) error {
	var request models.AcceptFriendRequestRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := request.Validate(strfmt.Default); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "Friend request accepted")
}

func declineFriendRequest(c echo.Context) error {
	var request models.DeclineFriendRequestRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := request.Validate(strfmt.Default); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "Friend request declined")
}

func getFriends(c echo.Context) error {
	return c.JSON(http.StatusOK, "Friends list retrieved")
}

func deleteFriend(c echo.Context) error {
	friendId := c.Param("friendId")
	if friendId == "" {
		return c.JSON(http.StatusBadRequest, "friendId parameter is required")
	}

	return c.JSON(http.StatusOK, "Friend deleted")
}

func InitSocialController(e *echo.Echo) {
	e.POST("/api/v1/social/profile", createProfile)
	e.GET("/api/v1/social/profiles/search", searchProfiles)
	e.POST("/api/v1/social/friend-request", createFriendRequest)
	e.GET("/api/v1/social/friend-requests", getFriendRequests)
	e.POST("/api/v1/social/friend-request/accept", acceptFriendRequest)
	e.POST("/api/v1/social/friend-request/decline", declineFriendRequest)
	e.GET("/api/v1/social/friends", getFriends)
	e.DELETE("/api/v1/social/friends/:friendId", deleteFriend)
}

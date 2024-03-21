package api

import (
	"net/http"

	"github.com/dayachettri/hotel-reservation/types"
	"github.com/labstack/echo/v4"
)

func HandleGetUsers(c echo.Context) error {
	u := types.User{
		FirstName: "James",
		LastName:  "At the watercooler",
	}

	return c.JSON(http.StatusOK, u)
}

func HandleGetUser(c echo.Context) error {
	return c.JSON(http.StatusOK, "James")
}

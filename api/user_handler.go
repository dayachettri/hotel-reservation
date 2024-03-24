package api

import (
	"net/http"
	"strconv"

	"github.com/dayachettri/hotel-reservation/db"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userStore db.UserStore
}

type Response struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUser(c echo.Context) error {
	id := c.Param("id")
	if _, err := strconv.Atoi(id); err != nil {
		return c.JSON(http.StatusBadRequest,
			&Response{
				Message:    "invalid id format only integers accepted",
				StatusCode: http.StatusBadRequest,
			})
	}

	user, err := h.userStore.GetUserByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			&Response{
				Message:    "no user found with this id",
				StatusCode: http.StatusNotFound})
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) HandleGetUsers(c echo.Context) error {
	users, err := h.userStore.GetUsers()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}

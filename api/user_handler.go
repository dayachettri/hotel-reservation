package api

import (
	"net/http"
	"strconv"

	"github.com/dayachettri/hotel-reservation/db"
	"github.com/dayachettri/hotel-reservation/types"
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

// TODO:Add request params validation
func (h *UserHandler) HandleUpdateUser(c echo.Context) error {
	id := c.Param("id")
	if _, err := strconv.Atoi(id); err != nil {
		return c.JSON(http.StatusBadRequest,
			&Response{
				Message:    "invalid id format only integers accepted",
				StatusCode: http.StatusBadRequest,
			})
	}

	params := &types.UpdateUserParams{}
	if err := c.Bind(&params); err != nil {
		return err
	}

	if err := params.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err := h.userStore.UpdateUser(c.Request().Context(), params, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &Response{
		Message:    "user updated successfully",
		StatusCode: http.StatusOK,
	})
}

func (h *UserHandler) HandleCreateUser(c echo.Context) error {
	params := &types.CreateUserParams{}
	if err := c.Bind(&params); err != nil {
		return err
	}

	if err := params.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user, err := types.NewUserfromParams(params)

	if err != nil {
		return err
	}

	createdUser, err := h.userStore.CreateUser(c.Request().Context(), user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, createdUser)
}

func (h *UserHandler) HandleDeleteUser(c echo.Context) error {
	id := c.Param("id")
	if _, err := strconv.Atoi(id); err != nil {
		return c.JSON(http.StatusBadRequest,
			&Response{
				Message:    "invalid id format only integers accepted",
				StatusCode: http.StatusBadRequest,
			})
	}

	if err := h.userStore.DeleteUser(c.Request().Context(), id); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &Response{
		Message:    "user deleted successfully",
		StatusCode: http.StatusOK,
	})
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

	user, err := h.userStore.GetUserByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound,
			&Response{
				Message:    "no user found with this id",
				StatusCode: http.StatusNotFound})
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) HandleGetUsers(c echo.Context) error {
	users, err := h.userStore.GetUsers(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}

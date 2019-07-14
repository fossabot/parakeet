package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/mohemohe/parakeet/server/models"
	"net/http"
	"strconv"
)

type (
	UserRequest struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	UserResponse struct {
		User  *models.User `json:"user"`
	}
)

func GetUsers(c echo.Context) error {
	limit, err := strconv.Atoi(c.Param("limit"))
	if err != nil {
		limit = 10
	}
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		page = 1
	}

	users := models.GetUsers(limit, page)
	if users == nil {
		panic("db error")
	}

	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	id := c.Param("id")
	user := models.GetUserById(id)
	if user == nil {
		panic("the user not found")
	}

	return c.JSON(http.StatusOK, UserResponse{
		User:  user,
	})
}

func CreateUser(c echo.Context) error {
	reqBody := new(UserRequest)
	if err := c.Bind(reqBody); err != nil {
		panic("bind error")
	}

	if user := models.GetUserByName(reqBody.Name); user != nil {
		panic("the user already registered")
	}

	user := &models.User{
		Name: reqBody.Name,
		Email: reqBody.Email,
		Password: reqBody.Password,
		Role: models.UserRole,
	}
	if err := models.UpsertUser(user); err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, UserResponse{
		User:  user,
	})
}

func UpdateUser(c echo.Context) error {
	reqBody := new(UserRequest)
	if err := c.Bind(reqBody); err != nil {
		panic("bind error")
	}

	id := c.Param("id")
	user := models.GetUserById(id)
	if user == nil {
		panic("the user not found")
	}

	nextUser := &models.User{
		Name: user.Name,
		Email: reqBody.Email,
		Password: reqBody.Password,
		Role: user.Role,
	}
	if err := models.UpsertUser(nextUser); err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, UserResponse{
		User:  user,
	})
}

func DeleteUser(c echo.Context) error {
	reqBody := new(UserRequest)
	if err := c.Bind(reqBody); err != nil {
		panic("bind error")
	}

	id := c.Param("id")
	user := models.GetUserById(id)
	if user == nil {
		panic("the user not found")
	}

	if err := models.DeleteUser(user); err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, struct{}{})
}
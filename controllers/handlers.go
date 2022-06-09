package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	db "server/main.go/database"
	"server/main.go/models"
	"strconv"
)

var database = db.Connect().Debug()

func CreateUser(c echo.Context) error {

	user := models.User{}

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	database.NewRecord(user)
	query := database.Create(&user)

	if query.Error != nil {
		return c.JSON(http.StatusInternalServerError, query.Error)
	}

	return c.JSON(http.StatusCreated, user)
}

func UpdateUserById(c echo.Context) error {

	user := models.User{}
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	query := database.Where("id =?", userId).Find(&user)

	if query != nil {
		return c.JSON(http.StatusInternalServerError, query.Error)
	}

	result := database.Save(&user)

	if result != nil {
		return c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	return c.JSON(http.StatusOK, user)
}

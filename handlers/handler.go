package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	db "server/main.go/database"
	"server/main.go/models"
	"strconv"
)

var database = db.Connect().Debug()

func CreateUser(c echo.Context) error {

	user := new(models.User)

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

	if query.Error != nil {
		return c.JSON(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}

	result := database.Save(&user)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	return c.JSON(http.StatusOK, user)
}

func UpdateUserByName(c echo.Context) error {

	user := models.User{}
	name := c.Param("name")

	query := database.Where("name =?", name).Find(&user)

	if query.Error != nil {
		return c.JSON(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}

	result := database.Save(&user)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	return c.JSON(http.StatusOK, user)
}

func DeleteUserByName(c echo.Context) error {

	user := new(models.User)
	name := c.Param("name")

	query := database.Where("name =?", name).Find(&user)

	if query.Error != nil {
		return c.JSON(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}

	result := database.Delete(&user)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	return c.JSON(http.StatusNoContent, result.RowsAffected)
}

func GetAllUsers(c echo.Context) error {

	var users []models.User
	query := database.Limit(100).Find(&users)

	if query.Error != nil {
		return c.JSON(http.StatusInternalServerError, query.Error)
	}

	return c.JSON(http.StatusOK, users)
}

func GetUserByName(c echo.Context) error {

	user := models.User{}
	name := c.Param("name")

	query := database.Where("name =?", name).Find(&user)

	if query.Error != nil {
		return c.JSON(http.StatusInternalServerError, query.Error)
	}

	result := database.Where(&models.User{Name: name}).Find(&user)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	return c.JSON(http.StatusOK, user)
}

func GetUserById(c echo.Context) error {

	user := models.User{}
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	query := database.Where("id =?", userId).Find(&user)

	if query.Error != nil {
		return c.JSON(http.StatusInternalServerError, query.Error)
	}

	return c.JSON(http.StatusOK, user)
}

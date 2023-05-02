package handlers

import (
	db "github.com/flakrimjusufi/grpc-with-rest/database"
	"github.com/flakrimjusufi/grpc-with-rest/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

var database = db.Connect().Debug()

// CreateUser -.
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

// UpdateUserByID -.
func UpdateUserByID(c echo.Context) error {

	user := models.User{}
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	query := database.Where("id =?", userID).Find(&user)

	if query.Error != nil {
		return c.JSON(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}

	result := database.Save(&user)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	return c.JSON(http.StatusOK, user)
}

// UpdateUserByName -.
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

// DeleteUserByName -.
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

// GetAllUsers -.
func GetAllUsers(c echo.Context) error {

	var users []models.User
	query := database.Limit(100).Find(&users)

	if query.Error != nil {
		return c.JSON(http.StatusInternalServerError, query.Error)
	}

	return c.JSON(http.StatusOK, users)
}

// GetUserByName -.
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

// GetUserByID -.
func GetUserByID(c echo.Context) error {

	user := models.User{}
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	query := database.Where("id =?", userID).Find(&user)

	if query.Error != nil {
		return c.JSON(http.StatusInternalServerError, query.Error)
	}

	return c.JSON(http.StatusOK, user)
}

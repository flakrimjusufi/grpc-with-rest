package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateUser(t *testing.T) {

	createUserJSON := `{"name":"Flakrim Jusufi","email":"flakrim.jusufi@gmail.com", "phoneNumber":"070123456"}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(createUserJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/create")

	// Assertions
	if assert.NoError(t, CreateUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestGetAllUsers(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/user")

	// Assertions
	if assert.NoError(t, GetAllUsers(c)) {
		if rec.Body.Len() == 0 {
			assert.Equal(t, http.StatusNoContent, rec.Code)
		} else {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, true, rec.Body.Len() > 0)
		}
	}
}

func TestGetUserById(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/user/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Assertions
	if assert.NoError(t, GetUserById(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, true, rec.Body.Len() > 0)
	}
}

func TestUpdateUserById(t *testing.T) {

	updateUserJSON := `{"name":"Flakrim Jusufi","email":"flakrimjusufi8@gmail.com", "phoneNumber":"070987654321"}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(updateUserJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/toDoList/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Assertions
	if assert.NoError(t, UpdateUserById(c)) {
		if rec.Code == http.StatusNotFound {
			assert.Errorf(t, fmt.Errorf(http.StatusText(http.StatusNotFound)), rec.Body.String())
		} else {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	}
}

func TestUpdateUserByName(t *testing.T) {

	updateUserJSON := `{"name":"Flakrim","email":"flakrimjusufi8@gmail.com", "phoneNumber":"070987654321"}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(updateUserJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/toDoList/:name")
	c.SetParamNames("name")
	c.SetParamValues("Flakrim")

	// Assertions
	if assert.NoError(t, UpdateUserByName(c)) {
		if rec.Code == http.StatusNotFound {
			assert.Equal(t, http.StatusText(http.StatusNotFound), "Not Found")
		} else {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	}
}

func TestDeleteUser(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/:id")
	c.SetParamNames("name")
	c.SetParamValues("Flakrim")

	// Assertions
	if assert.NoError(t, DeleteUserByName(c)) {
		if rec.Body.Len() == 0 {
			assert.Equal(t, http.StatusNotFound, rec.Code)
		} else {
			assert.Equal(t, http.StatusNoContent, rec.Code)
		}
	}
}

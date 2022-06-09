package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
	"server/main.go/handlers"
)

func main() {

	// Echo instance
	e := echo.New()
	if os.Getenv("SERVER_PORT") == "" {
		err := godotenv.Load() //Load .env file for local environment
		if err != nil {
			panic(err)
		}
	}

	// Middleware
	e.Use(middleware.Logger())

	e.GET("/user/", handlers.GetAllUsers)
	e.GET("/user/:name", handlers.GetUserByName)
	e.GET("/user/:id", handlers.GetUserById)
	e.POST("/user", handlers.CreateUser)
	e.PUT("/user/:id", handlers.UpdateUserById)
	e.PUT("/user/:name", handlers.UpdateUserByName)
	e.DELETE("/user/:name", handlers.DeleteUserByName)

}

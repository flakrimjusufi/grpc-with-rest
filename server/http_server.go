package main

import (
	"github.com/flakrimjusufi/grpc-with-rest/handlers"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
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
	e.Group("api/v1")

	e.GET("/user/", handlers.GetAllUsers)
	e.GET("/user/:name", handlers.GetUserByName)
	e.GET("/user/:id", handlers.GetUserByID)
	e.POST("/user", handlers.CreateUser)
	e.PUT("/user/:id", handlers.UpdateUserByID)
	e.PUT("/user/:name", handlers.UpdateUserByName)
	e.DELETE("/user/:name", handlers.DeleteUserByName)

}

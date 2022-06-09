package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
	"server/main.go/controllers"
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

	e.GET("/user/", controllers.GetAllUsers)
	e.GET("/user/:name", controllers.GetUserByName)
	e.GET("/user/:id", controllers.GetUserById)
	e.POST("/user", controllers.CreateUser)
	e.PUT("/user/:id", controllers.UpdateUserById)
	e.PUT("/user/:name", controllers.UpdateUserByName)
	e.DELETE("/user/:name", controllers.DeleteUserByName)

}

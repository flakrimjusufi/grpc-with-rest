package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"os"
	"server/main.go/controllers"
	db "server/main.go/database"
	"server/main.go/models"
)

var database = db.Connect().Debug()

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var users []models.User
	database.Limit(100).Find(&users)

	json.NewEncoder(w).Encode(users)
}

func getUserByName(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	name := vars["name"]
	var user models.User
	database.Where(&models.User{Name: name}).Find(&user)

	json.NewEncoder(w).Encode(user)
}

func getUserById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idFromMux := vars["id"]

	var user models.User
	database.Where("id = ?", idFromMux).Find(&user)
	json.NewEncoder(w).Encode(user)
}

func updateUserByName(w http.ResponseWriter, r *http.Request) {

	user := models.User{}
	vars := mux.Vars(r)
	name := vars["name"]

	database.Where("name =?", name).Find(&user)

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}

	database.Save(&user)
	json.NewEncoder(w).Encode(user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {

	user := models.User{}
	vars := mux.Vars(r)
	name := vars["name"]

	database.Where("name =?", name).Find(&user)
	database.Delete(&user)

	json.NewEncoder(w).Encode("User deleted successfully")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/user/updateByName/{name}", updateUserByName).Methods("POST")
	myRouter.HandleFunc("/user/delete/{name}", deleteUser).Methods("POST")
	myRouter.HandleFunc("/user/list", getAllUsers).Methods("GET", "OPTIONS")
	myRouter.HandleFunc("/user/getByName/{name}", getUserByName).Methods("GET")
	myRouter.HandleFunc("/user/getById/{id}", getUserById).Methods("GET")
	log.Fatal(http.ListenAndServe(":8088", myRouter))
}

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

	e.POST("/user", controllers.CreateUser)
	e.PUT("/user/:id", controllers.UpdateUserById)

}

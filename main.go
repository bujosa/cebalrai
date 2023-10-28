package main

import (
	"celbalrai/database"
	"celbalrai/user"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	database.ConnectDB()
	defer database.CloseDB()
	user.ModuleInit()

	userController := user.UserControllerInstance

	router := mux.NewRouter()

	router.HandleFunc("/users/{id}", userController.GetUserByID).Methods("GET")
	router.HandleFunc("/users", userController.GetUsers).Methods("GET")

	http.Handle("/", router)

	http.ListenAndServe(":8080", router)
}

package main

import (
	"celbalrai/database"
	"celbalrai/user"
	"net/http"
)

func main() {
	database.ConnectDB()
	defer database.CloseDB()
	user.ModuleInit()

	userController := user.UserControllerInstance

	http.HandleFunc("/users",  userController.GetUserByID)

	http.ListenAndServe(":8080", nil)
}

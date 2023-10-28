package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserController struct {
	userService *UserService
}

func (uc *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	parameter := vars["id"]

	id, err := strconv.Atoi(parameter)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	user, err := uc.userService.GetUserByID(r.Context(), id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (uc *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := uc.userService.GetUsers(r.Context())
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

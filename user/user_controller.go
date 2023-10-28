package user

import (
	"fmt"
	"net/http"
)

type UserController struct {
	userService *UserService
}

func (uc *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
    userID := r.URL.Query().Get("id")

    var id int
    _, _ = fmt.Sscanf(userID, "%d", &id)

    user, err := uc.userService.GetUserByID(r.Context(), id)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "User: %v", user)
}
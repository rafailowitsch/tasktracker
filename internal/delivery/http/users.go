package delivery

import (
	"encoding/json"
	"net/http"
	"tasktracker/pkg/log/sl"
)

type userInput struct {
	Name     string `json:"name" binding:"required, min=2, max=64"`
	Password string `json:"password" binding:"required, min=8, max=64"`
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	var input userInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.log.Error("Error decoding JSON", sl.Err(err))
		http.Error(w, "Error deconding JSON", http.StatusBadRequest)
		return
	}

	if input.Name == "" || input.Password == "" {
		h.log.Error("The name or password field is empty")
		http.Error(w, "One of the fields is empty", http.StatusBadRequest)
		return
	}

	_ = h.services.Users.SignUp(
		r.Context(),
		input.Name,
		input.Password,
	)

	// if err != nil {
	// }
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {

}

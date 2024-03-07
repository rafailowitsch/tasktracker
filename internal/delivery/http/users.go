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
	ctx := r.Context()

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

	err := h.services.Users.SignUp(
		ctx,
		input.Name,
		input.Password,
	)

	if err != nil {
		h.log.Error("Service layer error", sl.Err(err))
		http.Error(w, "Service layer error", http.StatusInternalServerError)
	}
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	var input userInput
	ctx := r.Context()

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

	res, err := h.services.Users.SignIn(
		ctx,
		input.Name,
		input.Password,
	)

	if err != nil {
		h.log.Error("Service layer error", sl.Err(err))
		http.Error(w, "Service layer error", http.StatusInternalServerError)
	}

	response := map[string]string{
		"AccessToken":  res.AccessToken,
		"RefreshToken": res.RefreshToken,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		h.log.Error("Error encoding JSON")
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

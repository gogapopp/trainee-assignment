package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gogapopp/trainee-assignment/internal/models"
	"github.com/gogapopp/trainee-assignment/internal/repository"
)

func (h *APIHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var req models.SignUpRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.service.SignUp(r.Context(), req)
	if err != nil {
		if errors.Is(err, repository.ErrUserExist) {
			http.Error(w, "user already exists", http.StatusBadRequest)
			return
		}
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (h *APIHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var req models.SignInRequest
	var resp models.SignInResponse
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := h.service.SignIn(r.Context(), req)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotExist) {
			http.Error(w, "user does not exists or wrong password", http.StatusBadRequest)
			return
		}
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	resp.Token = token
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
}

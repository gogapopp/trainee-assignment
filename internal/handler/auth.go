package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gogapopp/trainee-assignment/internal/models"
	"github.com/gogapopp/trainee-assignment/internal/repository"
	"github.com/gogapopp/trainee-assignment/internal/service"
)

func (h *APIHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	const op = "handler.auth.SignUp"
	var req models.SignUpRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.logger.Errorf("%s: %w", op, err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	err = h.authService.SignUp(r.Context(), req)
	if err != nil {
		h.logger.Errorf("%s: %w", op, err)
		if errors.Is(err, repository.ErrUserExist) {
			http.Error(w, "user already exists", http.StatusConflict)
			return
		}
		if errors.As(err, &valErr) {
			http.Error(w, "password and username is required field", http.StatusBadRequest)
			return
		}
		if errors.Is(err, service.ErrUndefinedRole) {
			http.Error(w, "undefined role (available roles: admin, user)", http.StatusBadRequest)
			return
		}
		internalServerErrorHandlerFunc(w)
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *APIHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	const op = "handler.auth.SignIn"
	var req models.SignInRequest
	var resp models.SignInResponse
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.logger.Errorf("%s: %w", op, err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	token, err := h.authService.SignIn(r.Context(), req)
	if err != nil {
		h.logger.Errorf("%s: %w", op, err)
		if errors.Is(err, repository.ErrUserNotExist) {
			http.Error(w, "user does not exist or wrong password", http.StatusBadRequest)
			return
		}
		if errors.As(err, &valErr) {
			http.Error(w, "password and username is required field", http.StatusBadRequest)
			return
		}
		internalServerErrorHandlerFunc(w)
	}
	resp.Token = token
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		h.logger.Errorf("%s: %w", op, err)
		internalServerErrorHandlerFunc(w)
	}
}

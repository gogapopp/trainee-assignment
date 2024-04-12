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
		badRequestHandlerFunc(w, "bad request", http.StatusBadRequest)
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
			badRequestHandlerFunc(w, "password and username is required field", http.StatusBadRequest)
			return
		}
		if errors.Is(err, service.ErrUndefinedRole) {
			badRequestHandlerFunc(w, "undefined role (available roles: admin, user)", http.StatusBadRequest)
			return
		}
		internalServerErrorHandlerFunc(w)
		return
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
		badRequestHandlerFunc(w, "bad request", http.StatusBadRequest)
		return
	}
	token, err := h.authService.SignIn(r.Context(), req)
	if err != nil {
		h.logger.Errorf("%s: %w", op, err)
		if errors.Is(err, repository.ErrUserNotExist) {
			badRequestHandlerFunc(w, "user does not exist or wrong password", http.StatusNotFound)
			return
		}
		if errors.As(err, &valErr) {
			badRequestHandlerFunc(w, "password and username is required field", http.StatusBadRequest)
			return
		}
		internalServerErrorHandlerFunc(w)
		return
	}
	resp.Token = token
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		h.logger.Errorf("%s: %w", op, err)
		internalServerErrorHandlerFunc(w)
		return
	}
}

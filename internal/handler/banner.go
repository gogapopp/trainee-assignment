package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gogapopp/trainee-assignment/internal/handler/middlewares"
	"github.com/gogapopp/trainee-assignment/internal/models"
)

func (h *APIHandler) GetBanner(w http.ResponseWriter, r *http.Request, params GetBannerParams) {
	// const op = "handler.banner.GetBanner"
	// ctx := r.Context()
}

// Создание нового баннера
// (POST /banner)
func (h *APIHandler) PostBanner(w http.ResponseWriter, r *http.Request, params PostBannerParams) {
	const op = "handler.banner.PostBanner"
	ctx := r.Context()
	var req models.PostBannerRequest
	var resp models.PostBannerResponse
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.logger.Errorf("%s: %w", op, err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	bannerId, err := h.bannerService.SaveBanner(ctx, req)
	if err != nil {
		h.logger.Errorf("%s: %w", op, err)
		if errors.As(err, &valErr) {
			http.Error(w, "fields in the JSON cant be empty", http.StatusBadRequest)
			return
		}
		internalServerErrorHandlerFunc(w)
	}
	resp.BannerID = bannerId
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		h.logger.Errorf("%s: %w", op, err)
		internalServerErrorHandlerFunc(w)
	}
}

// Удаление баннера по идентификатору
// (DELETE /banner/{id})
func (h *APIHandler) DeleteBannerId(w http.ResponseWriter, r *http.Request, id int, params DeleteBannerIdParams) {
	// const op = "handler.banner.DeleteBannerId"
	// ctx := r.Context()
}

// Обновление содержимого баннера
// (PATCH /banner/{id})
func (h *APIHandler) PatchBannerId(w http.ResponseWriter, r *http.Request, id int, params PatchBannerIdParams) {
	// const op = "handler.banner.PatchBanner"
	// ctx := r.Context()
}

// Получение баннера для пользователя
// (GET /user_banner)
func (h *APIHandler) GetUserBanner(w http.ResponseWriter, r *http.Request, params GetUserBannerParams) {
	userRole := r.Context().Value(middlewares.UserRoleKey)
	fmt.Println(r.Context().Value(middlewares.UserIDKey), userRole)
	if ia, ok := userRole.(string); ok {
		println(ia)
	}
	fmt.Println(params.FeatureId, params.TagId, params.UseLastRevision)
	w.WriteHeader(http.StatusConflict)
	_, _ = w.Write([]byte("bb"))
}

package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gogapopp/trainee-assignment/internal/lib/admin"
	"github.com/gogapopp/trainee-assignment/internal/models"
	"github.com/gogapopp/trainee-assignment/internal/repository"
	"github.com/gogapopp/trainee-assignment/internal/service"
)

// Получение всех баннеров c фильтрацией по фиче и/или тегу
// (GET /banner)
func (h *APIHandler) GetBanner(w http.ResponseWriter, r *http.Request, params GetBannerParams) {
	const op = "handler.banner.GetBanner"
	ctx := r.Context()
	if ok := admin.IsAdmin(ctx); !ok {
		http.Error(w, "you dont have premissions", http.StatusForbidden)
		return
	}
	banners, err := h.bannerService.GetBanners(ctx, models.BannersRequest{
		FeatureId: params.FeatureId,
		TagId:     params.TagId,
		Limit:     params.Limit,
		Offset:    params.Offset,
	})
	if err != nil {
		h.logger.Errorf("%s: %w", op, err)
		if errors.Is(err, service.ErrBannersNotExist) {
			badRequestHandlerFunc(w, "banners dont exist", http.StatusNotFound)
			return
		}
		internalServerErrorHandlerFunc(w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(banners); err != nil {
		h.logger.Errorf("%s: %w", op, err)
		internalServerErrorHandlerFunc(w)
		return
	}

}

// Создание нового баннера
// (POST /banner)
func (h *APIHandler) PostBanner(w http.ResponseWriter, r *http.Request, params PostBannerParams) {
	const op = "handler.banner.PostBanner"
	ctx := r.Context()
	if ok := admin.IsAdmin(ctx); !ok {
		http.Error(w, "you dont have premissions", http.StatusForbidden)
		return
	}
	var req models.PostBannerRequest
	var resp models.PostBannerResponse
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.logger.Errorf("%s: %w", op, err)
		badRequestHandlerFunc(w, "bad request", http.StatusBadRequest)
		return
	}
	bannerId, err := h.bannerService.SaveBanner(ctx, req)
	if err != nil {
		h.logger.Errorf("%s: %w", op, err)
		if errors.As(err, &valErr) {
			badRequestHandlerFunc(w, "fields in the JSON cant be empty", http.StatusBadRequest)
			return
		}
		internalServerErrorHandlerFunc(w)
		return
	}
	resp.BannerID = bannerId
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		h.logger.Errorf("%s: %w", op, err)
		internalServerErrorHandlerFunc(w)
		return
	}
}

// Удаление баннера по идентификатору
// (DELETE /banner/{id})
func (h *APIHandler) DeleteBannerId(w http.ResponseWriter, r *http.Request, id int, params DeleteBannerIdParams) {
	const op = "handler.banner.DeleteBannerId"
	ctx := r.Context()
	if ok := admin.IsAdmin(ctx); !ok {
		http.Error(w, "you dont have premissions", http.StatusForbidden)
		return
	}
	if id < 0 {
		badRequestHandlerFunc(w, "bad request", http.StatusBadRequest)
		return
	}
	err := h.bannerService.DeleteBanner(ctx, id)
	if err != nil {
		h.logger.Errorf("%s: %w", op, err)
		if errors.Is(err, repository.ErrBannerNotExist) {
			badRequestHandlerFunc(w, "banner does not exist", http.StatusNotFound)
			return
		}
		internalServerErrorHandlerFunc(w)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// Обновление содержимого баннера
// (PATCH /banner/{id})
func (h *APIHandler) PatchBannerId(w http.ResponseWriter, r *http.Request, id int, params PatchBannerIdParams) {
	const op = "handler.banner.PatchBanner"
	ctx := r.Context()
	if ok := admin.IsAdmin(ctx); !ok {
		http.Error(w, "you dont have premissions", http.StatusForbidden)
		return
	}
	var req models.PatchBanner
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.logger.Errorf("%s: %w", op, err)
		badRequestHandlerFunc(w, "bad request", http.StatusBadRequest)
		return
	}
	err = h.bannerService.PatchBannerId(ctx, id, req)
	if err != nil {
		h.logger.Errorf("%s: %w", op, err)
		if errors.Is(err, repository.ErrBannerNotExist) {
			badRequestHandlerFunc(w, repository.ErrBannerNotExist.Error(), http.StatusNotFound)
			return
		}
		if errors.Is(err, repository.ErrNoFieldsToUpdate) {
			badRequestHandlerFunc(w, repository.ErrNoFieldsToUpdate.Error(), http.StatusBadRequest)
			return
		}
		internalServerErrorHandlerFunc(w)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Получение баннера для пользователя
// (GET /user_banner)
func (h *APIHandler) GetUserBanner(w http.ResponseWriter, r *http.Request, params GetUserBannerParams) {
	const op = "handler.banner.GetUserBanner"
	ctx := r.Context()
	var resp models.Content
	banner, err := h.bannerService.GetUserBanner(ctx, models.UserBannerRequest{
		FeatureId:       params.FeatureId,
		TagId:           params.TagId,
		UseLastRevision: params.UseLastRevision,
	})
	if err != nil {
		h.logger.Errorf("%s: %w", op, err)
		if errors.Is(err, service.ErrBannerUnactive) {
			badRequestHandlerFunc(w, service.ErrBannerUnactive.Error(), http.StatusForbidden)
			return
		}
		if errors.Is(err, repository.ErrBannerNotExist) {
			badRequestHandlerFunc(w, "banner does not exist", http.StatusNotFound)
			return
		}
		internalServerErrorHandlerFunc(w)
		return
	}
	resp.Text = banner.Content.Text
	resp.Title = banner.Content.Title
	resp.URL = banner.Content.URL
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		h.logger.Errorf("%s: %w", op, err)
		internalServerErrorHandlerFunc(w)
		return
	}
}

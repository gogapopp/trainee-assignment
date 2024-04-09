package handler

import (
	"fmt"
	"net/http"

	"github.com/gogapopp/trainee-assignment/internal/handler/middlewares"
)

func (h *APIHandler) GetBanner(w http.ResponseWriter, r *http.Request, params GetBannerParams) {

}

// Создание нового баннера
// (POST /banner)
func (h *APIHandler) PostBanner(w http.ResponseWriter, r *http.Request, params PostBannerParams) {

}

// Удаление баннера по идентификатору
// (DELETE /banner/{id})
func (h *APIHandler) DeleteBannerId(w http.ResponseWriter, r *http.Request, id int, params DeleteBannerIdParams) {

}

// Обновление содержимого баннера
// (PATCH /banner/{id})
func (h *APIHandler) PatchBannerId(w http.ResponseWriter, r *http.Request, id int, params PatchBannerIdParams) {

}

// Получение баннера для пользователя
// (GET /user_banner)
func (h *APIHandler) GetUserBanner(w http.ResponseWriter, r *http.Request, params GetUserBannerParams) {
	fmt.Println(r.Context().Value(middlewares.UserIDKey), r.Context().Value(middlewares.UserIsAdminKey))
	fmt.Println(params.FeatureId, params.TagId, params.UseLastRevision)
	w.WriteHeader(http.StatusConflict)
	_, _ = w.Write([]byte("bb"))
}

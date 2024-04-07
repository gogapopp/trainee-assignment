package handler

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

type (
	Service interface {
	}

	APIHandler struct {
		logger  *zap.SugaredLogger
		service Service
	}
)

func New(logger *zap.SugaredLogger, service Service) *APIHandler {
	return &APIHandler{
		logger:  logger,
		service: service,
	}
}

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
	fmt.Println(params)
	w.WriteHeader(http.StatusConflict)
	w.Write([]byte("bb"))
}

func (h *APIHandler) SignUp(w http.ResponseWriter, r *http.Request) {

}

func (h *APIHandler) SignIn(w http.ResponseWriter, r *http.Request) {

}

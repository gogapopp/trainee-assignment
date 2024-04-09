package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gogapopp/trainee-assignment/internal/handler/middlewares"
	"github.com/gogapopp/trainee-assignment/internal/models"
	"go.uber.org/zap"
)

type (
	Service interface {
		SignUp(ctx context.Context, user models.SignUpRequest) error
		SignIn(ctx context.Context, user models.SignInRequest) (string, error)
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
	fmt.Println(r.Context().Value(middlewares.UserIDKey), r.Context().Value(middlewares.UserRoleKey))
	fmt.Println(params.FeatureId, params.TagId, params.UseLastRevision)
	w.WriteHeader(http.StatusConflict)
	w.Write([]byte("bb"))
}

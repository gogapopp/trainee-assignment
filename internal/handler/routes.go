package handler

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	mw "github.com/gogapopp/trainee-assignment/internal/handler/middlewares"
	"go.uber.org/zap"

	"github.com/go-chi/chi/v5"
)

func Routes(logger *zap.SugaredLogger, addr string, service Service) *http.Server {
	APIHandler := New(logger, service)

	r := chi.NewRouter()

	// вспомогательные ручки (по сути мы считаем, что юзер уже получил токен откуда то, например из сервиса авторизации)
	r.Get("/signup", APIHandler.SignUp)
	r.Get("/signin", APIHandler.SignIn)

	// создаем сервер на основе спецификации
	middlewares := []MiddlewareFunc{mw.AuthMiddleware, middleware.Logger, middleware.RequestID}
	chiOptions := ChiServerOptions{
		BaseRouter:  r,
		Middlewares: middlewares,
	}
	h := HandlerWithOptions(APIHandler, chiOptions)

	srv := &http.Server{
		Addr:    addr,
		Handler: h,
	}
	return srv
}

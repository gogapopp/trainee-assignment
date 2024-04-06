package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gogapopp/trainee-assignment/internal/handler"
	mw "github.com/gogapopp/trainee-assignment/internal/handler/middlewares"
)

func main() {
	r := chi.NewRouter()

	APIHandler := handler.New()

	// вспомогательные ручки (по сути мы считаем, что юзер уже получил токен откуда то, например из сервиса авторизации)
	r.Get("/signup", APIHandler.SignUp)
	r.Get("/signin", APIHandler.SignIn)

	// проверяем все синтаксические проверки на уровне спецификации
	middlewares := []handler.MiddlewareFunc{mw.AuthMiddleware, middleware.Logger}
	chiOptions := handler.ChiServerOptions{
		BaseRouter:  r,
		Middlewares: middlewares,
	}
	h := handler.HandlerWithOptions(APIHandler, chiOptions)

	srv := http.Server{
		Addr:    ":8080",
		Handler: h,
	}
	srv.ListenAndServe()
}

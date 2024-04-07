package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gogapopp/trainee-assignment/internal/config"
	"github.com/gogapopp/trainee-assignment/internal/handler"
	mw "github.com/gogapopp/trainee-assignment/internal/handler/middlewares"
	"github.com/gogapopp/trainee-assignment/internal/lib/logger"
	"github.com/gogapopp/trainee-assignment/internal/repository/postgres"
	"github.com/gogapopp/trainee-assignment/internal/service"
)

func main() {
	ctx := context.Background()
	logger, err := logger.New()
	if err != nil {
		logger.Fatal(err)
	}
	config, err := config.New(".env")
	if err != nil {
		logger.Fatal(err)
	}
	repo, err := postgres.New(config.PGConfig.DSN)
	if err != nil {
		logger.Fatal(err)
	}
	defer repo.Close(ctx)

	service := service.New(logger, repo)
	APIHandler := handler.New(logger, service)

	r := chi.NewRouter()

	// вспомогательные ручки (по сути мы считаем, что юзер уже получил токен откуда то, например из сервиса авторизации)
	r.Get("/signup", APIHandler.SignUp)
	r.Get("/signin", APIHandler.SignIn)

	// проверяем все синтаксические проверки на уровне спецификации
	middlewares := []handler.MiddlewareFunc{mw.AuthMiddleware, middleware.Logger, middleware.RequestID}
	chiOptions := handler.ChiServerOptions{
		BaseRouter:  r,
		Middlewares: middlewares,
	}
	h := handler.HandlerWithOptions(APIHandler, chiOptions)

	srv := &http.Server{
		Addr:    config.HTTPConfig.Addr,
		Handler: h,
	}

	go func() {
		logger.Infof("Running the server at: %s", config.HTTPConfig.Addr)
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			logger.Fatalf("error to start the server: %w", err)
		}
	}()

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-sigint

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("error shutdown the server: %w", err)
	}
}

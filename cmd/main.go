package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gogapopp/trainee-assignment/internal/config"
	"github.com/gogapopp/trainee-assignment/internal/handler"
	"github.com/gogapopp/trainee-assignment/internal/logger"
	"github.com/gogapopp/trainee-assignment/internal/repository/cache"
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
	defer repo.Close()

	cache := cache.New()

	authService := service.NewAuthService(config.JWT_SECRET, config.PASS_SECRET, logger, repo)
	bannerService := service.NewBannerService(logger, repo, cache)

	server := handler.Routes(config.HTTPConfig.Addr, logger, authService, bannerService)

	go func() {
		logger.Infof("Running the server at: %s", config.HTTPConfig.Addr)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			logger.Fatalf("error to start the server: %w", err)
		}
	}()

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-sigint

	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalf("error shutdown the server: %w", err)
	}
}

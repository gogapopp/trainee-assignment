package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/middleware"
	mw "github.com/gogapopp/trainee-assignment/internal/handler/middlewares"
	"go.uber.org/zap"

	"github.com/go-chi/chi/v5"
)

func Routes(addr string, logger *zap.SugaredLogger, authService authService, bannerService bannerService) *http.Server {
	APIHandler := New(logger, authService, bannerService)

	r := chi.NewRouter()

	r.Post("/signup", APIHandler.SignUp)
	r.Post("/signin", APIHandler.SignIn)

	middlewares := []MiddlewareFunc{mw.AuthMiddleware, middleware.Logger, middleware.RequestID}
	chiOptions := ChiServerOptions{
		BaseRouter:       r,
		Middlewares:      middlewares,
		ErrorHandlerFunc: errorHandlerFunc,
	}
	h := HandlerWithOptions(APIHandler, chiOptions)

	srv := &http.Server{
		Addr:    addr,
		Handler: h,
	}
	return srv
}

type errorResponse struct {
	Error string `json:"error"`
}

func errorHandlerFunc(w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	if err := json.NewEncoder(w).Encode(errorResponse{Error: err.Error()}); err != nil {
		internalServerErrorHandlerFunc(w)
	}
}

func internalServerErrorHandlerFunc(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(errorResponse{Error: "something went wrong"}); err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
}

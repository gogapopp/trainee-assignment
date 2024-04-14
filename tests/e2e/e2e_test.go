package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os/exec"
	"testing"
	"time"

	"github.com/gogapopp/trainee-assignment/internal/config"
	"github.com/gogapopp/trainee-assignment/internal/handler"
	"github.com/gogapopp/trainee-assignment/internal/handler/middlewares"
	"github.com/gogapopp/trainee-assignment/internal/logger"
	"github.com/gogapopp/trainee-assignment/internal/models"
	"github.com/gogapopp/trainee-assignment/internal/repository/cache"
	"github.com/gogapopp/trainee-assignment/internal/repository/postgres"
	"github.com/gogapopp/trainee-assignment/internal/service"
	"github.com/gogapopp/trainee-assignment/tests/e2e/internal/util"
	"github.com/stretchr/testify/assert"
)

// TestAPIHandlerBannerSaveAndGet тестирует сценарий сохранения и получения банера новым пользователем
func TestAPIHandlerBannerSaveAndGet(t *testing.T) {
	cmd := exec.Command("docker-compose", "up", "-d", "pg-local")
	err := cmd.Run()
	time.Sleep(time.Second * 3)
	assert.NoError(t, err)
	logger, err := logger.New()
	assert.NoError(t, err)
	config, err := config.New(".env")
	assert.NoError(t, err)
	repo, err := postgres.New(config.PGConfig.DSN)
	assert.NoError(t, err)
	defer repo.Close()

	cache := cache.New()

	authService := service.NewAuthService(config.JWT_SECRET, config.PASS_SECRET, logger, repo)
	bannerService := service.NewBannerService(logger, repo, cache)

	h := handler.New(logger, authService, bannerService)

	userName := fmt.Sprintf("user%d", util.GenerateRandomNumber())
	userPass := fmt.Sprintf("pass%d", util.GenerateRandomNumber())
	tests := []struct {
		name          string
		signUpReq     models.SignUpRequest
		signInReq     models.SignInRequest
		postBannerReq models.PostBannerRequest
		wantStatus    int
	}{
		{
			name: "Valid SignUp and SignIn",
			signUpReq: models.SignUpRequest{
				Username:     userName,
				PasswordHash: userPass,
				Role:         "admin",
			},
			signInReq: models.SignInRequest{
				Username:     userName,
				PasswordHash: userPass,
			},
			postBannerReq: models.PostBannerRequest{
				TagIDs:    []int{1, 2, 3},
				FeatureID: 1,
				Content: models.Content{
					Title: "Test Title",
					Text:  "Test Text",
					URL:   "https://test.url",
				},
				IsActive: true,
			},
			wantStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// регистрация
			signUpBody, _ := json.Marshal(tt.signUpReq)
			signUpReq, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(signUpBody))
			signUpResp := httptest.NewRecorder()
			h.SignUp(signUpResp, signUpReq)
			assert.Equal(t, http.StatusCreated, signUpResp.Code)
			// логин для получения токена
			signInBody, _ := json.Marshal(tt.signInReq)
			signInReq, _ := http.NewRequest(http.MethodPost, "/signin", bytes.NewBuffer(signInBody))
			signInResp := httptest.NewRecorder()
			h.SignIn(signInResp, signInReq)
			// проверяем ответ от signin
			assert.Equal(t, tt.wantStatus, signInResp.Code)
			var resp models.SignInResponse
			json.Unmarshal(signInResp.Body.Bytes(), &resp)
			assert.NotEmpty(t, resp.Token)
			// cоздание баннера
			postBannerBody, _ := json.Marshal(tt.postBannerReq)
			postBannerReq, _ := http.NewRequest(http.MethodPost, "/banner", bytes.NewBuffer(postBannerBody))
			postBannerReq.Header.Set("token", resp.Token) // Добавление токена в заголовок
			postBannerResp := httptest.NewRecorder()
			// новый хендлер с middleware
			handlerWithMiddleware := middlewares.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				h.PostBanner(w, r, handler.PostBannerParams{})
			}))
			handlerWithMiddleware.ServeHTTP(postBannerResp, postBannerReq)
			// проверяем ответ
			assert.Equal(t, http.StatusCreated, postBannerResp.Code)
			var postBannerRespModel models.PostBannerResponse
			json.Unmarshal(postBannerResp.Body.Bytes(), &postBannerRespModel)
			assert.NotEmpty(t, postBannerRespModel.BannerID)
			// получение баннера
			getBannerReq, _ := http.NewRequest(http.MethodGet, "/user_banner?tag_id=1&feature_id=1", nil)
			getBannerReq.Header.Set("token", resp.Token) // Добавление токена в заголовок
			getBannerResp := httptest.NewRecorder()
			// новый хендлер с middleware
			handlerWithMiddleware = middlewares.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				h.GetUserBanner(w, r, handler.GetUserBannerParams{
					FeatureId: 1,
					TagId:     1,
				})
			}))
			handlerWithMiddleware.ServeHTTP(getBannerResp, getBannerReq)
			// проверяем ответ
			assert.Equal(t, http.StatusOK, getBannerResp.Code)
			var getBannerRespModel models.Content
			json.Unmarshal(getBannerResp.Body.Bytes(), &getBannerRespModel)
			assert.NotEmpty(t, getBannerRespModel)
			assert.Equal(t, getBannerRespModel.Title, tt.postBannerReq.Content.Title)
			assert.Equal(t, getBannerRespModel.Text, tt.postBannerReq.Content.Text)
			assert.Equal(t, getBannerRespModel.URL, tt.postBannerReq.Content.URL)
		})
	}
	// отключаемся
	cmd = exec.Command("make", "stop")
	err = cmd.Run()
	assert.NoError(t, err)
}

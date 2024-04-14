package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gogapopp/trainee-assignment/internal/logger"
	"github.com/gogapopp/trainee-assignment/internal/models"
	"github.com/gogapopp/trainee-assignment/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockBannerService struct {
	mock.Mock
}

func (m *MockBannerService) SaveBanner(ctx context.Context, banner models.PostBannerRequest) (int, error) {
	args := m.Called(ctx, banner)
	return args.Int(0), args.Error(1)
}

func (m *MockBannerService) GetUserBanner(ctx context.Context, params models.UserBannerRequest) (models.UserBannerResponse, error) {
	args := m.Called(ctx, params)
	return args.Get(0).(models.UserBannerResponse), args.Error(1)
}

func (m *MockBannerService) GetBanners(ctx context.Context, params models.BannersRequest) ([]models.BannersResponse, error) {
	args := m.Called(ctx, params)
	return args.Get(0).([]models.BannersResponse), args.Error(1)
}

func (m *MockBannerService) DeleteBanner(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockBannerService) PatchBannerId(ctx context.Context, id int, banner models.PatchBanner) error {
	args := m.Called(ctx, id, banner)
	return args.Error(0)
}

func (m *MockBannerService) DeleteBannerByFeatureId(ctx context.Context, featureId int) error {
	args := m.Called(ctx, featureId)
	return args.Error(0)
}

type MockAuthService struct {
	mock.Mock
}

func (m *MockAuthService) SignUp(ctx context.Context, user models.SignUpRequest) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockAuthService) SignIn(ctx context.Context, user models.SignInRequest) (string, error) {
	args := m.Called(ctx, user)
	return args.String(0), args.Error(1)
}

func TestSignUp(t *testing.T) {
	tests := []struct {
		name           string
		body           models.SignUpRequest
		expectedStatus int
		wantErr        bool
	}{
		{
			name: "Successful SignUp",
			body: models.SignUpRequest{
				Username:     "testuser",
				PasswordHash: "testpassword",
				Role:         "user",
			},
			wantErr:        false,
			expectedStatus: http.StatusCreated,
		},
		{
			name: "User Already Exists",
			body: models.SignUpRequest{
				Username:     "existinguser",
				PasswordHash: "testpassword",
				Role:         "user",
			},
			wantErr:        true,
			expectedStatus: http.StatusConflict,
		},
	}

	logger, err := logger.New()
	assert.NoError(t, err)
	s := &MockAuthService{}

	h := New(logger, s, &MockBannerService{})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reqBody, _ := json.Marshal(tt.body)
			req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(reqBody))
			rec := httptest.NewRecorder()

			if tt.wantErr {
				s.On("SignUp", mock.Anything, models.SignUpRequest{
					Username:     "existinguser",
					PasswordHash: "testpassword",
					Role:         "user",
				}).Return(repository.ErrUserExist)
			} else {
				s.On("SignUp", mock.Anything, models.SignUpRequest{
					Username:     "testuser",
					PasswordHash: "testpassword",
					Role:         "user",
				}).Return(nil)
			}

			h.SignUp(rec, req)

			assert.Equal(t, tt.expectedStatus, rec.Code)
		})
	}
}

// func TestSignIn(t *testing.T) {
// 	tests := []struct {
// 		name           string
// 		body           models.SignInRequest
// 		wantErr        bool
// 		expectedStatus int
// 	}{
// 		{
// 			name: "Successful SignIn",
// 			body: models.SignInRequest{
// 				Username:     "testuser",
// 				PasswordHash: "testpassword",
// 			},
// 			wantErr:        false,
// 			expectedStatus: http.StatusOK,
// 		},
// 		{
// 			name: "User Does Not Exist",
// 			body: models.SignInRequest{
// 				Username:     "nonexistentuser",
// 				PasswordHash: "testpassword",
// 			},
// 			wantErr:        true,
// 			expectedStatus: http.StatusNotFound,
// 		},
// 	}

// 	logger, err := logger.New()
// 	assert.NoError(t, err)
// 	s := &MockAuthService{}

// 	h := New(logger, s, &MockBannerRepo{})

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			reqBody, _ := json.Marshal(tt.body)
// 			req, _ := http.NewRequest(http.MethodPost, "/signin", bytes.NewBuffer(reqBody))
// 			rec := httptest.NewRecorder()

// 			if tt.wantErr {
// 				s.On("SignIn", mock.Anything, models.SignInRequest{
// 					Username:     tt.body.Username,
// 					PasswordHash: tt.body.PasswordHash,
// 				}).Return("testtoken", repository.ErrUserNotExist)
// 			} else {
// 				s.On("SignIn", mock.Anything, mock.Anything).Return("testtoken", nil)
// 			}

// 			h.SignIn(rec, req)

// 			t.Log(rec.Body, rec.Code)
// 			assert.Equal(t, tt.expectedStatus, rec.Code)
// 		})
// 	}
// }

package service

import (
	"context"
	"crypto/sha256"
	"fmt"

	"github.com/gogapopp/trainee-assignment/internal/lib/jwt"
	"github.com/gogapopp/trainee-assignment/internal/models"
)

func (s *service) SignUp(ctx context.Context, user models.SignUpRequest) error {
	user.PasswordHash = s.generatePasswordHash(user.PasswordHash)
	return s.repo.SignUp(ctx, user)
}

func (s *service) SignIn(ctx context.Context, user models.SignInRequest) (string, error) {
	const op = "service.auth.SignIn"
	user.PasswordHash = s.generatePasswordHash(user.PasswordHash)
	userId, userRole, err := s.repo.SignIn(ctx, user)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	token, err := jwt.GenerateJWTToken(s.jwtSecret, userId, userRole, user.Username, user.PasswordHash)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return token, nil
}

func (s *service) generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(s.passSecret)))
}

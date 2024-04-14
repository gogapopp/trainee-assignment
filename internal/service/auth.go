package service

import (
	"context"
	"crypto/sha256"
	"fmt"

	"github.com/gogapopp/trainee-assignment/internal/lib/jwt"
	"github.com/gogapopp/trainee-assignment/internal/models"
)

func (a *authService) SignUp(ctx context.Context, user models.SignUpRequest) error {
	const op = "service.auth.SignUp"
	err := a.validator.Struct(user)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	err = validateRole(user.Role)
	if err != nil {
		return fmt.Errorf("%s: %w", op, ErrUndefinedRole)
	}
	user.PasswordHash = a.generatePasswordHash(user.PasswordHash)
	err = a.authRepo.SignUp(ctx, user)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (a *authService) SignIn(ctx context.Context, user models.SignInRequest) (string, error) {
	const op = "service.auth.SignIn"
	err := a.validator.Struct(user)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	user.PasswordHash = a.generatePasswordHash(user.PasswordHash)
	userRole, err := a.authRepo.SignIn(ctx, user)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	token, err := jwt.GenerateJWTToken(a.jwtSecret, userRole, user.Username, user.PasswordHash)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return token, nil
}

func (a *authService) generatePasswordHash(password string) string {
	hash := sha256.New()
	_, _ = hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(a.passSecret)))
}

func validateRole(role string) error {
	if role != "admin" && role != "user" {
		return ErrUndefinedRole
	}
	return nil
}

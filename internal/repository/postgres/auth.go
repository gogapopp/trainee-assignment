package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/gogapopp/trainee-assignment/internal/models"
	"github.com/gogapopp/trainee-assignment/internal/repository"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (s *storage) SignUp(ctx context.Context, user models.SignUpRequest) error {
	const (
		op    = "postgres.auth.SignUp"
		query = "INSERT INTO users (username, password_hash, role) VALUES ($1, $2, $3);"
	)
	_, err := s.db.Exec(ctx, query, user.Username, user.PasswordHash, user.Role)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				return fmt.Errorf("%s: %w", op, repository.ErrUserExist)
			}
		}
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *storage) SignIn(ctx context.Context, user models.SignInRequest) (string, error) {
	const (
		op    = "postgres.auth.SignIn"
		query = "SELECT role FROM users WHERE username=$1 AND password_hash=$2;"
	)
	var (
		role string
	)
	row := s.db.QueryRow(ctx, query, user.Username, user.PasswordHash)
	err := row.Scan(&role)
	if err != nil {
		if err == pgx.ErrNoRows {
			return "", fmt.Errorf("%s: %w", op, repository.ErrUserNotExist)
		}
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return role, nil
}

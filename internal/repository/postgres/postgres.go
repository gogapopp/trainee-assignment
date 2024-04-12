package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type storage struct {
	db *pgxpool.Pool
}

func New(dsn string) (*storage, error) {
	const op = "postgres.postgres.New"
	ctx := context.Background()
	db, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	err = db.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	tx, err := db.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer tx.Rollback(ctx)
	_, err = tx.Exec(ctx, `
	CREATE TABLE IF NOT EXISTS users (
		user_id SERIAL PRIMARY KEY,
		username TEXT,
		password_hash TEXT,
		role TEXT
	);
	CREATE UNIQUE INDEX IF NOT EXISTS username_idx ON users(username);
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	_, err = tx.Exec(ctx, `
	CREATE TABLE IF NOT EXISTS banners (
		banner_id SERIAL PRIMARY KEY,
		tag_ids INTEGER[],
		feature_id INTEGER,
		banner_data JSONB,
		is_active BOOLEAN,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	_, err = tx.Exec(ctx, `
	CREATE TABLE IF NOT EXISTS banners_tags (
		banner_id INTEGER,
		tag_id INTEGER
	);
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &storage{
		db: db,
	}, tx.Commit(ctx)
}

func (s *storage) Close() {
	s.db.Close()
}

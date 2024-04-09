package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type storage struct {
	db *pgxpool.Pool
}

func New(dsn string) (*storage, error) {
	ctx := context.Background()
	db, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping(ctx)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(ctx, `
	CREATE TABLE IF NOT EXISTS users (
		user_id SERIAL PRIMARY KEY,
		username TEXT,
		password_hash TEXT,
		is_admin BOOLEAN
	);
	CREATE UNIQUE INDEX IF NOT EXISTS username_idx ON users(username);
	`)
	if err != nil {
		return nil, err
	}

	return &storage{
		db: db,
	}, nil
}

func (s *storage) Close() {
	s.db.Close()
}

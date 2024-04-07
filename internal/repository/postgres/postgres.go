package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type repository struct {
	db *pgx.Conn
}

func New(dsn string) (*repository, error) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}
	err = conn.Ping(ctx)
	if err != nil {
		return nil, err
	}
	return &repository{
		db: conn,
	}, nil
}

func (r *repository) Close(ctx context.Context) error {
	return r.db.Close(ctx)
}

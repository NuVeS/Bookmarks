package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type API interface {
	AddUser(ctx context.Context, login string, password string) error
}

type Storage struct {
	DBPool *pgxpool.Pool
}

func (storage *Storage) AddUser(ctx context.Context, login string, password string) error {
	conn, err := storage.DBPool.Acquire(context.Background())
	if err != nil {
		return err
	}

	defer conn.Release()

	req := "INSERT INTO users(username, password) VALUES ($1, $2) RETURNING id"
	row := conn.QueryRow(ctx, req, login, password)

	var id string
	err = row.Scan(&id)
	return err
}

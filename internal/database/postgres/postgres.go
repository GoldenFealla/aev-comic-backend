package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	Conn *pgxpool.Pool
)

func New(uri string) error {
	var err error
	Conn, err = pgxpool.New(context.Background(), uri)
	if err != nil {
		return err
	}
	return nil
}

package repository

import (
	"context"

	"github.com/jackc/pgx/v5"

	"Goldenfealla/aev-comic/internal/database/postgres"
	"Goldenfealla/aev-comic/internal/domain"
)

type Comic struct{}

func NewComicRepository() *Comic {
	return &Comic{}
}

func (r *Comic) GetComicList(ctx context.Context) ([]*domain.Comic, error) {
	query := `
		SELECT * FROM comic
	`

	rows, _ := postgres.Conn.Query(ctx, query)

	comics, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[domain.Comic])

	if err != nil {
		return nil, err
	}

	return comics, err
}

func (r *Comic) GetComicImageList(ctx context.Context, code string) ([]*domain.ComicImage, error) {
	query := `
		SELECT * FROM image WHERE code=@code
	`

	args := &pgx.NamedArgs{
		"code": code,
	}

	rows, _ := postgres.Conn.Query(ctx, query, args)

	images, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[domain.ComicImage])

	if err != nil {
		return nil, err
	}

	return images, err
}

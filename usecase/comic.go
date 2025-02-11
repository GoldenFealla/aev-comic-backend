package usecase

import (
	"context"

	"Goldenfealla/aev-comic/internal/domain"
)

type ComicRepository interface {
	GetComicList(ctx context.Context) ([]*domain.Comic, error)
	GetComicImageList(ctx context.Context, code string) ([]*domain.ComicImage, error)
}

type Comic struct {
	r ComicRepository
}

func NewComicUsecase(r ComicRepository) *Comic {
	return &Comic{
		r,
	}
}

func (u *Comic) GetComicList(ctx context.Context) ([]*domain.Comic, error) {
	return u.r.GetComicList(ctx)
}

func (u *Comic) GetComicImageList(ctx context.Context, code string) ([]*domain.ComicImage, error) {
	return u.r.GetComicImageList(ctx, code)
}

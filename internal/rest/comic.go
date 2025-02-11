package rest

import (
	"context"
	"net/http"

	"Goldenfealla/aev-comic/internal/domain"

	"github.com/labstack/echo/v4"
)

type ComicUsecase interface {
	GetComicList(ctx context.Context) ([]*domain.Comic, error)
	GetComicImageList(ctx context.Context, code string) ([]*domain.ComicImage, error)
}

type handler struct {
	u ComicUsecase
}

func NewComicHandler(e *echo.Echo, u ComicUsecase) {
	group := e.Group("comic")

	h := handler{
		u,
	}

	group.GET("/list-comic", h.listComic)
	group.GET("/list-comic-image", h.listComicImage)
}

func (h *handler) listComic(c echo.Context) error {
	comics, err := h.u.GetComicList(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, comics)
}

func (h *handler) listComicImage(c echo.Context) error {
	if hasID := c.QueryParams().Has("code"); !hasID {
		return echo.NewHTTPError(http.StatusBadRequest, "query param 'code' is required")
	}

	code := c.QueryParams().Get("code")

	images, err := h.u.GetComicImageList(c.Request().Context(), code)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, images)
}

package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"Goldenfealla/template-go-echo/internal/domain"
	"Goldenfealla/template-go-echo/internal/validation"
)

type BookUsecase interface {
	GetBook(id string) (*domain.Book, error)
	GetBookList() ([]*domain.Book, error)
	CreateBook(book *domain.CreateBook) error
	UpdateBook(id string, book *domain.UpdateBook) error
}

type handler struct {
	u BookUsecase
}

func NewBookHandler(e *echo.Echo, u BookUsecase) {
	group := e.Group("book")

	h := handler{
		u,
	}

	group.GET("/get", h.get)
	group.GET("/list", h.list)
	group.POST("/create", h.create)
	group.PUT("/update", h.update)
}

func (h *handler) get(c echo.Context) error {
	if hasID := c.QueryParams().Has("id"); !hasID {
		return echo.NewHTTPError(http.StatusBadRequest, "query param 'id' is required")
	}

	id := c.QueryParams().Get("id")

	book, err := h.u.GetBook(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, book)
}

func (h *handler) list(c echo.Context) error {
	books, err := h.u.GetBookList()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, books)
}

func (h *handler) create(c echo.Context) error {
	var book domain.CreateBook

	err := c.Bind(&book)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	ves := validation.Validate(book)
	if ves != nil {
		return echo.NewHTTPError(http.StatusBadRequest, ves)
	}

	err = h.u.CreateBook(&book)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return nil
}

func (h *handler) update(c echo.Context) error {
	if hasID := c.QueryParams().Has("id"); !hasID {
		return echo.NewHTTPError(http.StatusBadRequest, "query param 'id' is required")
	}

	id := c.QueryParams().Get("id")

	var book domain.UpdateBook
	err := c.Bind(&book)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	ves := validation.Validate(book)
	if ves != nil {
		return echo.NewHTTPError(http.StatusBadRequest, ves)
	}

	err = h.u.UpdateBook(id, &book)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, book)
}

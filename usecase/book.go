package usecase

import "Goldenfealla/template-go-echo/internal/domain"

type BookRepository interface {
	GetBook(id string) (*domain.Book, error)
	GetBookList() ([]*domain.Book, error)
	CreateBook(book *domain.CreateBook) error
	UpdateBook(id string, book *domain.UpdateBook) error
}

type usecase struct{ r BookRepository }

func NewBookUsecase(r BookRepository) *usecase {
	return &usecase{r}
}

func (u usecase) GetBookList() ([]*domain.Book, error) {
	return u.r.GetBookList()
}

func (u usecase) GetBook(id string) (*domain.Book, error) {
	return u.r.GetBook(id)
}

func (u usecase) CreateBook(book *domain.CreateBook) error {
	return u.r.CreateBook(book)
}

func (u usecase) UpdateBook(id string, book *domain.UpdateBook) error {
	return u.r.UpdateBook(id, book)
}

package repository

import (
	"Goldenfealla/template-go-echo/internal/domain"
	"errors"

	"github.com/google/uuid"
)

var books = map[string]*domain.Book{}

type repository struct{}

func NewBookRepository() *repository {
	return &repository{}
}

func (r *repository) GetBook(id string) (*domain.Book, error) {
	lookID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	book, ok := books[lookID.String()]
	if !ok {
		return nil, errors.New("no book found")
	}

	return book, nil
}

func (r *repository) GetBookList() ([]*domain.Book, error) {
	var bookList []*domain.Book

	for _, book := range books {
		bookList = append(bookList, book)
	}

	return bookList, nil
}

func (r *repository) CreateBook(book *domain.CreateBook) error {
	newID, err := uuid.NewV7()
	if err != nil {
		return err
	}

	newBook := &domain.Book{
		ID:     newID,
		Name:   book.Name,
		Author: book.Author,
	}

	books[newID.String()] = newBook
	return nil
}

func (r *repository) UpdateBook(id string, newBook *domain.UpdateBook) error {
	lookID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	book, ok := books[lookID.String()]
	if !ok {
		return errors.New("no book found")
	}

	book.Name = newBook.Name
	book.Author = newBook.Author

	return nil
}

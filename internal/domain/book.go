package domain

import "github.com/google/uuid"

type Book struct {
	ID     uuid.UUID `json:"id" validate:"uuid"`
	Name   string    `json:"name" validate:"required,gte=8"`
	Author string    `json:"author" validate:"required,gte=8"`
}

type CreateBook struct {
	Name   string `json:"name" validate:"required,gte=8"`
	Author string `json:"author" validate:"required,gte=8"`
}

type UpdateBook struct {
	Name   string `json:"name" validate:"required,gte=8"`
	Author string `json:"author" validate:"required,gte=8"`
}

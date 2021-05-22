package service

import (
	"github.com/plutonio00/books-api/internal/model"
	"github.com/plutonio00/books-api/internal/repository"
)

type Books interface {
	FindById(id string) (*model.Book, error)
}

type Authors interface {
}

type Services struct {
	Books   Books
	Authors Authors
}

type Deps struct {
	Repos *repository.Repositories
}

func NewServices(deps Deps) *Services {
	return &Services{
		Books:   NewBooksService(deps.Repos.Books),
		Authors: NewAuthorsService(deps.Repos.Authors),
	}
}

package service

import (
	"github.com/plutonio00/books-api/internal/repository"
)

type Books interface {
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

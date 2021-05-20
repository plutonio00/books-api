package service

import (
	"github.com/plutonio00/books-api/internal/repository"
)

type Services struct {
	Books   BooksService
	Authors AuthorsService
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

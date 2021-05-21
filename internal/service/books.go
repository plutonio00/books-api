package service

import (
	"github.com/plutonio00/books-api/internal/repository"
)

type BooksService struct {
	repo repository.Books
}

func NewBooksService(repo repository.Books) *BooksService {
	return &BooksService{
		repo: repo,
	}
}

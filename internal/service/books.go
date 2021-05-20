package service

import (
	"github.com/plutonio00/books-api/internal/repository"
)

type BooksService struct {
	repo repository.BooksRepo
}

func NewBooksService(repo repository.BooksRepo) *BooksService {
	return &BooksService{
		repo: repo,
	}
}

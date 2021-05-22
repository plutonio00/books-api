package service

import (
	"github.com/plutonio00/books-api/internal/model"
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

func (s *BooksService) FindById(id string) (*model.Book, error) {
	return s.repo.FindById(id)
}

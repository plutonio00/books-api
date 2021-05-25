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

func (s *BooksService) GetBooksList() ([]model.Book, error) {
	return s.repo.GetBooksList()
}

func (s *BooksService) DeleteById(id string) error {
	return s.repo.DeleteById(id)
}

func (s *BooksService) UpdateById(keys []string, values []interface{}) error {
	return s.repo.UpdateById(keys, values)
}

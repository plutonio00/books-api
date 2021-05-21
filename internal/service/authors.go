package service

import (
	"github.com/plutonio00/books-api/internal/repository"
)

type AuthorsService struct {
	repo repository.Authors
}

func NewAuthorsService(repo repository.Authors) *AuthorsService {
	return &AuthorsService{
		repo: repo,
	}
}

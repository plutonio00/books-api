package service

import (
	"github.com/plutonio00/books-api/internal/repository"
)

type AuthorsService struct {
	repo repository.AuthorsRepo
}

func NewAuthorsService(repo repository.AuthorsRepo) *AuthorsService {
	return &AuthorsService{
		repo: repo,
	}
}

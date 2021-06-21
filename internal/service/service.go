package service

import (
	"github.com/plutonio00/books-api/internal/model"
	"github.com/plutonio00/books-api/internal/repository"
	"github.com/plutonio00/books-api/pkg/token"
)

type Token struct {
	AccessToken  string
}

type Users interface {
	SignUp(string, string) error
	SignIn(string, string) (*Token, error)
}

type Books interface {
	FindById(id string) (*model.Book, error)
	GetBooksList() ([]model.Book, error)
	DeleteById(string) error
	UpdateById([]string, []interface{}) (int64, error)
}

type Authors interface {
}

type Services struct {
	Books   Books
	Authors Authors
	Users   Users
}

type Deps struct {
	Repos        *repository.Repositories
	TokenManager *token.TokenManager
}

func NewServices(deps Deps) *Services {
	return &Services{
		Books:   NewBooksService(deps.Repos.Books),
		Authors: NewAuthorsService(deps.Repos.Authors),
		Users:   NewUsersService(deps.Repos.Users, deps.TokenManager),
	}
}

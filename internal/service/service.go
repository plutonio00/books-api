package service

import (
	"github.com/plutonio00/books-api/internal/model"
	"github.com/plutonio00/books-api/internal/repository"
	"github.com/plutonio00/books-api/pkg/token"
)

type Token struct {
	AccessToken string `example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJcdTAwMDEifQ.O3Lp6-VLUmGrcHUO11b2LqUBXq-tQeyRTWxwAe6aNLQ"`
}

type Users interface {
	SignUp(*model.User) error
	SignIn(*model.User) (*Token, error)
}

type Books interface {
	FindById(int) (*model.Book, error)
	GetBooksList() ([]model.Book, error)
	DeleteById(int) error
	Update(*model.Book) error
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

package repository

import (
	"database/sql"
	"github.com/plutonio00/books-api/internal/model"
	"github.com/plutonio00/books-api/internal/model/input"
	mysqlRepos "github.com/plutonio00/books-api/internal/repository/mysql"
)

type Users interface {
	Create(input.UserCredentials) error
	GetByEmail(string) (*model.User, error)
}

type Books interface {
	FindById(id string) (*model.Book, error)
	GetBooksList() ([]model.Book, error)
	DeleteById(string) error
	UpdateById([]string, []interface{}) error
}

type Authors interface {
}

type Repositories struct {
	Books   Books
	Authors Authors
	Users   Users
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Books:   mysqlRepos.NewBooksRepo(db),
		Authors: mysqlRepos.NewAuthorsRepo(db),
		Users:   mysqlRepos.NewUsersRepo(db),
	}
}

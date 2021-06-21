package repository

import (
	"database/sql"
	"github.com/plutonio00/books-api/internal/model"
	mysqlRepos "github.com/plutonio00/books-api/internal/repository/mysql"
)

type Users interface {
	Create(string, string) error
	//     SetToken()
	GetByEmail(string) (*model.User, error)
}

type Books interface {
	FindById(id string) (*model.Book, error)
	GetBooksList() ([]model.Book, error)
	DeleteById(string) error
	UpdateById([]string, []interface{}) (int64, error)
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

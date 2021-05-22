package repository

import (
	"database/sql"
	"github.com/plutonio00/books-api/internal/model"
	mysqlRepos "github.com/plutonio00/books-api/internal/repository/mysql"
)

type Books interface {
	FindById(id string) (*model.Book, error)
}

type Authors interface {
}

type Repositories struct {
	Books   Books
	Authors Authors
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Books:   mysqlRepos.NewBooksRepo(db),
		Authors: mysqlRepos.NewAuthorsRepo(db),
	}
}

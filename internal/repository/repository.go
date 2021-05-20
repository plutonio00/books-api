package repository

import (
	"database/sql"
	mysqlRepos "github.com/plutonio00/books-api/internal/repository/mysql"
)

type BooksRepo interface {
}

type AuthorsRepo interface {
}

type Repositories struct {
	Books   BooksRepo
	Authors AuthorsRepo
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Books:   mysqlRepos.NewBooksRepo(db),
		Authors: mysqlRepos.NewAuthorsRepo(db),
	}
}

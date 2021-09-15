package repository

import (
	"gorm.io/gorm"
	"github.com/plutonio00/books-api/internal/model"
	mysqlRepos "github.com/plutonio00/books-api/internal/repository/mysql"
)

type Users interface {
	Create(*model.User) error
	GetByEmail(string) (*model.User, error)
}

type Books interface {
	FindById(id string) (*model.Book, error)
	GetBooksList() ([]model.Book, error)
	DeleteById(string) error
	Update(*model.Book) error
}

type Authors interface {
}

type Repositories struct {
	Books   Books
	Authors Authors
	Users   Users
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Books:   mysqlRepos.NewBooksRepo(db),
		Authors: mysqlRepos.NewAuthorsRepo(db),
		Users:   mysqlRepos.NewUsersRepo(db),
	}
}

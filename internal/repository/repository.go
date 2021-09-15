package repository

import (
	"github.com/plutonio00/books-api/internal/model"
	"gorm.io/gorm"
)

type Users interface {
	Create(*model.User) error
	GetByEmail(string) (*model.User, error)
}

type Books interface {
	FindById(id string) (*model.Book, error)
	GetBooksList() ([]model.Book, error)
	DeleteById(int) error
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
		Books:   NewBooksRepo(db),
		Authors: NewAuthorsRepo(db),
		Users:   NewUsersRepo(db),
	}
}

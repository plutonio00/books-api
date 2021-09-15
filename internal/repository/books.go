package repository

import (
	"gorm.io/gorm"
	api_errors "github.com/plutonio00/books-api/internal/error"
	"github.com/plutonio00/books-api/internal/model"
)

type BooksRepo struct {
	db        *gorm.DB
	tableName string
}

func NewBooksRepo(db *gorm.DB) *BooksRepo {
	return &BooksRepo{
		db:        db,
	}
}

func (r *BooksRepo) FindById(id string) (*model.Book, error) {

}

func (r *BooksRepo) GetBooksList() ([]model.Book, error) {

}

func (r *BooksRepo) DeleteById(id string) error {

}

func (r *BooksRepo) Update(*model.Book) error {

}


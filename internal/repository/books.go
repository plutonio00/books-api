package repository

import (
	"github.com/plutonio00/books-api/internal/model"
	"gorm.io/gorm"
)

type BooksRepo struct {
	db        *gorm.DB
	tableName string
}

func NewBooksRepo(db *gorm.DB) *BooksRepo {
	return &BooksRepo{
		db: db,
	}
}

func (r *BooksRepo) FindById(id string) (*model.Book, error) {
	book := &model.Book{}
	result := r.db.First(book)

	if result.Error != nil {
		return nil, result.Error
	}

	return book, nil
}

func (r *BooksRepo) GetBooksList() ([]model.Book, error) {
	books := []model.Book{}
	result := r.db.Find(&books)

	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}

func (r *BooksRepo) DeleteById(id int) error {
	result := r.db.Delete(&model.Book{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *BooksRepo) Update(book *model.Book) error {
	result := r.db.Model(book).Updates(book)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

package mysql

import (
	"database/sql"
)

type BooksRepo struct {
	db *sql.DB
}

func NewBooksRepo(db *sql.DB) *BooksRepo {
	return &BooksRepo{
		db: db,
	}
}

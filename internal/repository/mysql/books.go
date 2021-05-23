package mysql

import (
	"database/sql"
	"fmt"
	"github.com/plutonio00/books-api/internal/model"
)

type BooksRepo struct {
	db        *sql.DB
	tableName string
}

func NewBooksRepo(db *sql.DB) *BooksRepo {
	return &BooksRepo{
		db:        db,
		tableName: "book",
	}
}

func (r *BooksRepo) FindById(id string) (*model.Book, error) {
	qb := newQueryBuilder()

	selectedParams := []string{
		"b.id",
		"title",
		"release_year",
		"author_id",
		"first_name",
		"last_name",
		"is_male",
		"birth_date",
	}

	qb.Select(selectedParams, r.tableName+" b").InnerJoin("author a", "b.author_id = a.id").Where("b.id")


    book := &model.Book{}

	err := r.db.QueryRow(qb.Query, id).Scan(
               &book.Id,
               &book.Title,
               &book.ReleaseYear,
               &book.Author.Id,
               &book.Author.FirstName,
               &book.Author.LastName,
               &book.Author.IsMale,
               &book.Author.BirthDate,
           )

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return book, nil
}

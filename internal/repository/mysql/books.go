package mysql

import (
	"database/sql"
	"github.com/plutonio00/books-api/internal/model"
	// 	"fmt"
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

	qb := r.findAllWithAuthors().Where("b.id")
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
		// 	    if err == sql.ErrNoRows {
		// 	        return nil, repository.ErrBookNotFound
		// 	    }
		return nil, err
	}

	return book, nil
}

func (r *BooksRepo) GetBooksList() ([]model.Book, error) {
	qb := r.findAllWithAuthors()
	books := []model.Book{}
	rows, err := r.db.Query(qb.Query)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		book := model.Book{}

		err := rows.Scan(
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
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

func (r *BooksRepo) DeleteById(id string) error {
	qb := newQueryBuilder()
	qb.Delete(r.tableName).Where("id")

	_, err := r.db.Exec(qb.Query, id)

	if err != nil {
		return err
	}

	return nil
}

func (r *BooksRepo) UpdateById(keys []string, values []interface{}) (int64, error) {
	qb := newQueryBuilder()
	qb.Update(r.tableName, keys).Where("id")

	result, err := r.db.Exec(qb.Query, values...)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (r *BooksRepo) findAllWithAuthors() *QueryBuilder {
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

	return qb.Select(selectedParams, r.tableName+" b").InnerJoin("author a", "b.author_id = a.id")
}

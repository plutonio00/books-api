package mysql

import (
	"database/sql"
	"github.com/plutonio00/books-api/internal/model"
)

type UsersRepo struct {
	db        *sql.DB
	tableName string
}

func NewUsersRepo(db *sql.DB) *UsersRepo {
	return &UsersRepo{
		db:        db,
		tableName: "user",
	}
}

func (r *UsersRepo) Create(email string, passwordHash string) error {
	qb := newQueryBuilder()
	qb.Insert(r.tableName, []string{"email", "password"})

	_, err := r.db.Exec(qb.Query, email, passwordHash)

	if err != nil {
		return err
	}

	return nil
}

func (r *UsersRepo) GetByEmail(email string) (*model.User, error) {
	qb := newQueryBuilder()
	user := &model.User{}

	selectedParams := []string{
		"id",
		"email",
		"password",
	}

	qb.Select(selectedParams, r.tableName).Where("email")

	err := r.db.QueryRow(qb.Query, email).Scan(
		&user.Id,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

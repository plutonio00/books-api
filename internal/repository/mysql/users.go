package mysql

import (
	"database/sql"
	mysql_driver "github.com/go-sql-driver/mysql"
	api_errors "github.com/plutonio00/books-api/internal/error"
	"github.com/plutonio00/books-api/internal/model"
	"github.com/plutonio00/books-api/internal/model/input"
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

func (r *UsersRepo) Create(credentials input.UserCredentials) error {
	qb := newQueryBuilder()
	qb.Insert(r.tableName, []string{"email", "password"})

	_, err := r.db.Exec(qb.Query, credentials.Email, credentials.Password)

	if err != nil {
		if driverErr, ok := err.(*mysql_driver.MySQLError); ok {
			if driverErr.Number == duplicateKeyErrorNumber {
				err = api_errors.ErrUserHasAlreadyRegistered
			}
		}
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

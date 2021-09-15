package repository

import (
	"gorm.io/gorm"
	mysql_driver "github.com/go-sql-driver/mysql"
	api_errors "github.com/plutonio00/books-api/internal/error"
	"github.com/plutonio00/books-api/internal/model"
	"github.com/plutonio00/books-api/internal/model/input"
)

type UsersRepo struct {
	db        *gorm.DB
	tableName string
}

func NewUsersRepo(db *gorm.DB) *UsersRepo {
	return &UsersRepo{
		db:        db,
	}
}

func (r *UsersRepo) Create(user *model.User) error {

}

func (r *UsersRepo) GetByEmail(email string) (*model.User, error) {

}

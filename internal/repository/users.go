package repository

import (
	"gorm.io/gorm"
	"github.com/plutonio00/books-api/internal/model"
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
    result := r.db.Create(user)

    if result.Error != nil {
        return result.Error
    }

    return nil
}

func (r *UsersRepo) GetByEmail(email string) (*model.User, error) {
    user := &model.User{}
    result := r.db.Where("email = ?", email).Find(user)

    if result.Error != nil {
        return nil, result.Error
    }

    return user, nil
}

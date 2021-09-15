package repository

import (
	"gorm.io/gorm"
)

type AuthorsRepo struct {
	db *gorm.DB
}

func NewAuthorsRepo(db *gorm.DB) *AuthorsRepo {
	return &AuthorsRepo{
		db: db,
	}
}

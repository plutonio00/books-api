package mysql

import (
    "database/sql"
)

type AuthorsRepo struct {
    db *sql.DB
}

func NewAuthorsRepo(db *sql.DB) *AuthorsRepo {
    return &AuthorsRepo {
        db: db,
    }
}
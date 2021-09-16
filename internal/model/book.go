package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"time"
)

type Book struct {
	Id          uint   `json:"id" example:"1"`
	Title       string `json:"title" example:"Games of thrones"`
	ReleaseYear int    `json:"release_year" example:"1997"`
	Author      Author `json:"author"`
	AuthorId    int    `json:"-"`
}

func (b *Book) Validate() error {
    year := time.Now().Year()

	return validation.ValidateStruct(
		b,
		validation.Field(&b.Id, validation.Required),
		validation.Field(&b.ReleaseYear, validation.Max(year)),
	)
}
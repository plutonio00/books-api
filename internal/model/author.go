package model

import (
	"time"
)

type Author struct {
	Id        uint      `json:"id" example:"10"`
	FirstName string    `json:"first_name" example:"George"`
	LastName  string    `json:"last_name" example:"Martin"`
	IsMale    bool      `json:"is_male" example:"true"`
	BirthDate time.Time `json:"birth_date" example:"1948-09-20"`
	Books     []Book    `json:"books,omitempty"`
}

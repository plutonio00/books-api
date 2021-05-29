package model

import (
	"time"
)

type Author struct {
	Id        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	IsMale    bool      `json:"is_male"`
	BirthDate time.Date `json:"birth_date"`
}

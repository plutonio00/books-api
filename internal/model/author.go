package model

import (
	//"time"
)

type Author struct {
	Id        int           `json:"id"`
	FirstName string        `json:"first_name"`
	LastName  string        `json:"last_name"`
	Gender    bool          `json:"gender"`
	BirthDate []uint8 `json:"birth_date"`
}

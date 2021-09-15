package model

import (
	"time"
)

type Author struct {
	Id        uint      `gorm:"primary_key" json:"id" example:"10"`
	FirstName string    `gorm:"type:varchar(50) not null" json:"first_name" example:"George"`
	LastName  string    `gorm:"type:varchar(50) not null" json:"last_name" example:"Martin"`
	IsMale    bool      `gorm:"tinyint(1) not null" json:"is_male" example:"true"`
	BirthDate time.Time `gorm:"DATE not null" json:"birth_date" example:"1948-09-20"`
	Books     []Book    `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
}

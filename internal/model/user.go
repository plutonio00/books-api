package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type User struct {
	Id       uint   `gorm:"primary_key" json:"id"`
	Email    string `gorm:"index:user_email_uindex,unique" json:"email"`
	Password string `gorm:"not null" json:"password,omitempty"`
}

func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Required, validation.Length(6, 100)),
	)
}
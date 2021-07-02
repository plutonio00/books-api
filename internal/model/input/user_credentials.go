package input

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type UserCredentials struct {
	Email    string
	Password string
}

func (uc *UserCredentials) Validate() error {
	return validation.ValidateStruct(
		uc,
		validation.Field(&uc.Email, validation.Required, is.Email),
		validation.Field(&uc.Password, validation.Required, validation.Length(6, 100)),
	)
}

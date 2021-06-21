package error

import (
	"errors"
)

var (
	ErrBookNotFound             = errors.New("Book doesn't exist")
	ErrInvalidCredentials       = errors.New("Invalid credentials: email and password are mandatory")
	ErrUserHasAlreadyRegistered = errors.New("Such user has already registered")
)

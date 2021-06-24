package error

import (
	"errors"
)

var (
	ErrBookNotFound             = errors.New("Book doesn't exist")
	ErrEmptyCredentials         = errors.New("Invalid credentials: email and password are mandatory")
	ErrInvalidCredentials       = errors.New("Invalid credentials: incorrect email or password")
	ErrUserHasAlreadyRegistered = errors.New("Such user has already registered")
	ErrUnauthorized             = errors.New("Unauthorized error")
	ErrInvalidAuthorId          = errors.New("Invalid author id")
)

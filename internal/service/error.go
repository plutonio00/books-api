package service

import (
    "errors"
)

var (
    ErrBookNotFound = errors.New("book doesn't exist")
)
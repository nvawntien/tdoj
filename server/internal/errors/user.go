package errors

import "errors"

var (
	UserConflict = errors.New("User already exists")
)

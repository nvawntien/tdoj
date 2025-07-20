package errors

import "errors"

var (
	UserConflict      = errors.New("User already exists")
	UserNotFound      = errors.New("User not found")
	PasswordIncorrect = errors.New("Password is incorrect")
)

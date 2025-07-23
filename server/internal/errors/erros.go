package errors

import "errors"

var (
	UserConflict      = errors.New("User already exists")
	UserNotFound      = errors.New("User not found")
	PasswordIncorrect = errors.New("Password is incorrect")
)

var (
	ErrGenerateToken = errors.New("Generate token failed")
)

var (
	ErrCheckExistsUserByEmail = errors.New("Check exists user by email failed")
)

package models

import "time"

type User struct {
	UserID    string    `json:"user_id" db:"user_id""`
	FullName  string    `json:"full_name" db:"full_name" binding:"required"`
	Email     string    `json:"email" db:"email" binding:"required,email"`
	Username  string    `json:"username" db:"username" binding:"required"`
	Password  string    `json:"password" db:"password" binding:"required,min=8"`
	Rating    string    `json:"rating" db:"rating"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

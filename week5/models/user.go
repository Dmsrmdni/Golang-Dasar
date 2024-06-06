package models

import "time"

type User struct {
	Id        int        `json:"id"`
	Name      string     `json:"name" validate:"required"`
	Email     string     `json:"email" validate:"required,email"`
	Password  string     `json:"password"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

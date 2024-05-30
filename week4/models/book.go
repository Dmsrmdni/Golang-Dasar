package models

import "time"

type Book struct {
	Id              int        `json:"id"`
	CategoryId      int        `json:"category_id"`
	Title           string     `json:"title"`
	UserId          int        `json:"user_id"`
	PublicationYear int        `json:"publication_year"`
	Description     string     `json:"description"`
	CreatedAt       *time.Time `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
}

type GetBook struct {
	Id              int        `json:"id"`
	Category        string     `json:"category"`
	Title           string     `json:"title"`
	Author          string     `json:"author"`
	PublicationYear int        `json:"publication_year"`
	Description     string     `json:"description"`
	CreatedAt       *time.Time `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
}

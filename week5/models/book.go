package models

import "time"

type Book struct {
	Id              int        `json:"id" form:"id"`
	CategoryId      int        `json:"category_id" form:"category_id"`
	Title           string     `json:"title" form:"title"`
	UserId          int        `json:"user_id" form:"user_id"`
	PublicationYear int        `json:"publication_year" form:"publication_year"`
	Description     string     `json:"description" form:"description"`
	BookImage       string     `json:"book_image" form:"book_image"`
	CreatedAt       *time.Time `json:"created_at" form:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at" form:"updated_at"`
}

type GetBook struct {
	Id              int        `json:"id"`
	Category        string     `json:"category"`
	Title           string     `json:"title"`
	Author          string     `json:"author"`
	PublicationYear int        `json:"publication_year"`
	Description     string     `json:"description"`
	BookImage       string     `json:"book_image"`
	CreatedAt       *time.Time `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
}

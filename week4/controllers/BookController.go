package controller

import (
	"database/sql"
	"net/http"
	"week4/config"
	"week4/models"

	"github.com/labstack/echo/v4"
)

// Get Book
func GetBook(ctx echo.Context) error {
	db, err := config.ConnectDB()

	if err != nil {
		return err
	}

	defer db.Close()

	title := "%" + ctx.QueryParam("title") + "%"
	author := "%" + ctx.QueryParam("author") + "%"
	limit := ctx.QueryParam("limit")

	if limit == "" {
		limit = "10"
	}
	offset := ctx.QueryParam("offset")

	if offset == "" {
		offset = "0"
	}

	query := `
		SELECT 
			books.id,
			categories.name AS category,
			books.title,
			users.name AS author,
			books.publication_year,
			books.description,
			books.created_at,
			books.updated_at 
		FROM 
			books 
			JOIN categories ON books.category_id = categories.id
			JOIN users ON books.user_id = users.id
		WHERE
			books.title LIKE $1
			AND users.name LIKE $2
		ORDER BY books.id ASC
		LIMIT $3 OFFSET $4
	`

	rows, err := db.Query(query, title, author, limit, offset)

	if err != nil {
		return err
	}

	var data_books []models.GetBook
	for rows.Next() {
		var book models.GetBook
		var updatedAt sql.NullTime

		err = rows.Scan(&book.Id, &book.Category, &book.Title, &book.Author, &book.PublicationYear, &book.Description, &book.CreatedAt, &updatedAt)

		if err != nil {
			return err
		}

		if updatedAt.Valid {
			book.UpdatedAt = &updatedAt.Time
		}

		data_books = append(data_books, book)
	}

	response := models.Response{
		Status:  http.StatusOK,
		Message: "Get Books successfully",
		Data:    data_books,
	}

	return ctx.JSON(http.StatusOK, response)
}

// Create Book
func CreateBook(ctx echo.Context) error {
	db, err := config.ConnectDB()

	if err != nil {
		return err
	}

	defer db.Close()

	var book models.Book

	if err := ctx.Bind(&book); err != nil {
		return err
	}

	query := "INSERT INTO books(category_id,title,user_id,publication_year,description) VALUES ($1,$2,$3,$4,$5) RETURNING id,category_id,title,user_id,publication_year,description,created_at,updated_at"

	err = db.QueryRow(query, &book.CategoryId, &book.Title, &book.UserId, &book.PublicationYear, &book.Description).Scan(&book.Id, &book.CategoryId, &book.Title, &book.UserId, &book.PublicationYear, &book.Description, &book.CreatedAt, &book.UpdatedAt)

	if err != nil {
		return err
	}

	response := models.Response{
		Status:  http.StatusCreated,
		Message: "Book created successfully",
		Data:    book,
	}

	return ctx.JSON(http.StatusCreated, response)
}

// Detail Book
func DetailBook(ctx echo.Context) error {
	db, err := config.ConnectDB()

	if err != nil {
		return err
	}

	defer db.Close()

	bookId := ctx.Param("id")

	var book models.GetBook

	query := `
		SELECT 
			books.id,
			categories.name AS category,
			books.title,
			users.name AS author,
			books.publication_year,
			books.description,
			books.created_at,
			books.updated_at 
		FROM 
			books 
			JOIN categories ON books.category_id = categories.id
			JOIN users ON books.user_id = users.id
		WHERE books.id = $1
	`

	err = db.QueryRow(query, &bookId).Scan(&book.Id, &book.Category, &book.Title, &book.Author, &book.PublicationYear, &book.Description, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		response := models.Response{
			Status:  http.StatusNotFound,
			Message: "Book not found",
		}

		return ctx.JSON(http.StatusNotFound, response)
	}

	response := models.Response{
		Status:  http.StatusOK,
		Message: "Get Book detail successfully",
		Data:    book,
	}

	return ctx.JSON(http.StatusOK, response)
}

// Update Book
func UpdateBook(ctx echo.Context) error {
	db, err := config.ConnectDB()

	if err != nil {
		return err
	}

	defer db.Close()

	bookId := ctx.Param("id")

	var book models.Book

	if err := ctx.Bind(&book); err != nil {
		return err
	}

	query := "UPDATE books SET category_id = $1 , title = $2 , user_id = $3 , publication_year = $4 , description = $5 , updated_at = CURRENT_TIMESTAMP WHERE id = $6 RETURNING id,category_id,title,user_id,publication_year,description,created_at,updated_at"

	err = db.QueryRow(query, &book.CategoryId, &book.Title, &book.UserId, &book.PublicationYear, &book.Description, bookId).Scan(&book.Id, &book.CategoryId, &book.Title, &book.UserId, &book.PublicationYear, &book.Description, &book.CreatedAt, &book.UpdatedAt)

	if err != nil {
		return err
	}

	response := models.Response{
		Status:  http.StatusOK,
		Message: "Book Updated successfully",
		Data:    book,
	}

	return ctx.JSON(http.StatusOK, response)
}

// Delete Book
func DeleteBook(ctx echo.Context) error {
	db, err := config.ConnectDB()

	if err != nil {
		return err
	}

	defer db.Close()

	bookId := ctx.Param("id")

	query := "DELETE FROM books WHERE id = $1 RETURNING id"

	result, err := db.Exec(query, bookId)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		response := models.Response{
			Status:  http.StatusNotFound,
			Message: "Book not found",
		}

		return ctx.JSON(http.StatusNotFound, response)
	}

	response := models.Response{
		Status:  http.StatusOK,
		Message: "Book Deleted successfully",
	}

	return ctx.JSON(http.StatusOK, response)
}

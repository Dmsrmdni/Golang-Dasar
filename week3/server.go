package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

// type Siswa struct {
// 	Name   string `json:"name" form:"name"`
// 	Alamat string `json:"alamat" form:"alamat"`
// }

// func addSiswa(ctx echo.Context) error {
// 	var siswa Siswa

// 	if err := ctx.Bind(&siswa); err != nil {
// 		return err
// 	}

// 	return ctx.JSON(http.StatusOK, siswa)
// }

type Books struct {
	Id              int    `json:"id"`
	Title           string `json:"title"`
	Author          string `json:"author"`
	PublicationYear int    `json:"publication_year"`
	Description     string `json:"description"`
	CreatedAt       time.Time
	UpdatedAt       *time.Time
}

func main() {
	// Instansi objek Echo
	e := echo.New()

	//Routing
	e.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello, world!")
	})

	// e.GET("/string", func(ctx echo.Context) error {
	// 	data := "Hello world"
	// 	return ctx.String(http.StatusOK, data)
	// })

	// e.GET("/html", func(ctx echo.Context) error {
	// 	data := "<h1>Hello, World!</h1>"
	// 	return ctx.HTML(http.StatusOK, data)
	// })

	// e.GET("/redirect", func(ctx echo.Context) error {
	// 	return ctx.Redirect(http.StatusTemporaryRedirect, "/")
	// })

	// e.GET("/json", func(ctx echo.Context) error {
	// 	data := map[string]any{"Message": "Hello", "Counter": 2}
	// 	return ctx.JSON(http.StatusOK, data)
	// })

	// e.POST("/siswa", addSiswa)
	// // /users/1
	// e.GET("/users/:id", func(ctx echo.Context) error {
	// 	id := ctx.Param("id")
	// 	return ctx.String(http.StatusOK, id)
	// })

	// // /movie?genre=action
	// e.GET("/movie", func(c echo.Context) error {
	// 	genre := c.QueryParam("genre")
	// 	return c.String(http.StatusOK, "Genre:"+genre)
	// })

	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUsername, dbPassword, dbName, dbPort))

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Get Data Buku
	e.GET("/books", func(ctx echo.Context) error {
		query := "SELECT id,title,author,publication_year,description,created_at,updated_at FROM books"
		rows, err := db.Query(query)
		if err != nil {
			fmt.Println(err)
			return err
		}

		var data_books []Books

		for rows.Next() {
			var books Books
			var updatedAt sql.NullTime
			err := rows.Scan(&books.Id, &books.Title, &books.Author, &books.PublicationYear, &books.Description, &books.CreatedAt, &updatedAt)

			if err != nil {
				fmt.Println(err)
				return err
			}

			if updatedAt.Valid {
				books.UpdatedAt = &updatedAt.Time
			}

			data_books = append(data_books, books)
		}

		return ctx.JSON(http.StatusOK, data_books)
	})

	// Detail Buku
	e.GET("/books/:id", func(ctx echo.Context) error {
		bookID := ctx.Param("id")

		query := "SELECT id, title, author, publication_year, description, created_at, updated_at FROM books WHERE id = $1"

		row := db.QueryRow(query, bookID)

		var book Books

		err := row.Scan(&book.Id, &book.Title, &book.Author, &book.PublicationYear, &book.Description, &book.CreatedAt, &book.UpdatedAt)
		if err != nil {
			fmt.Println(err)
			return ctx.String(http.StatusNotFound, "Buku Tidak Ada")
		}

		return ctx.JSON(http.StatusOK, book)
	})

	// Delete Buku
	e.DELETE("/books/:id", func(ctx echo.Context) error {
		bookID := ctx.Param("id")

		query := "DELETE FROM books WHERE id = $1"

		_, err := db.Exec(query, bookID)

		if err != nil {
			fmt.Println(err)
			return err
		}

		return ctx.String(http.StatusOK, "Book deleted successfully")
	})

	// Tambah Buku
	e.POST("/books", func(ctx echo.Context) error {

		var book Books
		if err := ctx.Bind(&book); err != nil {
			return err
		}

		query := "INSERT INTO books (title, author, publication_year, description) VALUES ($1, $2, $3, $4)"

		_, err := db.Exec(query, &book.Title, &book.Author, &book.PublicationYear, &book.Description)

		if err != nil {
			fmt.Println(err)
			return err
		}

		response := map[string]any{
			"status":  200,
			"message": "Created book successfully",
		}

		return ctx.JSON(http.StatusOK, response)
	})

	// Update Buku
	e.PUT("/books/:id", func(ctx echo.Context) error {
		bookID := ctx.Param("id")
		var userId int

		var book Books

		if err := ctx.Bind(&book); err != nil {
			return err
		}

		query := "UPDATE books SET title = $1, author = $2, publication_year = $3, description = $4,updated_at = CURRENT_TIMESTAMP WHERE id = $5 RETURNING id"

		err := db.QueryRow(query, &book.Title, &book.Author, &book.PublicationYear, &book.Description, bookID).Scan(&userId)

		if err != nil {
			fmt.Println(err)
			return err
		}

		response := map[string]any{
			"status":  200,
			"message": "Updated book successfully",
			"data":    userId,
		}

		return ctx.JSON(http.StatusOK, response)
	})

	// Mulai Server
	e.Logger.Fatal(e.Start(":8000"))
}

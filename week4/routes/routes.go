package routes

import (
	"net/http"
	controller "week4/controllers"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	// Instansi objek Echo
	e := echo.New()

	//Routing
	e.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello, world!")
	})

	// Users
	e.GET("/users", controller.GetUser)
	e.POST("/users", controller.CreateUser)
	e.GET("/users/:id", controller.DetailUser)
	e.PUT("/users/:id", controller.UpdateUser)
	e.DELETE("/users/:id", controller.DeleteUser)
	//endUser

	// Categories
	e.GET("/categories", controller.GetCategory)
	e.POST("/categories", controller.CreateCategory)
	e.GET("/categories/:id", controller.DetailCategory)
	e.PUT("/categories/:id", controller.UpdateCategory)
	e.DELETE("/categories/:id", controller.DeleteCategory)
	// endCategories

	// Books
	e.GET("/books", controller.GetBook)
	e.POST("/books", controller.CreateBook)
	e.GET("/books/:id", controller.DetailBook)
	e.PUT("/books/:id", controller.UpdateBook)
	e.DELETE("/books/:id", controller.DeleteBook)
	// endBooks

	return e
}

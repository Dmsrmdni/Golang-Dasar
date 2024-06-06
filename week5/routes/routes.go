package routes

import (
	"week5/controllers"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	// Instansi objek Echo
	e := echo.New()

	e.Static("media", "assets/images")

	// Middleware
	// digunakan untuk mencatat log HTTP reques
	// e.Use(middleware.Logger())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// Routes
	e.POST("/login", controllers.Login)

	// Middleware JWT
	authMiddleware := (echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte("secret"),
	}))

	// Group for authenticated routes
	authenticated := e.Group("", authMiddleware)

	// Users
	authenticated.GET("/users", controllers.GetUser)
	authenticated.POST("/users", controllers.CreateUser)
	authenticated.GET("/users/:id", controllers.DetailUser)
	authenticated.PUT("/users/:id", controllers.UpdateUser)
	authenticated.DELETE("/users/:id", controllers.DeleteUser)

	// Categories
	authenticated.GET("/categories", controllers.GetCategory)
	authenticated.POST("/categories", controllers.CreateCategory)
	authenticated.GET("/categories/:id", controllers.DetailCategory)
	authenticated.PUT("/categories/:id", controllers.UpdateCategory)
	authenticated.DELETE("/categories/:id", controllers.DeleteCategory)

	// Books
	authenticated.GET("/books", controllers.GetBook)
	authenticated.POST("/books", controllers.CreateBook)
	authenticated.GET("/books/:id", controllers.DetailBook)
	authenticated.PUT("/books/:id", controllers.UpdateBook)
	authenticated.DELETE("/books/:id", controllers.DeleteBook)
	// endBooks
	return e

}

package controllers

import (
	"database/sql"
	"net/http"
	"week5/config"
	"week5/models"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// Get Category
func GetCategory(ctx echo.Context) error {
	db, err := config.ConnectDB()

	if err != nil {
		return err
	}

	defer db.Close()

	query := "SELECT id,name,created_at,updated_at FROM categories ORDER BY id ASC"

	rows, err := db.Query(query)

	var data_categories []models.Category
	if err != nil {
		return err
	}

	for rows.Next() {
		var category models.Category
		var updatedAt sql.NullTime

		err = rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &updatedAt)

		if err != nil {
			return err
		}

		if updatedAt.Valid {
			category.UpdatedAt = &updatedAt.Time
		}

		data_categories = append(data_categories, category)
	}

	response := models.Response{
		Status:  http.StatusOK,
		Message: "Get Categories successfully",
		Data:    data_categories,
	}

	return ctx.JSON(http.StatusOK, response)
}

// Create Category
func CreateCategory(ctx echo.Context) error {
	db, err := config.ConnectDB()

	if err != nil {
		return err
	}

	defer db.Close()

	var category models.Category

	if err := ctx.Bind(&category); err != nil {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(&category); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			Status:  http.StatusBadRequest,
			Message: "Validation error",
			Data:    err.Error(),
		})
	}

	query := "INSERT INTO categories(name) VALUES ($1) RETURNING id,name,created_at,updated_at"

	err = db.QueryRow(query, &category.Name).Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt)

	if err != nil {
		return err
	}

	response := models.Response{
		Status:  http.StatusCreated,
		Message: "Category created successfully",
		Data:    category,
	}

	return ctx.JSON(http.StatusCreated, response)
}

// Detail Category
func DetailCategory(ctx echo.Context) error {
	db, err := config.ConnectDB()

	if err != nil {
		return err
	}

	defer db.Close()

	categoryId := ctx.Param("id")

	var category models.Category

	query := "SELECT id,name,created_at,updated_at FROM categories WHERE id = $1"

	err = db.QueryRow(query, &categoryId).Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt)

	if err != nil {
		response := models.Response{
			Status:  http.StatusNotFound,
			Message: "Category not found",
		}

		return ctx.JSON(http.StatusNotFound, response)
	}

	response := models.Response{
		Status:  http.StatusOK,
		Message: "Get category detail successfully",
		Data:    category,
	}

	return ctx.JSON(http.StatusOK, response)
}

// Update Category
func UpdateCategory(ctx echo.Context) error {
	db, err := config.ConnectDB()

	if err != nil {
		return err
	}

	defer db.Close()

	categoryId := ctx.Param("id")

	var category models.Category

	if err := ctx.Bind(&category); err != nil {
		return err
	}

	query := "UPDATE categories SET name = $1 , updated_at = CURRENT_TIMESTAMP WHERE id = $2 RETURNING id,name,created_at,updated_at"

	err = db.QueryRow(query, &category.Name, categoryId).Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt)

	if err != nil {
		return err
	}

	response := models.Response{
		Status:  http.StatusOK,
		Message: "Category Updated successfully",
		Data:    category,
	}

	return ctx.JSON(http.StatusOK, response)
}

// Delete Category
func DeleteCategory(ctx echo.Context) error {
	db, err := config.ConnectDB()

	if err != nil {
		return err
	}

	defer db.Close()

	categoryId := ctx.Param("id")

	query := "DELETE FROM categories WHERE id = $1 RETURNING id"

	result, err := db.Exec(query, categoryId)

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
			Message: "Category not found",
		}

		return ctx.JSON(http.StatusNotFound, response)
	}

	response := models.Response{
		Status:  http.StatusOK,
		Message: "Category Deleted successfully",
	}

	return ctx.JSON(http.StatusOK, response)
}

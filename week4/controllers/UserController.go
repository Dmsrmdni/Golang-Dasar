package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"week4/config"
	"week4/models"

	"github.com/labstack/echo/v4"
)

// Get User
func GetUser(ctx echo.Context) error {

	db, err := config.ConnectDB()

	if err != nil {
		return err
	}

	defer db.Close()

	query := "SELECT id,name,email,created_at,updated_at FROM users ORDER BY id ASC"
	rows, err := db.Query(query)

	if err != nil {
		return (err)
	}

	var data_users []models.User

	for rows.Next() {
		var user models.User
		var updatedAt sql.NullTime
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt, &updatedAt)

		if err != nil {
			fmt.Println(err)
			return err
		}

		if updatedAt.Valid {
			user.UpdatedAt = &updatedAt.Time
		}

		data_users = append(data_users, user)
	}

	response := models.Response{
		Status:  http.StatusOK,
		Message: "Get user successfully",
		Data:    data_users,
	}

	return ctx.JSON(http.StatusOK, response)
}

// Create User
func CreateUser(ctx echo.Context) error {

	db, err := config.ConnectDB()

	if err != nil {
		return err
	}

	defer db.Close()

	var user models.User

	if err := ctx.Bind(&user); err != nil {
		return err
	}

	query := "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, name, email, created_at, updated_at"
	err = db.QueryRow(query, user.Name, user.Email).Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return err
	}

	response := models.Response{
		Status:  http.StatusCreated,
		Message: "User created successfully",
		Data:    user,
	}

	return ctx.JSON(http.StatusCreated, response)
}

// Detail User
func DetailUser(ctx echo.Context) error {
	db, err := config.ConnectDB()

	if err != nil {
		return err
	}

	defer db.Close()

	var user models.User

	userId := ctx.Param("id")

	query := "SELECT id,name,email,created_at,updated_at FROM users WHERE id = $1"

	err = db.QueryRow(query, userId).Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		response := models.Response{
			Status:  http.StatusNotFound,
			Message: "User not found",
		}

		return ctx.JSON(http.StatusNotFound, response)
	}

	response := models.Response{
		Status:  http.StatusOK,
		Message: "Get user detail successfully",
		Data:    user,
	}

	return ctx.JSON(http.StatusOK, response)
}

// Update User
func UpdateUser(ctx echo.Context) error {

	db, err := config.ConnectDB()

	if err != nil {
		return err
	}

	defer db.Close()

	userId := ctx.Param("id")

	var user models.User

	if err = ctx.Bind(&user); err != nil {
		return err
	}

	query := "UPDATE users SET name = $1, email = $2, updated_at = CURRENT_TIMESTAMP WHERE id = $3 RETURNING id, name, email, created_at, updated_at"

	err = db.QueryRow(query, &user.Name, &user.Email, userId).Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return err
	}

	response := models.Response{
		Status:  http.StatusOK,
		Message: "User Updated successfully",
		Data:    user,
	}

	return ctx.JSON(http.StatusOK, response)
}

// Delete User
func DeleteUser(ctx echo.Context) error {

	db, err := config.ConnectDB()

	if err != nil {
		return err
	}

	defer db.Close()

	userId := ctx.Param("id")

	query := "DELETE FROM users WHERE id = $1 RETURNING id"

	result, err := db.Exec(query, userId)
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
			Message: "User not found",
		}

		return ctx.JSON(http.StatusNotFound, response)
	}

	response := models.Response{
		Status:  http.StatusOK,
		Message: "User Deleted successfully",
	}

	return ctx.JSON(http.StatusOK, response)
}

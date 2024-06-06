package controllers

import (
	"crypto/sha1"
	"encoding/hex"
	"net/http"
	"time"
	"week5/config"
	"week5/models"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var signingKey = []byte("secret")

func hashPassword(password string) string {
	// Buat objek hasher SHA-1
	hasher := sha1.New()

	// Tulis string password ke hasher
	hasher.Write([]byte(password))

	// Dapatkan hasil hash dalam bentuk byte
	hashedBytes := hasher.Sum(nil)

	// Konversi byte menjadi string heksadesimal
	hashedPassword := hex.EncodeToString(hashedBytes)

	// Kembalikan string hash
	return hashedPassword
}

func Login(c echo.Context) error {

	db, err := config.ConnectDB()

	if err != nil {
		return err
	}

	defer db.Close()

	email := c.FormValue("email")
	password := c.FormValue("password")

	hashedPassword := hashPassword(password)
	query := "SELECT password FROM users WHERE email = $1"

	var passwordString string

	err = db.QueryRow(query, email).Scan(&passwordString)

	if err != nil {
		response := models.Response{
			Status:  http.StatusNotFound,
			Message: "User not found",
		}

		return c.JSON(http.StatusNotFound, response)
	}

	if hashedPassword == passwordString {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["email"] = email
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

		// Generate encoded token and send it as response
		t, err := token.SignedString(signingKey)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})

	}

	return echo.ErrUnauthorized
}

package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" // Import PostgreSQL driver
)

func ConnectDB() (*sql.DB, error) {

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	koneksi := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUsername, dbPassword, dbName, dbPort)

	db, err := sql.Open("postgres", koneksi)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping the database: %w", err)
	}

	return db, nil
}

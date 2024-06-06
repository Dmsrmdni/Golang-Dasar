package main

import (
	"os"
	"week5/routes"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		return
	}

	//Routing
	e := routes.Init()

	appPort := os.Getenv("APP_PORT")

	// Mulai Server
	e.Logger.Fatal(e.Start(":" + appPort))
}

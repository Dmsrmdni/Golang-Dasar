package main

import (
	"week4/routes"

	_ "github.com/lib/pq"
)

func main() {
	//Routing
	e := routes.Init()

	// Mulai Server
	e.Logger.Fatal(e.Start(":8000"))
}

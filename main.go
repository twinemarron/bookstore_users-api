package main

import (
	"github.com/joho/godotenv"
	"github.com/twinemarron/bookstore_users-api/app"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	app.StartApplication()
}

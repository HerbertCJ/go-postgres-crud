package main

import (
	"log"

	"github.com/HerbertCJ/my-store/app"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	var a app.App
	a.CreateConeection()
	a.Routes()
	a.Run()
}

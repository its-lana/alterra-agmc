package main

import (
	"Task-Code/config"
	"Task-Code/routes"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":3000"))
}

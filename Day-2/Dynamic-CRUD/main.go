package main

import (
	"Dynamic-CRUD/config"
	"Dynamic-CRUD/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":3000"))
}

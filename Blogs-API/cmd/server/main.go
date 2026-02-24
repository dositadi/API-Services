package main

import (
	a "blog/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	app := &a.App{}

	app.InitializeDB()
	app.Run()
}

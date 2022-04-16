package main

import (
	"database/sql"
	db2 "github.com/loxt/fullcycle2.0-ports-and-adapters/adapters/db"
	"github.com/loxt/fullcycle2.0-ports-and-adapters/application"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)
	product, _ := productService.Create("test", 30)

	_, _ = productService.Enable(product)
}

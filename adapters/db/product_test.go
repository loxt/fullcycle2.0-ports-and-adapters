package db

import (
	"database/sql"
	"log"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")

	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	createTableQuery := `CREATE TABLE products (
                 "id" string,
                 "name" string,
                 "price" float,
                 "status" string
								);`

	stmt, err := db.Prepare(createTableQuery)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, _ = stmt.Exec()
}
func createProduct(db *sql.DB) {
	insertProductQuery := `insert into products values ("abc", "Product Test", 0, "disabled");`
	stmt, err := db.Prepare(insertProductQuery)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, _ = stmt.Exec()
}

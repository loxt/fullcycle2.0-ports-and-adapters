package db_test

import (
	"database/sql"
	"github.com/loxt/fullcycle2.0-ports-and-adapters/adapters/db"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
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

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)

	product, err := productDb.Get("abc")
	require.Nil(t, err)

	require.Equal(t, "Product Test", product.GetName())
	require.Equal(t, float64(0), product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

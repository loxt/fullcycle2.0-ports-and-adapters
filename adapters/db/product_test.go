package db_test

import (
	"database/sql"
	"github.com/loxt/fullcycle2.0-ports-and-adapters/adapters/db"
	"github.com/loxt/fullcycle2.0-ports-and-adapters/application"
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

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)

	product := application.NewProduct()
	product.Name = "Product Test"
	product.Price = 25.0

	productResult, err := productDb.Save(product)
	require.Nil(t, err)

	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())

	product.Status = "enabled"
	productResult, err = productDb.Save(product)

	require.Nil(t, err)
	require.Equal(t, product.Status, productResult.GetStatus())

}

package main

import (
	"database/sql"

	db2 "github.com/fsgabriel/go-hexagonal-architecture/adapters/db"
	"github.com/fsgabriel/go-hexagonal-architecture/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", ":memory")
	if err != nil {
		panic(err)
	}

	productDB := db2.NewProductDB(db)
	productService := application.NewProductService(productDB)
}

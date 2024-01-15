package main

import (
	"database/sql"
	"fmt"

	db2 "github.com/fsgabriel/go-hexagonal-architecture/adapters/db"
	"github.com/fsgabriel/go-hexagonal-architecture/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", ":memory")
	if err != nil {
		panic(err)
	}
	db2.Setup(db)
	productDB := db2.NewProductDB(db)
	productService := application.NewProductService(productDB)

	p, _ := productService.Create("product exemplo", 10)
	x, _ := productService.Get(p.GetID())
	fmt.Println(x)
}

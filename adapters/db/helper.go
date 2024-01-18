package db

import (
	"database/sql"
	"log"
)

func Setup(db *sql.DB) {
	CreateTable(db)
	CreateProduct(db)
}

func CreateTable(db *sql.DB) {
	table := `CREATE TABLE products (
			"id" string,
			"name" string,
			"price" float,
			"status" string
			);`
	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func CreateProduct(db *sql.DB) {
	insert := `insert into products values("abc","Product Test",0,"disabled")`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

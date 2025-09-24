package main

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "./test.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	usecase := NewUseCase(db)

	product, err := usecase.GetProductName(1)
	if err != nil {
		panic(err)
	}
	println(product.Name)
}

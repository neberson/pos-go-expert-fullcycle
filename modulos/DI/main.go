package main

import (
	"database/sql"

	"github.com/neberson/pos-go-expert-fullcycle/modulos/DI/product"
	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "./test.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Create a new product repository
	repository := product.NewProductRepository(db)

	// Create a new product use case
	usecase := product.NewProductUseCase(repository)

	product, err := usecase.GetProductName(1)
	if err != nil {
		panic(err)
	}
	println(product.Name)
}

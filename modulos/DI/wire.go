//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/neberson/pos-go-expert-fullcycle/modulos/DI/product"
)

var (
	setRepositoryDependencies = wire.NewSet(
		product.NewProductRepository,
		wire.Bind(new(product.ProductRepositoryInterface), new(*product.ProductRepository)),
	)
)

func NewUseCase(db *sql.DB) *product.ProductUseCase {
	wire.Build(
		setRepositoryDependencies,
		product.NewProductUseCase,
	)
	return &product.ProductUseCase{}
}

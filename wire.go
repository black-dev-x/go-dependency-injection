//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/black-dev-x/go-dependency-injection/product"
	"github.com/google/wire"
)

var setRepository = wire.NewSet(
	product.NewProductRepository,
	wire.Bind(new(product.ProductRepositoryInterface), new(*product.ProductRepository)),
)

func NewUseCase(db *sql.DB) *product.ProductUseCase {
	wire.Build(
		setRepository,
		product.NewProductUseCase,
	)
	return &product.ProductUseCase{}
}

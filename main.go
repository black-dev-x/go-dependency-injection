package main

import (
	"database/sql"

	"github.com/black-dev-x/go-dependency-injection/product"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	repository := product.NewProductRepository(db)
	useCase := product.NewProductUseCase(repository)

	product := useCase.GetProductByID(3)
	println("Product ID:", product.ID)
	println("Product Name:", product.Name)
}

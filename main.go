package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	useCase := NewUseCase(db)

	product := useCase.GetProductByID(3)
	println("Product ID:", product.ID)
	println("Product Name:", product.Name)
}

package main

import (
	"fmt"
	"github/carlosmajimo/isbn-validator/isbn"
)

func main() {
	isbnValue := "0471958697"
	validator := isbn.NewValidateISBN(isbnValue)

	if validator.IsValid() {
		fmt.Printf("El ISBN %s es válido\n", isbnValue)
	} else {
		fmt.Printf("El ISBN %s no es válido\n", isbnValue)
	}
}

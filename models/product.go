package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name     string
	Desc     string
	Price    float64
	Quantity int
}

func CreateNewProduct(name, description string, price float64, quantity int) Product {
	var newProduct = Product{
		Name:     name,
		Desc:     description,
		Price:    price,
		Quantity: quantity}

	return newProduct
}

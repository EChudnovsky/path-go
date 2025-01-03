package main

import "fmt"

// Type: ProductList
type Product struct {
	name, category string
	price          float64
}

func newProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

func (product *Product) printDetails() {
	fmt.Println("Name:", product.name, "Category:", product.category, "Price", product.calcTax(0.2, 100))
}

func (product *Product) calcTax(rate, threshold float64) float64 {
	if product.price > threshold {
		return product.price + (product.price * rate)
	}
	return product.price
}

// Методы интерфейса Expense
func (p Product) getName() string {
	return p.name
}
func (p Product) getCost(_ bool) float64 {
	return p.price
}

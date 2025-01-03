package main

import (
	"fmt"
	_ "packages/data"
	. "packages/fmt"
	"packages/store"
	"packages/store/cart"

	"github.com/fatih/color"
)

func main() {

	product := store.NewProduct("Kayak", "Watersports", 279)

	fmt.Println("Name:", product.Name)
	fmt.Println("Category:", product.Category)
	fmt.Println("Price:", ToCurrency(product.Price()))

	cart := cart.Cart{
		CustomerName: "Alice",
		Products:     []store.Product{*product},
	}
	color.Green("Name: " + cart.CustomerName)
	color.Cyan("Total: " + ToCurrency(cart.GetTotal()))
}

package main

import "fmt"

func printPrice(product string, price, taxRate float64, order ...string) {

	taxAmount := price * taxRate
	fmt.Println("Price:", product, "Tax:", taxAmount)
	for index, value := range order {
		fmt.Println(index, value)
	}
}

func main() {
	printPrice("Kayak", 275, 0.2, "pupkin", "bubukin")
	names := []string{"Acme Kayaks", "Bob's Boats", "Crazy Canoes"}

	printPrice("Lifejacket", 48.95, 0.2, names...)
	printPrice("Soccer Ball", 19.50, 0.15)
}

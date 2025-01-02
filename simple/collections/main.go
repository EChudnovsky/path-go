package main

import (
	"fmt"
)

func main() {

	product := map[string]bool{
		"Kayak":      true,
		"Lifejacket": true,
		"Hat":        true,
	}

	value, ok := product["Hat1"]

	fmt.Println("Stored value:", value, ok)
}

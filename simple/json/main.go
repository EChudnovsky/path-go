package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Product struct {
	Name, Category string
	Price          float64
}

func main() {
	var b bool = true
	var str string = "Hello"
	var fval float64 = 99.99
	var ival int = 200
	var pointer *int = &ival

	names := []string{"Kayak", "Lifejacket", "Soccer Ball"}
	numbers := [3]int{10, 20, 30}
	var byteArray [5]byte

	copy(byteArray[0:], []byte(names[0]))
	byteSlice := []byte(names[0])

	m := map[string]float64{
		"Kayak":      279,
		"Lifejacket": 49.95,
	}

	var Kayak = Product{
		Name:     "Kayak",
		Category: "Watersports",
		Price:    279,
	}

	var writer strings.Builder

	dp := DiscountedProduct{
		Product:  &Kayak,
		Discount: 10.50,
	}
	namedItems := []Named{&dp, &Person{PersonName: "Alice"}}
	encoder := json.NewEncoder(&writer)

	for _, val := range []interface{}{b, str, fval, ival, pointer, names, numbers, byteSlice, m, Kayak, namedItems} {
		encoder.Encode(val)
	}
	fmt.Print(writer.String())
}

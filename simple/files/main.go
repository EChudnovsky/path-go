package main

import "fmt"

func main() {
	fmt.Println("Read file config...",
		"\n___________________________")

	for _, p := range Products {
		Printfln("Product: %v, Category: %v, Price: $%.2f",
			p.Name, p.Category, p.Price)
	}
	fmt.Println("Example  write file...",
		"\n___________________________")

	writefile()

}

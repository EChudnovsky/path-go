package main

// Type: ProductList
type ProductList []*Product

/*
Exsemple:
	products := ProductList{
		newProduct("Kayak", "Watersports", 275),
		newProduct("Lifejacket", "Watersports", 48.95),
		newProduct("Soccer Ball", "Soccer", 19.50),
	}

	for category, total := range products.calcCategoryTotals() {
		products.calcCategoryTotals()
		fmt.Println("Category: ", category, "Total:", total)
	}
*/
func (products *ProductList) calcCategoryTotals() map[string]float64 {
	totals := make(map[string]float64)
	for _, p := range *products {
		totals[p.category] = totals[p.category] + p.price
	}
	return totals
}

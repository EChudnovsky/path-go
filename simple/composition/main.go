package main

import (
	"composition/store"
	"fmt"
)

func compLevel2() {
	boats := []*store.Boat{
		store.NewBoat("Kayak", 275, 1, false),
		store.NewBoat("Canoe", 400, 3, false),
		store.NewBoat("Tender", 650.25, 2, true),
	}
	for _, b := range boats {
		fmt.Println("Boat:", b.Name, "Price:", b.Price(0.2))
	}
}

func compLevel3() {
	rentals := []*store.RentalBoat{
		store.NewRentalBoat(
			"Rubber  Ring",
			10,
			1,
			false,
			false,
			"",
			""),
		store.NewRentalBoat(
			"Yacht",
			50000,
			5,
			true,
			true,
			"",
			""),
		store.NewRentalBoat(
			"Super Yacht",
			100000,
			15,
			true,
			true,
			"",
			""),
	}

	for _, r := range rentals {
		fmt.Println("Rental Boat:", r.Name, "Rental Price:", r.Price(0.2))
	}
}

func compLevel3AndSomeFields() {
	rentals := []*store.RentalBoat{
		store.NewRentalBoat(
			"Rubber  Ring",
			10,
			1,
			false,
			false,
			"N/A",
			"N/A"),
		store.NewRentalBoat(
			"Yacht",
			50000,
			5,
			true,
			true,
			"Bob",
			"Alice"),
		store.NewRentalBoat(
			"Super Yacht",
			100000,
			15,
			true,
			true,
			"Dora",
			"Charlie"),
	}
	for _, r := range rentals {
		fmt.Println(
			"Rental Boat:", r.Name,
			"Rental Price:", r.Price(0.2),
			"Captain:", r.Captain)
	}
}

func compInterface() {
	products := map[string]store.ItemForSale{
		"Kayak": store.NewBoat("Kayak", 279, 1, false),
		"Ball": store.NewProduct("Soccer  Ball", "Soccer",
			19.50),
	}
	for key, p := range products {
		switch item := p.(type) {
		case store.Describable:
			fmt.Println(
				"Name:", item.GetName(),
				"Category:", item.GetCategory(),
				"Price:", item.(store.ItemForSale).Price(0.2))
		default:
			fmt.Println("Key:", key, "Price:", p.Price(0.2))
		}
	}

}
func main() {

	fmt.Println(
		"_____________________________\n",
		"Composition level 2")
	compLevel2()

	fmt.Println(
		"_____________________________\n",
		"Composition level 3")
	compLevel3()

	fmt.Println(
		"_____________________________\n",
		"Composition level 3 with some fields")

	compLevel3AndSomeFields()

	fmt.Println(
		"_____________________________\n",
		"Composition with Interface")
	compInterface()
}

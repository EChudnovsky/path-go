package main

import "fmt"

func expWorkInterface() {

	expenses := ExpenseList{
		newProduct("Kayak", "Watersports", 275),
		newService("Boat Cover", 12, 89.50, []string{}),
	}

	type Account struct {
		accountNumber int
		expenses      *ExpenseList
	}

	account := Account{
		accountNumber: 12345,
		expenses:      &expenses,
	}

	fmt.Println("Input ExpenseList:")
	expenses.printList()
	fmt.Println("Total:", expenses.calcTotal())

	fmt.Println("________________________")
	fmt.Println("Input Account:")
	for _, expense := range *(account.expenses) {
		fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
	}
	fmt.Println("Total:", account.expenses.calcTotal())

}

func emptyInterface() {

	var expense Expense = &Product{"Kayak", "Watersports", 275}
	data := []interface{}{
		expense,
		Product{"Lifejacket", "Watersports", 48.95},
		Service{"Boat Cover", 12, 89.50, []string{}},
		Person{"Alice", "London"},
		&Person{"Bob", "New York"},
		"This is a string",
		100,
		true,
	}

	for _, item := range data {
		switch value := item.(type) {
		case Product:
			fmt.Println("Product:", value.name, "Price:", value.price)
		case *Product:
			fmt.Println("Product Pointer:", value.name,
				"Price:", value.price)
		case Service:
			fmt.Println("Service:", value.description,
				"Price:", value.monthlyFee*float64(value.durationMonths))
		case Person:
			fmt.Println("Person:", value.name, "City:", value.city)
		case *Person:
			fmt.Println("Person Pointer:", value.name, "City:", value.city)
		case string, bool, int:
			fmt.Println("Built-in type:", value)
		default:
			fmt.Println("Default:", value)
		}
	}

}

func processItems(items ...interface{}) {
	for _, item := range items {

		switch value := item.(type) {
		case Product:
			fmt.Println("Product:", value.name, "Price:",
				value.price)
		case *Product:
			fmt.Println("Product  Pointer:", value.name,
				"Price:", value.price)
		case Service:
			fmt.Println("Service:", value.description,
				"Price:",
				value.monthlyFee*
					float64(value.durationMonths))
		case Person:
			fmt.Println("Person:", value.name, "City:",
				value.city)
		case *Person:
			fmt.Println("Person  Pointer:", value.name,
				"City:", value.city)
		case string, bool, int:
			fmt.Println("Built-in type:", value)
		default:
			fmt.Println("Default:", value)
		}
	}
}

func main() {

	fmt.Println("Example work interface:")
	expWorkInterface()

	fmt.Println("_____________________\n",
		"Example work empty interface:")

	emptyInterface()

	fmt.Println("_____________________\n",
		"Example work empty interface in function:")

	var expense Expense = &Product{"Kayak", "Watersports", 275}
	data := []interface{}{
		expense,
		Product{"Lifejacket", "Watersports", 48.95},
		Service{"Boat Cover", 12, 89.50, []string{}},
		Person{"Alice", "London"},
		&Person{"Bob", "New York"},
		"This is a string",
		100,
		true,
	}

	// for _, item := range data {
	// 	processItems(item)
	// }
	processItems(data...)

}

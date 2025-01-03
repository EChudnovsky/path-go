package main

import "fmt"

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

type ExpenseList []Expense

func (expenses *ExpenseList) calcTotal() (total float64) {
	for _, item := range *expenses {
		total += item.getCost(true)
	}

	return total
}
func (expenses *ExpenseList) printList() {

	for _, expense := range *expenses {
		fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
	}

}

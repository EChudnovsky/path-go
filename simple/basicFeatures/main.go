package main

import (
	"fmt"
)

func main() {

	first := 100
	second := &first

	first++
	*second++

	var myNewPointer *int

	fmt.Println("first:", first)
	fmt.Println("second:", *second)
	fmt.Println("second:", myNewPointer)

}

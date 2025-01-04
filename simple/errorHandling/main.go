package main

import "fmt"

func totalPrice() {

	categories := []string{"Watersports", "Chess", "Running"}
	for _, cat := range categories {
		total, err := Products.TotalPrice(cat)
		if err == nil {
			fmt.Println(cat, "Total:", ToCurrency(total))
		} else {
			fmt.Println(cat, "(no such category)")
		}
	}
}

func totalPriceAsync() {

	recoveryFunc := func() {
		if arg := recover(); arg != nil {
			if err, ok := arg.(error); ok {
				fmt.Println("Error:", err.Error())

			} else if str, ok := arg.(string); ok {
				fmt.Println("Message:", str)
			} else {
				fmt.Println("Panic recovered")
			}
		}
	}

	defer recoveryFunc()

	categories := []string{"Watersports", "Chess", "Running"}

	channel := make(chan ChannelMessage, 10)
	go Products.TotalPriceAsync(categories, chan<- ChannelMessage(channel))
	for message := range channel {
		if message.CategoryError == nil {
			fmt.Println(message.Category, "Total:", ToCurrency(message.Total))
		} else {
			fmt.Println(message.CategoryError.Error())
		}
	}

}

func runProcess() {
	categories := []string{"Watersports", "Chess", "Running"}
	channel := make(chan CategoryCountMessage)
	go processCategories(categories, channel)
	for message := range channel {
		if message.TerminalError == nil {
			fmt.Println(message.Category, "Total:", message.Count)
		} else {
			fmt.Println("A terminal error occured")
		}
	}

}

func main() {

	// totalPrice()
	// totalPriceAsync()
	runProcess()
}

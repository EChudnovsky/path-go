package main

import "fmt"

// type CategoryError struct {
// 	requestedCategory string
// }

type ChannelMessage struct {
	Category      string
	Total         float64
	CategoryError error
}

// func (e *CategoryError) Error() string {
// 	return "Category  " + e.requestedCategory + "  does  not	exist"
// }

func (slice ProductSlice) TotalPrice(category string) (total float64, err error) {
	productCount := 0
	for _, p := range slice {
		if p.Category == category {
			total += p.Price
			productCount++
		}
	}
	if productCount == 0 {
		// err = &CategoryError{requestedCategory: category}
		err = fmt.Errorf("Cannot  find  category:  %v", category)

	}

	return total, err
}

func (slice ProductSlice) TotalPriceAsync(categories []string, channel chan<- ChannelMessage) {
	for _, c := range categories {
		total, err := slice.TotalPrice(c)
		channel <- ChannelMessage{
			Category:      c,
			Total:         total,
			CategoryError: err,
		}
	}
	close(channel)
}

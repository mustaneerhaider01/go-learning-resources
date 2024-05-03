package lists

import "fmt"

func lists() {
	prices := []float64{10.99, 7.99}
	fmt.Println(prices[0:1])

	prices[1] = 9.99
	// prices[2] = 6.99 -> not allowed to assigned values to indexes that doesn't exist in case of dynamic lists created via slice
	prices = append(prices, 7.99)
	fmt.Println(prices)

	discountedPrices := []float64{101.99, 75.99, 355.19}
	prices = append(prices, discountedPrices...) // spread operator to extract list elements and pass them as separated elements to append()

	fmt.Println(prices)

}

// func main() {
// 	var productNames [4]string = [4]string{"A Bike"}
// 	prices := [4]float64{9.99, 10.99, 5.59, 7.99}
// 	fmt.Println(prices)

// 	productNames[2] = "A Book"

// 	fmt.Println(productNames)

// 	fmt.Println(prices[2])

// 	featuredPrices := prices[1:] // [10.99, 5.59, 7.99] -> Slice is reference to part of original array
// 	featuredPrices[0] = 100.99
// 	highlightedPrices := featuredPrices[:1] // [10.99] -> This slice is created based on another slice
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(prices)
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

// 	highlightedPrices = highlightedPrices[:3]
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
// }

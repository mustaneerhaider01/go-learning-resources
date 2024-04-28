package main

import "fmt"

func main() {
	var productNames [4]string = [4]string{"A Bike"}
	prices := [4]float64{9.99, 10.99, 5.59, 7.99}
	fmt.Println(prices)

	productNames[2] = "A Book"

	fmt.Println(productNames)

	fmt.Println(prices[2])

	featuredPrices := prices[1:]            // [10.99, 5.59, 7.99]
	highlightedPrices := featuredPrices[:1] // [10.99] -> This slice is created based on another slice
	fmt.Println(highlightedPrices)
}

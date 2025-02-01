package main

import "fmt"

func main() {
	var productNames [4]string = [4]string{"A book"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)
	fmt.Println(prices[3])

	fmt.Println(productNames)
	productNames[3] = "A carpet"
	fmt.Println(productNames)

	featuredPrices := prices[1:]
	featuredPrices[0] = 199.99
	highlitedPrices := featuredPrices[:1]
	fmt.Println(featuredPrices)
	fmt.Println(highlitedPrices)
	fmt.Println(len(highlitedPrices), cap(highlitedPrices))

	highlitedPrices = highlitedPrices[:3]
	fmt.Println(highlitedPrices)
	fmt.Println(len(highlitedPrices), cap(highlitedPrices))
}

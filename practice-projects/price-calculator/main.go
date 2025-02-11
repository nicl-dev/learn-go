package main

import (
	"fmt"

	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.10, 0.15}

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println(err)
			return
		}

	}
}

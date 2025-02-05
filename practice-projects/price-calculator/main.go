package main

import "fmt"

func main() {
	prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.07, 0.10, 0.15}

	result := make(map[float64][]float64)

	for _, price := range prices {
		for _, taxRate := range taxRates {
			result[taxRate] = append(result[taxRate], price*(1+taxRate))
		}
	}

	fmt.Println(result)
}

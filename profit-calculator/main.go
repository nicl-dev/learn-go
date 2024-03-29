package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var tax_rate float64

	fmt.Print("Enter your revenue in dollars: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter your expenses in dollars: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter your tax rate: ")
	fmt.Scan(&tax_rate)

	var ebt = revenue - expenses
	var profit = ebt * (1 - tax_rate/100)
	var ratio = ebt / profit

	fmt.Printf("Your earnings before taxes are: $%v \n", ebt)
	fmt.Printf("Your profit after taxes is: $%v \n", profit)
	fmt.Printf("Your ebt to profit ratio is: %v \n", ratio)
}

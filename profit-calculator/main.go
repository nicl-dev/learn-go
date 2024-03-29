package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Enter your revenue in dollars: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter your expenses in dollars: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter your tax rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := revenue*(1-taxRate/100) - expenses
	ratio := ebt / profit

	fmt.Printf("Your earnings before taxes are: $%v \n", ebt)
	fmt.Printf("Your profit after taxes is: $%v \n", profit)
	fmt.Printf("Your ebt to profit ratio is: %v%% \n", ratio)
}

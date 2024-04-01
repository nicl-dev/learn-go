package main

import "fmt"

func main() {
	revenue := getUserInput("Enter your revenue in dollars: ")
	expenses := getUserInput("Enter your expenses in dollars: ")
	taxRate := getUserInput("Enter your tax rate: ")

	ebt, profit, ratio := calculateAll(revenue, expenses, taxRate)

	fmt.Printf("Your earnings before taxes are: $%.3v\n", ebt)
	fmt.Printf("Your profit after taxes is: $%.3v\n", profit)
	fmt.Printf("Your ebt to profit ratio is: %.3v\n", ratio)
}

func getUserInput(infoText string) float64 {
	var userInput float64

	fmt.Print(infoText)
	fmt.Scan(&userInput)

	return userInput
}

func calculateAll(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = revenue*(1-taxRate/100) - expenses
	ratio = ebt / profit

	return ebt, profit, ratio
}

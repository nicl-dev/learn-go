package main

import (
	"fmt"
	"os"
)

func main() {
	revenue := getUserInput("Enter your revenue in dollars: ")
	expenses := getUserInput("Enter your expenses in dollars: ")
	taxRate := getUserInput("Enter your tax rate: ")

	ebt, profit, ratio := calculateAll(revenue, expenses, taxRate)
	storeResults(ebt, profit, ratio)

	fmt.Printf("Your earnings before taxes are: $%.2f\n", ebt)
	fmt.Printf("Your profit after taxes is: $%.2f\n", profit)
	fmt.Printf("Your ebt to profit ratio is: %.2f\n", ratio)
}

func getUserInput(infoText string) float64 {
	var userInput float64

	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		panic("only positive numbers allowed")
	}

	return userInput
}

func calculateAll(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = revenue*(1-taxRate/100) - expenses
	ratio = ebt / profit

	return ebt, profit, ratio
}

func storeResults(ebt, profit, ratio float64) {
	text := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f", ebt, profit, ratio)
	os.WriteFile("profit.txt", []byte(text), 0644)
}

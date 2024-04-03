package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err := getUserInput("Enter your revenue in dollars: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	expenses, err := getUserInput("Enter your expenses in dollars: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	taxRate, err := getUserInput("Enter your tax rate: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, profit, ratio := calculateAll(revenue, expenses, taxRate)
	writeToFile(ebt, profit, ratio)

	fmt.Printf("Your earnings before taxes are: $%.3v\n", ebt)
	fmt.Printf("Your profit after taxes is: $%.3v\n", profit)
	fmt.Printf("Your ebt to profit ratio is: %.3v\n", ratio)
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64

	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("only positive numbers allowed")
	}

	return userInput, nil
}

func calculateAll(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = revenue*(1-taxRate/100) - expenses
	ratio = ebt / profit

	return ebt, profit, ratio
}

func writeToFile(ebt, profit, ratio float64) {
	text := fmt.Sprint(ebt, profit, ratio)
	os.WriteFile("profit.txt", []byte(text), 0644)
}

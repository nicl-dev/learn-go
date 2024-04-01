package main

import "fmt"

func main() {
	var accountBalance float64 = 1000

	fmt.Println("Welcome to the Bank!")
	fmt.Println("Please select an option:")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	fmt.Print("Enter your choice: ")
	var choice int
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Your balance is: ", accountBalance)
	} else if choice == 2 {
		var depositAmount float64
		fmt.Print("Enter the amount you want to deposit: ")
		fmt.Scan(&depositAmount)

		if depositAmount <= 0 {
			fmt.Println("Invalid deposit amount. Must be greater than 0.")
			return
		}

		accountBalance += depositAmount
		fmt.Println("Your new balance is: ", accountBalance)
	} else if choice == 3 {
		var withdrawAmount float64
		fmt.Print("Enter the amount you want to withdraw: ")
		fmt.Scan(&withdrawAmount)

		if withdrawAmount >= accountBalance {
			fmt.Println("Invalid deposit amount. Not enough funds in account.")
			return
		}

		accountBalance -= withdrawAmount
		fmt.Println("Your new balance is: ", accountBalance)
	} else {
		fmt.Println("Goodbye!")
	}
}

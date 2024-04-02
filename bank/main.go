package main

import "fmt"

func main() {
	var accountBalance float64 = 1000
	var choice int

	fmt.Println("Welcome to the Bank!")

	for {
		fmt.Println("Please select an option:")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Enter the amount you want to deposit: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid deposit amount. Must be greater than 0.")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Your new balance is:", accountBalance)
		case 3:
			var withdrawAmount float64
			fmt.Print("Enter the amount you want to withdraw: ")
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid deposit amount. Must be greater than 0.")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Invalid deposit amount. Not enough funds in account.")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Println("Your new balance is: ", accountBalance)
		default:
			fmt.Println("Thank you for choosing our Bank!")
			return
		}
	}
}

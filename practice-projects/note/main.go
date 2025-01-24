package main

import (
	"errors"
	"fmt"
)

func main() {
	_, err := getUserInput("Enter your note.")

	if err != nil {
		fmt.Println(err)
		return
	}
}

func getUserInput(prompt string) (string, error) {
	fmt.Println(prompt)

	var value string
	_, err := fmt.Scanln(&value)
	if err != nil {
		return "", errors.New("Invalid input.")
	}

	return value, nil
}

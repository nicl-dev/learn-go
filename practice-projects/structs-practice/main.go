package main

import (
	"errors"
	"fmt"
)

func main() {
	title, err := getUserInput("Note title:")

	if err != nil {
		fmt.Println(err)
		return
	}

	content, err := getUserInput("Note content:")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(title)
	fmt.Println(content)
}

func getUserInput(prompt string) (string, error) {
	fmt.Println(prompt)
	var value string
	fmt.Scanln(&value)
	if value == "" {
		return "", errors.New("invalid input")
	}
	return value, nil
}

package main

import (
	"errors"
	"fmt"
)

func main() {
	title, content, err := getNoteData()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(title)
	fmt.Println(content)
}

func getNoteData() (string, string, error) {
	title, err := getUserInput("Enter a title:")
	if err != nil {
		return "", "", err
	}

	content, err := getUserInput("Enter the content:")
	if err != nil {
		return "", "", err
	}

	return title, content, nil
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

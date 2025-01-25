package main

import (
	"fmt"

	"example.com/note/note"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(userNote)
}

func getNoteData() (string, string) {
	title := getUserInput("Enter a title:")
	content := getUserInput("Enter the content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Println(prompt)

	var value string
	fmt.Scanln(&value)

	return value
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Enter todo:")

	userTodo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	userTodo.Display()

	err = userTodo.Save()
	if err != nil {
		fmt.Println("Saving the todo failed.")
	}
	fmt.Println("Saving the todo succeeded!")

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()

	err = userNote.Save()
	if err != nil {
		fmt.Println("Saving the note failed.")
		return
	}

	fmt.Println("Saved Note!")
}

func getNoteData() (string, string) {
	title := getUserInput("Enter a title:")
	content := getUserInput("Enter the content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Println(prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

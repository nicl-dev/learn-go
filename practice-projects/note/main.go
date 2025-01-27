package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Enter todo:")

	userTodo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	userTodo.Display()

	err = saveData(userTodo)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Saving the todo succeeded!")

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = saveData(userNote)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		return err
	}

	fmt.Println("Saved Note!")
	return nil
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

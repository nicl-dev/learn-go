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

type outputtable interface {
	saver
	Display()
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Enter todo:")

	userTodo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userTodo)
	if err != nil {
		fmt.Println(err)
		return
	}
	printSomething("Saving the todo succeeded!")

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userNote)
	if err != nil {
		printSomething(err)
		return
	}
	printSomething("Saving the note succeeded!")
}

func printSomething(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("Integer:", value)
	case float64:
		fmt.Println("Float64:", value)
	case string:
		fmt.Println(value)
	}
}

func outputData(data outputtable) error {
	data.Display()
	err := saveData(data)
	if err != nil {
		return err
	}
	return nil
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		return err
	}
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

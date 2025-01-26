package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (t Todo) Display() {
	fmt.Println(t.Text)
}

func (t Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(t)
	if err != nil {
		return err
	}

	err = os.WriteFile(fileName, json, 0644)
	if err != nil {
		return err
	}
	return nil
}

func New(todoContent string) (Todo, error) {
	if todoContent == "" {
		return Todo{}, errors.New("Invalid input.")
	}

	return Todo{
		Text: todoContent,
	}, nil
}

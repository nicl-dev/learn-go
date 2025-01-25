package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (n Note) Display() {
	fmt.Printf("%v\n\n%v\n", n.Title, n.Content)
}

func (n Note) Save() error {
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(n)
	if err != nil {
		return err
	}

	err = os.WriteFile(fileName, json, 0644)
	if err != nil {
		return err
	}
	return nil
}

func New(noteTitle, noteContent string) (Note, error) {
	if noteTitle == "" || noteContent == "" {
		return Note{}, errors.New("Invalid input.")
	}

	return Note{
		Title:     noteTitle,
		Content:   noteContent,
		CreatedAt: time.Now(),
	}, nil
}

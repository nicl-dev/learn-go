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
	title     string
	content   string
	createdAt time.Time
}

func (n Note) Display() {
	fmt.Printf("%v\n\n%v\n", n.title, n.content)
}

func (n Note) Save() error {
	fileName := strings.ReplaceAll(n.title, " ", "_")
	fileName = strings.ToLower(fileName)

	json, err := json.Marshal(n)
	if err != nil {
		return err
	}

	err = os.WriteFile(fmt.Sprintf("testdata/%s.json", fileName), json, 0644)
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
		title:     noteTitle,
		content:   noteContent,
		createdAt: time.Now(),
	}, nil
}

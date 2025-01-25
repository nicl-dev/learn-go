package note

import (
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
	fmt.Printf("%v\n\n%v", n.title, n.content)
}

func (n Note) Save() {
	fileName := strings.ReplaceAll(n.title, " ", "_")
	fileName = strings.ToLower(fileName)
	err := os.WriteFile(fmt.Sprintf("testdata/%s", fileName), []byte(n.title+"\n"+n.content), 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
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

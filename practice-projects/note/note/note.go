package note

import (
	"errors"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
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

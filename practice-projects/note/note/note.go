package note

import "time"

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func New(noteTitle, noteContent string) Note {
	return Note{
		title:     noteTitle,
		content:   noteContent,
		createdAt: time.Now(),
	}
}

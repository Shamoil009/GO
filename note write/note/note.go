package note

import (
	"fmt"
	"time"
)

type Note struct {
	title      string
	content    string
	created_at time.Time
}

func (note Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n\n%v\n\n", note.title, note.content)
}

func NewNote(title, content string) (*Note, error) {
	return &Note{
		title:      title,
		content:    content,
		created_at: time.Now(),
	}, nil
}

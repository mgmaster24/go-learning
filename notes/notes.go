package gol_notes

import (
	"errors"
	"fmt"
	"time"

	gol_io "go-learning.com/learning/io"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type Notes []Note

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Invalid Note values")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (notes Notes) Save(fileName string) error {
	return gol_io.SaveJsonData(fileName, notes)
}

func Load(fileName string) (Notes, error) {
	var notes Notes
	notes, err := gol_io.LoadJsonData[Notes](fileName)
	if err != nil {
		return notes, err
	}

	return notes, nil
}

func (notes *Notes) Add(note Note) {
	*notes = append(*notes, note)
}

func (note Note) Print() {
	fmt.Printf("Title: %v\n", note.Title)
	fmt.Printf("Content: %v\n", note.Content)
	fmt.Printf("Created: %v\n", note.CreatedAt)
}

func (notes Notes) Print() {
	for i := 0; i < len(notes); i++ {
		notes[i].Print()
	}
}

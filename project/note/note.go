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
	Title       string
	Description string
	CreatedAt   time.Time
}

func New(title, desc string) (*Note, error) {
	if title == "" || desc == "" {
		return &Note{}, errors.New("user inputs cant be empty")
	}
	return &Note{
		Title:       title,
		Description: desc,
		CreatedAt:   time.Now(),
	}, nil
}

func (n Note) Display() {
	fmt.Printf("\n Title : %v \n Description : %v \n Timestamp : %v\n", n.Title, n.Description, n.CreatedAt)
}

func (note Note) SaveNote() error {
	fileName := strings.ToLower(strings.ReplaceAll(note.Title, " ", "_")) + ".json"
	jsonObject, err := json.Marshal(note)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonObject, 0644)
}

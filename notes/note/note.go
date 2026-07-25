package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("all fields are required.")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (n Note) SaveJSON() error {
	title := fmt.Sprintf(`%s.json`, strings.ToLower(strings.ReplaceAll(n.Title, " ", "_")))
	data, err := json.Marshal(n)
	if err != nil {
		log.Print("error encoding json")
	}

	return os.WriteFile(title, data, 0644)
}

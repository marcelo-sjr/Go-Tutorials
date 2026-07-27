package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("all fields are required.")
	}

	return Todo{
		Text: content,
	}, nil
}

func (n Todo) SaveJSON() error {
	title := "todo.json"
	data, err := json.Marshal(n)
	if err != nil {
		log.Print("error encoding json")
	}

	return os.WriteFile(title, data, 0644)
}

func (n Todo) Print() {
	fmt.Printf("Text: %s\n", n.Text)
}

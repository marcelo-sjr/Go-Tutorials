package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"example/note"
)

func getUserInput(r *bufio.Reader) (string, error) {
	str, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}
	str = strings.TrimSuffix(str, "\n")
	str = strings.TrimSuffix(str, "\r")

	return str, nil
}

func main() {
	log.SetFlags(0)
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Title: ")
	title, err := getUserInput(reader)
	if err != nil {
		log.Printf("Error: %v", err)
	}

	fmt.Print("Content: ")
	content, err := getUserInput(reader)
	if err != nil {
		log.Printf("Error: %v", err)
	}

	note, err := note.New(title, content)
	if err != nil {
		log.Printf("Error: %v", err)
	}

	err = note.SaveJSON()
	if err != nil {
		log.Printf("Error: %v", err)
	}

	fmt.Println("Note saved!")
}

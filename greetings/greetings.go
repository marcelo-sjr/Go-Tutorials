package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name argument can't be empty!")
	}
	message := fmt.Sprintf("Hello, %s. Welcome!", name)
	return message, nil
}

package greetings

import "fmt"

func Hello(name string) string{
	message := fmt.Sprintf("Hello, %s. Welcome!", name)
	return message
}


package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name argument can't be empty!")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

// starting functions with lowercase turns it into a non exported func(not visible outside package)
func randomFormat() string {
	possibleGreetings := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
		"Hey, %v! Nice to see you!",
		"Welcome, %v!",
		"Greetings, %v!",
		"Good to have you here, %v!",
		"It's great to see you, %v!",
		"Hope you're doing well, %v!",
		"Howdy, %v!",
		"Good day, %v!",
		"Cheers, %v!",
		"Welcome aboard, %v!",
		"Happy to see you, %v!",
		"Hi there, %v!",
		"Hello and welcome, %v!",
		"Nice to meet you, %v!",
		"Hey there, %v!",
		"Pleasure to have you here, %v!",
		"Warm greetings, %v!",
		"Have a wonderful day, %v!",
	}

	return possibleGreetings[rand.Intn(len(possibleGreetings))]
}

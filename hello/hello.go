package main

import (
	"fmt"
	_ "fmt"
	"log"
	_ "log"

	"github.com/marcelo-sjr/Go-Tutorials/greetings"
)

func main() {
	name := "Alex"
	message, err := greetings.Hello(name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	names := []string{
		"Alex",
		"Brian",
		"Roger",
		"Melissa",
		"Emma",
	}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	//'message' local to this for loop, so it shadows the previous
	//'message' variable without a redefinition error
	for _, message := range messages {
		fmt.Println(message)
	}
}

package main

import (
	"fmt"
	"log"

	"github.com/marcelo-sjr/Go-Tutorials/greetings"
)

func main() {

	name := "Alex"
	message, err := greetings.Hello(name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}

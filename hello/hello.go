package main

import (
	"fmt"

	"example.com/greetings"

	"rsc.io/quote"

	"log"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	message, _ := greetings.Hello("Me")
	fmt.Println(message)

	greetingAGroup()

	message, err := greetings.Hello("")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
}

func greetingAGroup() {
	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned map of
	// messages to the console.
	fmt.Println(messages)
}

package main

import (
	"log"
	"example.com/greetings"
	"fmt"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)


	// Get a greeting message  for a single person and print it.
	//message, err := greetings.Hello("Bobby")

	// Get a slice of names
	names := []string {"Bob", "Jil", "Ted"}
	messages, err := greetings.Hellos(names)

	// If an error was returned, print it to the console
	// and exit the program
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, printthe returned message
	// to the console
	fmt.Println(messages)
}

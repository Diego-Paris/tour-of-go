package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
		
		//If no name was given, return an error with a message
		if name == "" {
			return "", errors.New("empty name")
		}
	
		// Return a greeting that embeds the name in a message.
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

// Hellos returns a map that associates each of the 
// named people with a greetings message
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages
	messages := make(map[string]string)

	// Loop through the received slice, we call the 
	// hello function to get a msg for each name
	for _, name := range names {
		
		// get the msg and err
		message, err := Hello(name)
		
		// if err is not nil then return nil and error
		if err != nil {
			return nil, err
		}

		// In the map, associate the retrieved message
		// with the name
		messages[name] = message
	}

	// return a successful map of messages and nil (no error)
	return messages, nil
}


// init sets an initial value for variables used in the function
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting
// messages. The returned message is selected at random
func randomFormat() string {
	// A slice of message formats
	formats := []string {
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by
	// specifying a random index for the slice of formats
	return formats[rand.Intn(len(formats))]

}
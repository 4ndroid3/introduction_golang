package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return name, errors.New("empty name")
	}
	// Create a message using a random format.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// init sets initial values for variables used in the function.
func init() { // se inicia automaticamente antes de que se inicialicen todos los campos de el modulo.
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string { // cuando una funcion empieza en minusculas se indica que la funcion va a ser local. Osea no va a poder accederse desde afuera.
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}

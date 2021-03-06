package main

import (
	"fmt"
)

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// very custom logic for generating an english greeting
	return "Hi there"
}

// same as above, we can remove eb as a argument because we do not use it inside of a receiver
// func (eb englishBot) getGreeting() string {
// 	// very custom logic for generating an english greeting
// 	return "Hi there"
// }

func (sb spanishBot) getGreeting() string {
	// very custom logic for generating a spanish greeting
	return "Hola"
}

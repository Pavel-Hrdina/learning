package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	// Get a greetings and print it
	message := greetings.Hello("Pavel")
	fmt.Println(message)
}

package greetings

import "fmt"

func Hello(name string) string {
	// Return a greeting
	var message string
	message = fmt.Sprintf("Hello %v.", name)
	return message
}

package main

import "fmt"

const spanish = "Spanish"
const EnglishHelloPrefix = "Hello, "
const SpanishHelloPrefix = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World!"
	}

	prefix := EnglishHelloPrefix

	switch language {
	case spanish:
		prefix = SpanishHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello(""))
}

package main

import "fmt"

const spanish = "Spanish"
const EnglishHelloPrefix = "Hello, "
const SpanishHelloPrefix = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World!"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = SpanishHelloPrefix
	default:
		prefix = EnglishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("", "English"))
}

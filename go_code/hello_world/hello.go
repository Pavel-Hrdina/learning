package main

import "fmt"

const spanish = "Spanish"
const EnglishHelloPrefix = "Hello, "
const SpanishHelloPrefix = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		return EnglishHelloPrefix + "World!"
	}
	return HollyHello + name
}

func main() {
	fmt.Println(Hello(""))
}

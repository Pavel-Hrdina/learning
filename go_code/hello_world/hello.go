package main

import "fmt"

const spanish = "Spanish"
const EnglishHelloPrefix = "Hello, "
const SpanishHelloPrefix = "Hola, "

func Hello(name string) string {
	if name == "" {
		return HollyHello + "World!"
	}
	return HollyHello + name
}

func main() {
	fmt.Println(Hello(""))
}

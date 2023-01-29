package main

import "fmt"

const holly_hello = "Hello, "

func hello(name string) string {
	return holly_hello + name
}

func main() {
	fmt.Printf(hello("World"))
}

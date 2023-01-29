package main

import "fmt"

const HollyHello = "Hello, "

func Hello(name string) string {
	if name == "" {
		return HollyHello + "World!"
	}
	return HollyHello + name
}

func main() {
	fmt.Println(Hello(""))
}

package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("greet people", func(t *testing.T) {
		got := Hello("Joe", "")
		want := "Hello, Joe"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say, 'Hello, World' when empty string is given", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Joe", "Spanish")
		want := "Hola, Joe"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

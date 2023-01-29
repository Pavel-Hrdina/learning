package main

import "testing"

func test_hello(t *testing.T) {
	t.Run("greet people", func(t *testing.T) {
		got := hello("Joe")
		want := "Hello, Joe"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("say, 'Hello, World' when empty string is given", func(t *testing.T) {
		got := hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

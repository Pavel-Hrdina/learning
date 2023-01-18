package main

import "testing"

func test_hello(t *testing.T) {
	got := hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

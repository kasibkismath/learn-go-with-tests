package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Kasib", "")
		want := "Hello, Kasib!"
		assert(t, got, want)
	})
	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assert(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Kasib", "Spanish")
		want := "Hola, Kasib!"
		assert(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Kasib", "French")
		want := "Bonjour, Kasib!"
		assert(t, got, want)
	})
}

func assert(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

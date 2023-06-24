package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello, Chris"
		asserCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		asserCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		asserCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Jean", "French")
		want := "Bonjour, Jean"
		asserCorrectMessage(t, got, want)
	})
	t.Run("in Hindi", func(t *testing.T) {
		got := Hello("Kumar", "Hindi")
		want := "Ramram, Kumar"
		asserCorrectMessage(t, got, want)
	})
}

func asserCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got: %q\nwant: %q", got, want)
	}
}

package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say hellop to people", func(t *testing.T) {
		got := Hello("Calin", "")
		want := "Hello, Calin"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "a")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Jack", "French")
		want := "Bonjour, Jack"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Romanian", func(t *testing.T) {
		got := Hello("Ion", "Romanian")
		want := "Salut, Ion"
		assertCorrectMessage(t, got, want)
	})
}

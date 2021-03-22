package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Calin")
	want := "Hello, Calin"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
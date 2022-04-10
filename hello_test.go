package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to people with name", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("say hello to people with name", func(t *testing.T) {
		got := Hello("David")
		want := "Hello, David"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("say hello to people with name", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

}

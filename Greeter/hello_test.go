package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q wanted %q", got, want)
		}
	}

	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Dan", "")
		want := "Hello, Dan!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Says 'Hello, World!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("In Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("In French", func(t *testing.T) {
		got := Hello("Pablo", "French")
		want := "Bonjour, Pablo!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("In Japanese", func(t *testing.T) {
		got := Hello("Kim", "Japanese")
		want := "Kon'nichiwa, Kim!"

		assertCorrectMessage(t, got, want)
	})
}

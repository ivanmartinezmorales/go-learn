package main

import "testing"

func TestHello(t *testing.T) {

	// helper function to help assertCorrectMessage
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got: '%s', want: '%s'", got, want)
		}
	}

	t.Run("Saying hello to Carlos!", func(t *testing.T) {
		got := Greeting("Carlos")
		want := "Hello, Carlos"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello to a stranger.", func(t *testing.T) {
		got := Greeting("Stranger")
		want := "Hello, Stranger"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello to an empty string.", func(t *testing.T) {
		got := Greeting("")
		want := "Hello, Stranger"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Greeting the user in Spanish.", func(t *testing.T) {
		got := Hello("Juan", "Spanish")
		want := "Hola, Juan"
		assertCorrectMessage(t, got, want)
	})
}

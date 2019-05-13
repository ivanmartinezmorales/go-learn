/**
 * Testing our Greeting function.
 */

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
		got := Greeting("Carlos", "English")
		want := "Hello, Carlos"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello to a stranger.", func(t *testing.T) {
		got := Greeting("Stranger", "English")
		want := "Hello, Stranger"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello to an empty string.", func(t *testing.T) {
		got := Greeting("", "English")
		want := "Hello, Stranger"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Greeting the user in Spanish.", func(t *testing.T) {
		got := Greeting("Juan", "Spanish")
		want := "Hola, Juan"
		assertCorrectMessage(t, got, want)
	})

	t.Run("A greeting in french.", func(t *testing.T) {
		got := Greeting("Stephen", "French")
		want := "Bonjour, Stephen"
		assertCorrectMessage(t, got, want)
	})
}

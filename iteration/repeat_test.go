package iteration

import "testing"

func TestRepeat(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, repeated, expected string) {
		t.Helper()
		if repeated != expected {
			t.Errorf("expected: '%s', but got '%s'", expected, repeated)
		}
	}

	t.Run("Testing repeating 'a' 0 times.", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := ""
		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("Testing 'a' repeated 1 time", func(t *testing.T) {
		repeated := Repeat("a", 1)
		expected := "a"
		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("Repeating 'zyx' repeated 5 times.", func(t *testing.T) {
		repeated := Repeat("zyx", 5)
		expected := "zyxzyxzyxzyxzyx"
		assertCorrectMessage(t, repeated, expected)
	})
}

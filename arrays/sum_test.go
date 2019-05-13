package arrays

import "testing"

func TestSum(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("Got: '%d' but wanted: '%d'", got, want)
		}
	}

	t.Run("Testing if the sum of the numbers is 15.", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		assertCorrectMessage(t, got, want)
	})
}

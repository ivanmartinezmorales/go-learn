package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("Got: '%.2f' but wanted: '%.2f'", got, want)
		}
	}

	t.Run("Testing perimeter of a square with sides: 10.0", func(t *testing.T) {
		got := Perimeter(10.0, 10.0)
		want := 40.0
		assertCorrectMessage(t, got, want)
	})

	t.Run("Testing perimeter with height 1000.0, width: 1234.0", func(t *testing.T) {
		got := Perimeter(1000.0, 1234.0)
		want := 4468.0
		assertCorrectMessage(t, got, want)
	})

	t.Run("Testing perimeter with height 0, width 0.", func(t *testing.T) {
		got := Perimeter(0.0, 0.0)
		want := 0.0
		assertCorrectMessage(t, got, want)
	})
}

func TestArea(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("Got: '%.2f' but wanted: '%.2f'", got, want)
		}
	}

	t.Run("Testing area with height 0, width 0", func(t *testing.T) {
		got := Area(0.0, 0.0)
		want := 0.0
		assertCorrectMessage(t, got, want)
	})

	t.Run("Testing area with height 10, width 100", func(t *testing.T) {
		got := Area(10.0, 100.0)
		want := 1000.0
		assertCorrectMessage(t, got, want)
	})

	t.Run("Testing area with height -1, width -1", func(t *testing.T) {
		got := Area(-1.0, -1.0)
		want := 1.0
		assertCorrectMessage(t, got, want)
	})
}

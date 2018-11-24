package geometry

import (
	"testing"
)

func TestGeometry(t *testing.T) {
	r := Rectangle{2.5, 2.0}
	c := Circle{1.0}

	assertCorrectMessage := func(t *testing.T, got, want float64) {
		//t.Helper()
		if got != want {
			t.Errorf("got %f, wanted %f", got, want)
		}
	}

	t.Run("calculating the perimeter of a rectangle given a height and a width", func(t *testing.T) {
		got := r.Perimeter()
		want := 9.0

		assertCorrectMessage(t, got, want)
	})

	t.Run("calculating the area of a recatangle givben a height and a width", func(t *testing.T) {
		got := r.Area()
		want := 5.0

		assertCorrectMessage(t, got, want)
	})

	t.Run("calculating the perimeter of a circle", func(t *testing.T) {
		got := c.Perimeter()
		want := 6.28

		assertCorrectMessage(t, got, want)
	})

	t.Run("calculating the area of a circle", func(t *testing.T) {
		got := c.Area()
		want := 3.14

		assertCorrectMessage(t, got, want)
	})
}

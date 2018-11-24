package geometry

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rect := Rectangle{2.5, 2.0}

	assertCorrectMessage := func(t *testing.T, got, want float64) {
		//t.Helper()
		if got != want {
			t.Errorf("got %f, wanted %f", got, want)
		}
	}

	t.Run("calculating the perimeter of a rectangle given a height and a width", func(t *testing.T) {
		got := Perimeter(rect)
		want := 9.0

		assertCorrectMessage(t, got, want)
	})

	t.Run("calculating the area of a recatangle givben a height and a width", func(t *testing.T) {
		got := Area(rect)
		want := 5.0

		assertCorrectMessage(t, got, want)
	})
}

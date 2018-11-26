package geometry

import (
	"testing"
)

func TestGeometry(t *testing.T) {
	r := Rectangle{2.5, 2.0}
	c := Circle{1.0}

	checkArea := func(t *testing.T, s Shape, want float64) {
		//t.Helper()
		got := s.Area()
		if got != want {
			t.Errorf("got %f, wanted %f", got, want)
		}
	}
	checkPerimeter := func(t *testing.T, s Shape, want float64) {
		//t.Helper()
		got := s.Perimeter()
		if got != want {
			t.Errorf("got %f, wanted %f", got, want)
		}
	}

	t.Run("calculating the perimeter of a rectangle given a height and a width", func(t *testing.T) {
		want := 9.0

		checkPerimeter(t, r, want)
	})

	t.Run("calculating the area of a recatangle givben a height and a width", func(t *testing.T) {
		want := 5.0

		checkArea(t, r, want)
	})

	t.Run("calculating the perimeter of a circle", func(t *testing.T) {
		want := 6.28

		checkPerimeter(t, c, want)
	})

	t.Run("calculating the area of a circle", func(t *testing.T) {
		want := 3.14

		checkArea(t, c, want)
	})
}

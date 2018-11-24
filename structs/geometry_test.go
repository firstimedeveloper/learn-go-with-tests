package geometry

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Run("calculating the perimeter of a rectangle given a height and a width", func(t *testing.T) {
		got := Perimeter(2.5, 2.0)
		want := 9.0

		if got != want {
			t.Errorf("got %f, wanted %f", got, want)
		}
	})
}

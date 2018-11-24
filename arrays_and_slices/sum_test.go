package sum

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	t.Run("collection of slices of numbers", func(t *testing.T) {
		got := SumAll([]int{3, 4}, []int{4, 5})
		want := []int{7, 9}

		// comparing the values of got and want
		// returns true if any two values are equal
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func ExampleSumAll() {
	fmt.Println(SumAll([]int{3, 5}, []int{3, 2}))
	fmt.Println(SumAll([]int{}, []int{3, 2}))

	// Output: [8 5]
	// [0 5]
}

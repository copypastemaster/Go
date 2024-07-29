package main

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("colection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("Got %d want %d given %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("Got %d want %d given: %v", got, want, numbers)
		}
	})
}

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func ExampleSum() {
	numbers := []int{1, 2, 3, 4, 5}
	total := Sum(numbers)
	fmt.Println(total)
	// Output: 15
}

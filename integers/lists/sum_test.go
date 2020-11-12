package main

import "testing"

func TestSum(t *testing.T) {
	numbers := []int{1,2,3,4,5}

	t.Run("collection of 5 numbers", func(t *testing.T) {
		// alternate way to declare
		//numbers := [...]int{1,2,3,4,5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v",got, want, numbers)
		}

	})
}

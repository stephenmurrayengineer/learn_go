package main

import "testing"

func TestArea(t *testing.T) {

	// slice of structs to iterate over:
	areaTests := []struct {
			shape Shape
			want float64
	}{
		{Rectangle{12,6}, 72},
		{Circle{10}, 314.1592653589793},
		{Triangle{12,6}, 36},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}

}
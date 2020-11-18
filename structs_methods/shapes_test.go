package main

import "testing"

func TestArea(t *testing.T) {

	// slice of structs to iterate over:
	areaTests := []struct {
			shape Shape
			want float64
	}{
		{shape: Rectangle{Width: 12,Height: 6}, want: 72},
		{shape: Circle{Radius: 10}, want: 314.1592653589793},
		{shape: Triangle{Base: 12, Height:6}, want: 36},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
		}
	}

}
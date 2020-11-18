package main

import "testing"

func TestArea(t *testing.T) {

	// slice of structs to iterate over:
	areaTests := []struct {
			shape Shape
			hasArea float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, hasArea: 72},
		{shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, hasArea: 36},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v got %g hasArea %g", tt.shape, got, tt.hasArea)
		}
	}

}
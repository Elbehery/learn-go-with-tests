package structs

import "testing"

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		exp   float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, exp: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, exp: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, exp: 36.0},
	}

	for _, tc := range areaTests {
		act := tc.shape.Area()
		if act != tc.exp {
			t.Errorf("%#v expected %v, but got %v instead", tc.shape, tc.exp, act)
		}
	}
}

package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	expect := 40.0

	if got != expect {
		t.Errorf("got %.2f expected %.2f", got, expect)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name   string
		shape  Shape
		expect float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, expect: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, expect: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, expect: 36.0},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {

		})
		got := tt.shape.Area()
		if got != tt.expect {
			t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.expect)
		}
	}
}

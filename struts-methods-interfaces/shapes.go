package shapes

import "math"

// Shape - interface
type Shape interface {
	Area() float64
}

// Rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

// Area - method on Rectangle struct
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle struct
type Circle struct {
	Radius float64
}

// Area - method on Circle struct
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle struct
type Triangle struct {
	Base   float64
	Height float64
}

// Area - method on Triangle struct
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

// Perimeter - retuns the perimeter of a shape given the width and height
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Height + rectangle.Width)
}

// Area - retuns the Area of a shape given the width and height
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

package shapes

import "math"

// Shape is implemented by anything that can tell us its area
type Shape interface {
	Area() float64
}

// Rectangle contains the dimensions of a rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter returns the perimeter of a rectangle
func Perimeter(r Rectangle) float64 {
	return (r.Width * 2) + (r.Height * 2)
}

// Area returns the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle contains the dimensions of a circle
type Circle struct {
	Radius float64
}

// Area returns the area of a circle
func (c Circle) Area() float64 {
	return math.Pow(c.Radius, 2) * math.Pi
}

// Circumference returns the circumference of the circle
func Circumference(c Circle) float64 {
	return 2 * math.Pi * c.Radius
}

// Triangle represents the dimensions of the triangle
type Triangle struct {
	Base   float64
	Height float64
}

// Area returns the area of the triangle
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
